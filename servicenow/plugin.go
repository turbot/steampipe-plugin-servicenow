package servicenow

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-servicenow"

type contextKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.From(getFieldFromSObjectMapByColumnName).NullIfZero(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: ignoreError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: pluginTableDefinitions,
	}

	return p
}

func pluginTableDefinitions(ctx context.Context, d *plugin.TableMapData) (map[string]*plugin.Table, error) {
	client, err := ConnectUncached(ctx, d.Connection)
	if err != nil {
		return nil, err
	}

	// Initialize tables with static tables with static and dynamic columns(if credentials are set)
	tables := map[string]*plugin.Table{
		"servicenow_cmdb_ci_server":                              tableServicenowCmdbCiServer(),
		"servicenow_cmdb_ci_service":                             tableServicenowCmdbCiService(),
		"servicenow_cmdb_ci":                                     tableServicenowCmdbCi(),
		"servicenow_incident":                                    tableServicenowIncident(),
		"servicenow_now_consumer":                                tableServicenowNowConsumer(),
		"servicenow_now_contact":                                 tableServicenowNowContact(),
		"servicenow_sn_chg_rest_change_affected_cmdb_ci":         tableServicenowSnChgRestChangeAffectedCmdbCi(),
		"servicenow_sn_chg_rest_change_impacted_cmdb_ci_service": tableServicenowSnChgRestChangeImpactedCmdbCiService(),
		"servicenow_sn_chg_rest_change_model":                    tableServicenowSnChgRestChangeModel(),
		"servicenow_sn_chg_rest_change_task":                     tableServicenowSnChgRestChangeTask(),
		"servicenow_sn_chg_rest_change":                          tableServicenowSnChgRestChange(),
		"servicenow_sys_audit_relation":                          tableServicenowSysAuditRelation(),
		"servicenow_sys_audit":                                   tableServicenowSysAudit(),
		"servicenow_sys_group_has_role":                          tableServicenowSysGroupHasRole(),
		"servicenow_sys_user_grmember":                           tableServicenowSysUserGroupMember(),
		"servicenow_sys_user_group":                              tableServicenowSysUserGroup(),
		"servicenow_sys_user_has_role":                           tableServicenowSysUserHasRole(),
		"servicenow_sys_user_role":                               tableServicenowSysUserRole(),
		"servicenow_sys_user":                                    tableServicenowSysUser(),
	}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``
	servicenowTables := []string{}
	config := GetConfig(d.Connection)
	if config.Objects != nil && len(*config.Objects) > 0 {
		servicenowTables = append(servicenowTables, *config.Objects...)
	}

	builder, err := NewServiceNowTableBuilder(client)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(servicenowTables))
	for _, snowTable := range servicenowTables {
		tableName := "servicenow_" + strcase.ToSnake(re.ReplaceAllString(snowTable, substitution))
		plugin.Logger(ctx).Warn("tableName", tableName)
		go func(name string) {
			defer wg.Done()
			tableName := "servicenow_" + strcase.ToSnake(re.ReplaceAllString(name, substitution))
			plugin.Logger(ctx).Debug("servicenow.pluginTableDefinitions", "object_name", name, "table_name", tableName)
			ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
			ctx = context.WithValue(ctx, contextKey("ServicenowTableName"), name)
			table := generateDynamicTables(ctx, d, *builder)
			// Ignore if the requested Servicenow object is not present.
			if table != nil {
				tables[tableName] = table
			}
		}(snowTable)
	}
	wg.Wait()
	return tables, nil
}

func generateDynamicTables(ctx context.Context, _ *plugin.TableMapData, builder ServiceNowTableBuilder) *plugin.Table {
	logger := plugin.Logger(ctx)

	// Get the query for the metric (required)
	servicenowTableName := ctx.Value(contextKey("ServicenowTableName")).(string)
	tableName := ctx.Value(contextKey("PluginTableName")).(string)
	logger.Info("generating dynamic table:", tableName)

	serviceNowTableObject, err := builder.GetTableByName(servicenowTableName)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "table_generation_error", err)
		return nil
	}

	var columnsData = make(map[string]ServiceNowTableColumn)
	err = builder.GetTableColumns(serviceNowTableObject.Name, serviceNowTableObject.SuperClass, columnsData)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "column_build_error", err)
	}

	// Top columns
	cols := []*plugin.Column{}
	servicenowCols := map[string]string{}
	// Key columns
	keyColumns := plugin.KeyColumnSlice{}

	for columnFieldName, columnData := range columnsData {
		columnDescription := columnData.Description
		if columnDescription == "" {
			columnDescription = columnData.Label
		}
		column := plugin.Column{
			Name:        columnFieldName,
			Description: fmt.Sprintf("%s.", columnDescription),
			Transform:   transform.FromP(getFieldFromSObjectMap, columnFieldName),
		}
		// Adding column type in the map to help in qual handling
		servicenowCols[columnFieldName] = columnData.Type

		// Set column type based on the `soapType` from servicenow schema
		switch columnData.Type {
		case "string", "glide_date", "date", "time":
			column.Type = proto.ColumnType_STRING
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
		case "glide_time":
			column.Type = proto.ColumnType_STRING
			column.Transform.Transform(parseGlideTime)
		case "datetime":
			column.Type = proto.ColumnType_TIMESTAMP
			column.Transform.Transform(parseDateTime)
		case "boolean":
			column.Type = proto.ColumnType_BOOL
		case "double", "decimal", "float":
			column.Type = proto.ColumnType_DOUBLE
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "int", "integer", "longint":
			column.Type = proto.ColumnType_INT
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "GUID":
			column.Type = proto.ColumnType_JSON
		default:
			column.Type = proto.ColumnType_STRING
		}

		cols = append(cols, &column)
	}

	Table := plugin.Table{
		Name:        tableName,
		Description: fmt.Sprintf("%s.", serviceNowTableObject.Label),
		List: &plugin.ListConfig{
			KeyColumns: keyColumns,
			Hydrate:    listServicenowObjectsByTable(servicenowTableName, servicenowCols),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("sys_id"),
			Hydrate:    getServicenowObjectbyID(servicenowTableName),
		},
		Columns: cols,
	}

	return &Table
}

func parseGlideTime(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil {
		return nil, nil
	}
	timeStr := input.Value.(string)
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return nil, err
	}
	return t.Format("15:04:05"), nil
}
func parseDateTime(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil {
		return nil, nil
	}
	timeStr := input.Value.(string)
	return time.Parse("2006-01-02 15:04:05", timeStr)
}
