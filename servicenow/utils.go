package servicenow

import (
	"context"
	"fmt"
	"strings"

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

		// Check only if filter qual map matches with optional column name
		if filterQual.Name == filterQualItem.Name {
			if filterQual.Quals == nil {
				continue
			}

			for _, qual := range filterQual.Quals {
				if qual.Value != nil {
					value := qual.Value
					switch filterQualItem.Type {
					case proto.ColumnType_STRING:
						// Ignore the IN clause
						if value.GetListValue() != nil {
							continue
						}
						switch qual.Operator {
						case "=":
							filters = append(filters, fmt.Sprintf("%s=%s", filterQualItem.Name, value.GetStringValue()))
						case "<>":
							filters = append(filters, fmt.Sprintf("%s!=%s", filterQualItem.Name, value.GetStringValue()))
						}
					case proto.ColumnType_INT:
						if qual.Operator == "=" {
							filters = append(filters, fmt.Sprintf("%s%s%d", filterQualItem.Name, qual.Operator, value.GetInt64Value()))
						}
					case proto.ColumnType_DOUBLE:
						if qual.Operator == "=" {
							filters = append(filters, fmt.Sprintf("%s%s%f", filterQualItem.Name, qual.Operator, value.GetDoubleValue()))
						}
					}
				}
			}
		}
	}

	if len(filters) > 0 {
		return strings.Join(filters, "^")
	}

	return ""
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
	err := builder.client.NowTable.Read(model.SysDbObjectTableName, sysId, true, &returned)
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
		builder.GetTableColumns(serviceNowParentTable.Name, serviceNowParentTable.SuperClass, serviceNowColumns)
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
