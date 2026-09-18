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
			if qual.Value == nil || qual.Value.GetListValue() != nil {
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
