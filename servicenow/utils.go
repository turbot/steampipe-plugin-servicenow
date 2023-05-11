package servicenow

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/go-servicenow/servicenow"
	"github.com/turbot/steampipe-plugin-sdk/v5/connection"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-servicenow/model"
)

// connect:: returns servicenow client after authentication
func connectRaw(ctx context.Context, cm *connection.Manager, c *plugin.Connection) (*servicenow.ServiceNow, error) {
	// // Load connection from cache, which preserves throttling protection etc
	// cacheKey := "simpleforce"
	// if cachedData, ok := cm.Cache.Get(cacheKey); ok {
	// 	return cachedData.(*servicenow.ServiceNow), nil
	// }

	// config := GetConfig(c)
	// apiVersion := simpleforce.DefaultAPIVersion
	// clientID := "steampipe"
	// securityToken := ""

	// if config.ClientId != nil {
	// 	clientID = *config.ClientId
	// }

	// if config.APIVersion != nil {
	// 	apiVersion = *config.APIVersion
	// }

	// if config.Username == nil {
	// 	plugin.Logger(ctx).Warn("servicenow.connectRaw", "'username' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	// 	return nil, nil
	// }

	// if config.Password == nil {
	// 	plugin.Logger(ctx).Warn("servicenow.connectRaw", "'password' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	// 	return nil, nil
	// }

	// // The Servicenow security token is only required If the client's IP address is not added to the organization's list of trusted IPs
	// // https://help.servicenow.com/s/articleView?id=sf.security_networkaccess.htm&type=5
	// // https://migration.trujay.com/help/how-to-add-an-ip-address-to-whitelist-on-servicenow/
	// if config.Token != nil {
	// 	securityToken = *config.Token
	// }

	// // setup client
	// client := simpleforce.NewClient(*config.URL, clientID, apiVersion)
	// if client == nil {
	// 	plugin.Logger(ctx).Error("servicenow.connectRaw", "couldn't get servicenow client. Client setup error.")
	// 	return nil, fmt.Errorf("servicenow.connectRaw couldn't get servicenow client. Client setup error.")
	// }

	// // LoginPassword signs into servicenow using password. token is optional if trusted IP is configured.
	// // Ref: https://developer.servicenow.com/docs/atlas.en-us.214.0.api_rest.meta/api_rest/intro_understanding_username_password_oauth_flow.htm
	// // Ref: https://developer.servicenow.com/docs/atlas.en-us.214.0.api.meta/api/sforce_api_calls_login.htm
	// err := client.LoginPassword(*config.Username, *config.Password, securityToken)
	// if err != nil {
	// 	plugin.Logger(ctx).Error("servicenow.connectRaw", "client login error", err)
	// 	return nil, fmt.Errorf("client login error %v", err)
	// }

	// // Save to cache
	// cm.Cache.Set(cacheKey, client)

	// return client, nil
	return nil, nil
}

// generateQuery:: returns sql query based on the column names, table name passed
func generateQuery(columns []*plugin.Column, tableName string) string {
	var queryColumns []string
	for _, column := range columns {
		queryColumns = append(queryColumns, getServicenowColumnName(column.Name))
	}

	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(queryColumns, ", "), tableName)
}

// decodeQueryResult(ctx, apiResponse, responseStruct):: converts raw apiResponse to required output struct
func decodeQueryResult(ctx context.Context, response interface{}, respObject interface{}) error {
	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// For debugging purpose - commenting out to avoid unnecessary logs
	// plugin.Logger(ctx).Info("decodeQueryResult", "Items", string(resp))
	err = json.Unmarshal(resp, respObject)
	if err != nil {
		return err
	}

	return nil
}

// buildQueryFromQuals :: generate SOQL based on the contions specified in sql query
// refrences
// - https://developer.servicenow.com/docs/atlas.en-us.234.0.soql_sosl.meta/soql_sosl/sforce_api_calls_soql_select_comparisonoperators.htm
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
						// In case of IN caluse
						if value.GetListValue() != nil {
							// continue
							switch qual.Operator {
							case "=":
								stringValueSlice := []string{}
								for _, q := range value.GetListValue().Values {
									stringValueSlice = append(stringValueSlice, q.GetStringValue())
								}
								if len(stringValueSlice) > 0 {
									filters = append(filters, fmt.Sprintf("%s IN ('%s')", getServicenowColumnName(filterQualItem.Name), strings.Join(stringValueSlice, "','")))
								}
							case "<>":
								stringValueSlice := []string{}
								for _, q := range value.GetListValue().Values {
									stringValueSlice = append(stringValueSlice, q.GetStringValue())
								}
								if len(stringValueSlice) > 0 {
									filters = append(filters, fmt.Sprintf("%s NOT IN ('%s')", getServicenowColumnName(filterQualItem.Name), strings.Join(stringValueSlice, "','")))
								}
							}
						} else {
							switch qual.Operator {
							case "=":
								filters = append(filters, fmt.Sprintf("%s = '%s'", getServicenowColumnName(filterQualItem.Name), value.GetStringValue()))
							case "<>":
								filters = append(filters, fmt.Sprintf("%s != '%s'", getServicenowColumnName(filterQualItem.Name), value.GetStringValue()))
							}
						}
					case proto.ColumnType_BOOL:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s = %s", getServicenowColumnName(filterQualItem.Name), "FALSE"))
						case "=":
							filters = append(filters, fmt.Sprintf("%s = %s", getServicenowColumnName(filterQualItem.Name), "TRUE"))
						}
					case proto.ColumnType_INT:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s != %d", getServicenowColumnName(filterQualItem.Name), value.GetInt64Value()))
						default:
							filters = append(filters, fmt.Sprintf("%s %s %d", getServicenowColumnName(filterQualItem.Name), qual.Operator, value.GetInt64Value()))
						}
					case proto.ColumnType_DOUBLE:
						switch qual.Operator {
						case "<>":
							filters = append(filters, fmt.Sprintf("%s != %f", getServicenowColumnName(filterQualItem.Name), value.GetDoubleValue()))
						default:
							filters = append(filters, fmt.Sprintf("%s %s %f", getServicenowColumnName(filterQualItem.Name), qual.Operator, value.GetDoubleValue()))
						}
					// Need a way to distinguish b/w date and dateTime fields
					case proto.ColumnType_TIMESTAMP:
						// https://developer.servicenow.com/docs/atlas.en-us.234.0.soql_sosl.meta/soql_sosl/sforce_api_calls_soql_select_dateformats.htm
						if servicenowCols[filterQual.Name] == "date" {
							switch qual.Operator {
							case "=", ">=", ">", "<=", "<":
								filters = append(filters, fmt.Sprintf("%s %s %s", getServicenowColumnName(filterQualItem.Name), qual.Operator, value.GetTimestampValue().AsTime().Format("2006-01-02")))
							}
						} else {
							switch qual.Operator {
							case "=", ">=", ">", "<=", "<":
								filters = append(filters, fmt.Sprintf("%s %s %s", getServicenowColumnName(filterQualItem.Name), qual.Operator, value.GetTimestampValue().AsTime().Format("2006-01-02T15:04:05Z")))
							}
						}
					}
				}
			}

		}
	}

	if len(filters) > 0 {
		return strings.Join(filters, " AND ")
	}

	return ""
}

func getServicenowColumnName(name string) string {
	var columnName string
	// Servicenow custom fields are suffixed with '__c' and are not converted to
	// snake case in the table schema, so use the column name as is
	if strings.HasSuffix(name, "__c") {
		columnName = name
	} else {
		columnName = strcase.ToCamel(name)
	}
	return columnName
}

// append the dynamic columns with static columns for the table
func mergeTableColumns(ctx context.Context, p *plugin.Plugin, dynamicColumns []*plugin.Column, staticColumns []*plugin.Column) []*plugin.Column {
	var columns []*plugin.Column
	columns = append(columns, staticColumns...)
	for _, col := range dynamicColumns {
		if isColumnAvailable(col.Name, staticColumns) {
			continue
		}
		columns = append(columns, col)
	}
	return columns
}

// isColumnAvailable:: Checks if the column is not present in the existing columns slice
func isColumnAvailable(columnName string, columns []*plugin.Column) bool {
	for _, col := range columns {
		if col.Name == columnName {
			return true
		}
	}
	return false
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
