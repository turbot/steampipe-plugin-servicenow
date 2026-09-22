package servicenow

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/turbot/go-servicenow/servicenow"
	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
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
				// lists through unaltered. The FDW still pushes the limit down for them, so lists
				// are sent with ServiceNow's IN operator, which is a superset of the SQL predicate
				// that rowMatchesQuals narrows. A `not in` list is left entirely client-side, since
				// negating a case-insensitive match drops rows Postgres would keep.
				if qual.Operator != "=" || len(list.Values) == 0 {
					continue
				}
				values := make([]string, 0, len(list.Values))
				switch filterQualItem.Type {
				case proto.ColumnType_INT:
					for _, v := range list.Values {
						values = append(values, strconv.FormatInt(v.GetInt64Value(), 10))
					}
				case proto.ColumnType_STRING:
					for _, v := range list.Values {
						// IN is comma delimited, so a value containing a comma cannot be expressed.
						// Pushing the rest of the list would narrow the filter, so drop all of it
						// and let the client-side recheck do the work.
						if !snowValuePushable(v.GetStringValue(), true) {
							values = nil
							break
						}
						values = append(values, v.GetStringValue())
					}
				}
				if len(values) > 0 {
					filters = append(filters, fmt.Sprintf("%sIN%s", filterQualItem.Name, strings.Join(values, ",")))
				}
				continue
			}

			value := qual.Value
			switch filterQualItem.Type {
			case proto.ColumnType_STRING:
				// ServiceNow compares strings case-insensitively, so = returns a superset of what
				// Postgres asked for and rowMatchesQuals narrows it. != is the complement of that
				// superset: it excludes rows differing only by case, which Postgres would keep, and
				// no client-side filtering can add them back, so it is not pushed down.
				if qual.Operator == "=" && snowValuePushable(value.GetStringValue(), false) {
					filters = append(filters, fmt.Sprintf("%s=%s", filterQualItem.Name, value.GetStringValue()))
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

// snowValuePushable reports whether a string can be carried inside sysparm_query unchanged.
// There is no escaping in an encoded query: ^ separates conditions and , separates IN values, so a
// value containing one is silently reinterpreted rather than rejected - the API drops the malformed
// fragment and answers a different question. Those quals are not pushed down at all; the unfiltered
// rows are a superset, which the client-side recheck then narrows.
func snowValuePushable(value string, inList bool) bool {
	if strings.Contains(value, "^") {
		return false
	}
	return !inList || !strings.Contains(value, ",")
}

// rawValueEquals compares a raw ServiceNow field value against a qual value exactly. ServiceNow
// returns every field as a string, so each type is parsed back before comparing. The second return
// value reports whether the comparison could be made at all; when it is false the row is left for
// Postgres to judge rather than being dropped on a guess.
func rawValueEquals(raw interface{}, columnType proto.ColumnType, qualValue *proto.QualValue) (bool, bool) {
	text, isText := raw.(string)

	switch columnType {
	case proto.ColumnType_STRING:
		if !isText {
			return false, false
		}
		return text == qualValue.GetStringValue(), true
	case proto.ColumnType_INT:
		if !isText {
			return false, false
		}
		number, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return false, false
		}
		return number == qualValue.GetInt64Value(), true
	case proto.ColumnType_DOUBLE:
		if !isText {
			return false, false
		}
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return false, false
		}
		return number == qualValue.GetDoubleValue(), true
	case proto.ColumnType_BOOL:
		if !isText {
			if flag, isBool := raw.(bool); isBool {
				return flag == qualValue.GetBoolValue(), true
			}
			return false, false
		}
		flag, err := strconv.ParseBool(text)
		if err != nil {
			return false, false
		}
		return flag == qualValue.GetBoolValue(), true
	case proto.ColumnType_TIMESTAMP:
		rowTime, ok := parseSnowTime(raw)
		if !ok || qualValue.GetTimestampValue() == nil {
			return false, false
		}
		return rowTime.Equal(qualValue.GetTimestampValue().AsTime()), true
	}

	return false, false
}

// parseSnowTime parses a datetime as ServiceNow returns it. The values are naive, so reading them as
// UTC matches what Postgres does with the same string and keeps this recheck in agreement with it.
func parseSnowTime(raw interface{}) (time.Time, bool) {
	text, isText := raw.(string)
	if !isText {
		return time.Time{}, false
	}
	parsed, err := time.Parse("2006-01-02 15:04:05", text)
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

// rowMatchesQuals applies the pushed down quals to a row exactly. The FDW pushes the limit down
// when every qual is on a key column, so each row streamed must really match, or a row Postgres
// discards consumes part of the limit and the query silently under-returns. Three cases are not
// exact in sysparm_query: timestamp bounds are widened by the max UTC offset, string = and IN are
// case-insensitive, and filters that cannot be expressed at all are not pushed. ServiceNow's !=
// also returns empty values, which never satisfy a comparison in Postgres.
func rowMatchesQuals(row map[string]interface{}, quals plugin.KeyColumnQualMap, columns []*plugin.Column) bool {
	for _, column := range columns {
		if quals[column.Name] == nil {
			continue
		}
		for _, qual := range quals[column.Name].Quals {
			if qual.Value == nil || snowOperator(qual.Operator) == "" {
				continue
			}

			// A field absent from the API response says nothing about the qual, but a field that is
			// present and empty (nil after sanitizeTableObject) never satisfies one in Postgres,
			// including `in` and `not in`.
			value, present := row[column.Name]
			if !present {
				continue
			}
			if value == nil {
				return false
			}

			// Only equality is rechecked here. The ordered operators are exact in sysparm_query
			// for every type except timestamps, which are handled separately below, and Postgres
			// collates strings differently from Go, so ordering is never second-guessed.
			isEquality := qual.Operator == "=" || qual.Operator == "<>"

			// An `in` list reaches the API as a case-insensitive IN, or not at all; either way
			// membership has to be rechecked here. A `not in` list is never pushed down.
			if list := qual.Value.GetListValue(); list != nil {
				if !isEquality {
					continue
				}
				matched, decided := false, true
				for _, listValue := range list.Values {
					equal, ok := rawValueEquals(value, column.Type, listValue)
					if !ok {
						// Nothing reliable can be said about this list, so leave it to Postgres.
						decided = false
						break
					}
					if equal {
						matched = true
						break
					}
				}
				if decided && matched != (qual.Operator == "=") {
					return false
				}
				continue
			}

			switch column.Type {
			case proto.ColumnType_STRING:
				// ServiceNow matched this case-insensitively, so repeat it case-sensitively.
				equal, ok := rawValueEquals(value, column.Type, qual.Value)
				if !isEquality || !ok {
					continue
				}
				if equal != (qual.Operator == "=") {
					return false
				}
			case proto.ColumnType_TIMESTAMP:
				// The bounds went out widened by the max UTC offset, so re-apply them exactly.
				rowTime, ok := parseSnowTime(value)
				if !ok || qual.Value.GetTimestampValue() == nil {
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
			// INT, DOUBLE and BOOL scalars are exact in sysparm_query once the empty value case
			// above is handled, so they need no recheck.
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
