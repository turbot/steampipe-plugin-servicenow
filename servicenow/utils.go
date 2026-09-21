package servicenow

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/turbot/go-servicenow/servicenow"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-servicenow/model"
)

func buildQueryFromQuals(equalQuals plugin.KeyColumnQualMap, tableColumns []*plugin.Column, servicenowCols map[string]string) string {
	filters := []string{}

	for _, filterQualItem := range tableColumns {
		filterQual := equalQuals[filterQualItem.Name]
		if filterQual == nil {
			continue
		}

		if filterQual.Name != filterQualItem.Name {
			continue
		}
		if filterQual.Quals == nil {
			continue
		}

		for _, qual := range filterQual.Quals {
			if qual.Value == nil {
				continue
			}
			if list := qual.Value.GetListValue(); list != nil {
				// The SDK splits a single IN list into one call per value, but passes two or more
				// lists through unaltered. The FDW still pushes the limit down for them, so integer
				// lists are sent with ServiceNow's IN operator to keep the filter exact.
				if filterQualItem.Type == proto.ColumnType_INT && qual.Operator == "=" && len(list.Values) > 0 {
					values := make([]string, len(list.Values))
					for i, v := range list.Values {
						values[i] = strconv.FormatInt(v.GetInt64Value(), 10)
					}
					filters = append(filters, fmt.Sprintf("%sIN%s", filterQualItem.Name, strings.Join(values, ",")))
				}
				continue
			}

			value := qual.Value
			switch filterQualItem.Type {
			case proto.ColumnType_STRING:
				switch qual.Operator {
				case "=":
					filters = append(filters, fmt.Sprintf("%s=%s", filterQualItem.Name, value.GetStringValue()))
				case "<>":
					filters = append(filters, fmt.Sprintf("%s!=%s", filterQualItem.Name, value.GetStringValue()))
				}
			case proto.ColumnType_INT:
				op := snowOperator(qual.Operator)
				if op != "" {
					filters = append(filters, fmt.Sprintf("%s%s%d", filterQualItem.Name, op, value.GetInt64Value()))
				}
			case proto.ColumnType_DOUBLE:
				op := snowOperator(qual.Operator)
				if op != "" {
					filters = append(filters, fmt.Sprintf("%s%s%s", filterQualItem.Name, op, strconv.FormatFloat(value.GetDoubleValue(), 'f', -1, 64)))
				}
			case proto.ColumnType_TIMESTAMP:
				// ServiceNow interprets datetime literals in the API user's timezone.
				// Widen bounds by the max UTC offset so the API returns a superset;
				// Steampipe re-applies the exact predicate client-side.
				ts := value.GetTimestampValue()
				if ts == nil {
					continue
				}
				const maxOffset = 14 * time.Hour
				const layout = "2006-01-02 15:04:05"
				t := ts.AsTime().UTC()
				switch qual.Operator {
				case ">", ">=":
					filters = append(filters, fmt.Sprintf("%s>=%s", filterQualItem.Name, t.Add(-maxOffset).Format(layout)))
				case "<", "<=":
					filters = append(filters, fmt.Sprintf("%s<=%s", filterQualItem.Name, t.Add(maxOffset).Format(layout)))
				case "=":
					filters = append(filters, fmt.Sprintf("%s>=%s^%s<=%s",
						filterQualItem.Name, t.Add(-maxOffset).Format(layout),
						filterQualItem.Name, t.Add(maxOffset).Format(layout)))
				}
			case proto.ColumnType_BOOL:
				op := snowOperator(qual.Operator)
				if op == "=" || op == "!=" {
					filters = append(filters, fmt.Sprintf("%s%s%t", filterQualItem.Name, op, value.GetBoolValue()))
				}
			}
		}
	}

	if len(filters) > 0 {
		return strings.Join(filters, "^")
	}

	return ""
}

func snowOperator(op string) string {
	switch op {
	case "=":
		return "="
	case "<>":
		return "!="
	case ">":
		return ">"
	case ">=":
		return ">="
	case "<":
		return "<"
	case "<=":
		return "<="
	default:
		return ""
	}
}

// rowMatchesQuals applies the pushed down quals to a row exactly. The FDW pushes the limit down
// when every qual is on a key column, so each row streamed must really match. Two cases are not
// exact in sysparm_query: timestamp bounds are widened by the max UTC offset, and ServiceNow's
// != also returns empty values, which never satisfy a comparison in Postgres.
func rowMatchesQuals(row map[string]interface{}, quals plugin.KeyColumnQualMap, columns []*plugin.Column) bool {
	for _, column := range columns {
		if quals[column.Name] == nil {
			continue
		}
		for _, qual := range quals[column.Name].Quals {
			if qual.Value == nil || qual.Value.GetListValue() != nil || snowOperator(qual.Operator) == "" {
				continue
			}

			// A field absent from the API response says nothing about the qual, but a field that is
			// present and empty (nil after sanitizeTableObject) never satisfies one in Postgres.
			value, present := row[column.Name]
			if !present {
				continue
			}
			if value == nil {
				return false
			}

			// Only timestamps need an exact recheck. INT, DOUBLE and BOOL filters are exact in
			// sysparm_query once the empty value case above is handled.
			if column.Type != proto.ColumnType_TIMESTAMP || qual.Value.GetTimestampValue() == nil {
				continue
			}
			raw, ok := value.(string)
			if !ok {
				continue
			}
			// ServiceNow returns naive datetimes. Parsing as UTC here matches what Postgres does
			// with the same string, so this recheck agrees with the Postgres recheck.
			rowTime, err := time.Parse("2006-01-02 15:04:05", raw)
			if err != nil {
				continue
			}
			cmp := rowTime.Compare(qual.Value.GetTimestampValue().AsTime())
			var match bool
			switch qual.Operator {
			case "=":
				match = cmp == 0
			case ">":
				match = cmp > 0
			case ">=":
				match = cmp >= 0
			case "<":
				match = cmp < 0
			case "<=":
				match = cmp <= 0
			}
			if !match {
				return false
			}
		}
	}
	return true
}

func ignoreError(errors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range errors {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}

type ServiceNowTableColumn struct {
	Name        string
	Type        string
	Label       string
	Description string
}

type ServiceNowTableBuilder struct {
	client *servicenow.ServiceNow
	glides map[string]model.SysGlideObject
}

func NewServiceNowTableBuilder(client *servicenow.ServiceNow) (*ServiceNowTableBuilder, error) {
	builder := &ServiceNowTableBuilder{
		client: client,
	}

	err := builder.loadGlideObjectList()
	if err != nil {
		return nil, err
	}

	return builder, nil
}

func (builder *ServiceNowTableBuilder) loadGlideObjectList() error {
	var glidesResponse model.SysGlideObjectListResult
	err := builder.client.NowTable.List(model.SysGlideObjectTableName, 1000, 0, "", false, &glidesResponse)
	if err != nil {
		return err
	}
	if builder.glides == nil {
		builder.glides = make(map[string]model.SysGlideObject)
	}
	for _, glide := range glidesResponse.Result {
		builder.glides[glide.Name] = glide
	}
	return nil
}

func (builder *ServiceNowTableBuilder) GetTableByName(tableName string) (*model.SysDbObject, error) {
	var returned model.SysDbObjectListResult
	err := builder.client.NowTable.List(model.SysDbObjectTableName, 1, 0, fmt.Sprintf("name=%s", tableName), true, &returned)
	if err != nil {
		return nil, err
	}
	if len(returned.Result) == 0 {
		return nil, fmt.Errorf("table %s not found on ServiceNow", tableName)
	}
	return &returned.Result[0], nil
}

func (builder *ServiceNowTableBuilder) GetTableById(sysId string) (*model.SysDbObject, error) {
	var returned model.SysDbObjectGetResult
	err := builder.client.NowTable.Get(model.SysDbObjectTableName, sysId, true, &returned)
	if err != nil {
		return nil, err
	}
	return &returned.Result, nil
}

func (builder *ServiceNowTableBuilder) GetTableColumns(tableName string, parentTableSysId string, serviceNowColumns map[string]ServiceNowTableColumn) error {
	if parentTableSysId != "" {
		serviceNowParentTable, err := builder.GetTableById(parentTableSysId)
		if err != nil {
			return err
		}
		err = builder.GetTableColumns(serviceNowParentTable.Name, serviceNowParentTable.SuperClass, serviceNowColumns)
		if err != nil {
			return err
		}
	}

	servicenowObjectFields, err := builder.GetTableColumnsTypes(tableName)
	if err != nil {
		return err
	}

	fieldsDescriptions, err := builder.GetTableColumnsDescriptions(tableName)
	if err != nil {
		return err
	}

	for fieldName, fieldType := range servicenowObjectFields {
		serviceNowColumns[fieldName] = ServiceNowTableColumn{
			Name:        fieldName,
			Type:        fieldType,
			Label:       fieldsDescriptions[fieldName].Label,
			Description: fieldsDescriptions[fieldName].Hint,
		}
	}
	return nil
}

func (builder *ServiceNowTableBuilder) GetTableColumnsTypes(tableName string) (map[string]string, error) {
	columns := map[string]string{}
	limit := 1000
	offset := 0
	for {
		var returned model.SysDictionaryListResult
		err := builder.client.NowTable.List(model.SysDictionaryTableName, limit, offset, fmt.Sprintf("name=%s", tableName), false, &returned)
		if err != nil {
			return nil, err
		}
		totalReturned := len(returned.Result)
		for _, returnedObject := range returned.Result {
			if returnedObject.Element == "" {
				continue
			}
			if returnedObject.InternalType.Value == "glide_time" {
				columns[returnedObject.Element] = "glide_time"
				continue
			}

			// Find the scalar type of the column
			glide := builder.glides[returnedObject.InternalType.Value]

			// non-visible GUID fields are string typed
			if glide.ScalarType == "GUID" && glide.Visible != "true" {
				columns[returnedObject.Element] = "string"
				continue
			}

			columns[returnedObject.Element] = glide.ScalarType
		}

		if totalReturned < limit {
			break
		}
		offset += limit
	}

	return columns, nil
}

func (builder *ServiceNowTableBuilder) GetTableColumnsDescriptions(tableName string) (map[string]model.SysDocumentation, error) {
	columnsDescriptions := map[string]model.SysDocumentation{}
	limit := 1000
	offset := 0
	for {
		var returned model.SysDocumentationListResult
		err := builder.client.NowTable.List(model.SysDocumentationTableName, limit, offset, fmt.Sprintf("name=%s", tableName), false, &returned)
		if err != nil {
			return nil, err
		}
		totalReturned := len(returned.Result)
		for _, returnedObject := range returned.Result {
			columnsDescriptions[returnedObject.Element] = returnedObject
		}

		if totalReturned < limit {
			break
		}
		offset += limit
	}

	return columnsDescriptions, nil
}
