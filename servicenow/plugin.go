package servicenow

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/turbot/go-servicenow/servicenow"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-servicenow/model"
)

const pluginName = "steampipe-plugin-servicenow"

type contextKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.From(getFieldFromSObjectMapByColumnName).NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: pluginTableDefinitions,
	}

	return p
}

type dynamicMap struct {
	cols              []*plugin.Column
	keyColumns        plugin.KeyColumnSlice
	servicenowColumns map[string]string
}

func pluginTableDefinitions(ctx context.Context, d *plugin.TableMapData) (map[string]*plugin.Table, error) {
	// If unable to connect to servicenow instance, log warning and abort dynamic table creation
	client, err := ConnectUncached(ctx, d.Connection)
	if err != nil {
		// do not abort the plugin as static table needs to be generated
		plugin.Logger(ctx).Warn("servicenow.pluginTableDefinitions", "connection_error: unable to generate dynamic tables because of invalid steampipe servicenow configuration", err)
	}

	// staticTables := []string{}

	// dynamicColumnsMap := map[string]dynamicMap{}
	// var mapLock sync.Mutex

	// // If Servicenow client was obtained, don't generate dynamic columns for
	// // defined static tables
	// if client != nil {
	// 	var wgd sync.WaitGroup
	// 	wgd.Add(len(staticTables))
	// 	for _, st := range staticTables {
	// 		go func(staticTable string) {
	// 			defer wgd.Done()
	// 			dynamicCols, dynamicKeyColumns, servicenowCols := dynamicColumns(ctx, client, staticTable, nil)
	// 			// dynamicCols, dynamicKeyColumns, servicenowCols := dynamicColumns(ctx, client, staticTable, p)
	// 			mapLock.Lock()
	// 			dynamicColumnsMap[staticTable] = dynamicMap{dynamicCols, dynamicKeyColumns, servicenowCols}
	// 			defer mapLock.Unlock()
	// 		}(st)
	// 	}
	// 	wgd.Wait()
	// }

	// Initialize tables with static tables with static and dynamic columns(if credentials are set)
	tables := map[string]*plugin.Table{}

	var re = regexp.MustCompile(`\d+`)
	var substitution = ``
	servicenowTables := []string{}
	config := GetConfig(d.Connection)
	if config.Objects != nil && len(*config.Objects) > 0 {
		for _, tableName := range *config.Objects {
			pluginTableName := "servicenow_" + strcase.ToSnake(re.ReplaceAllString(tableName, substitution))
			if _, ok := tables[pluginTableName]; !ok {
				servicenowTables = append(servicenowTables, tableName)
			}
		}
	}

	if client == nil {
		plugin.Logger(ctx).Warn("servicenow.pluginTableDefinitions", "client_not_found: unable to generate dynamic tables because of invalid steampipe servicenow configuration", err)
		return tables, nil
	}

	var wg sync.WaitGroup
	wg.Add(len(servicenowTables))
	for _, sfTable := range servicenowTables {
		tableName := "servicenow_" + strcase.ToSnake(re.ReplaceAllString(sfTable, substitution))
		if tables[tableName] != nil {
			wg.Done()
			continue
		}
		go func(name string) {
			defer wg.Done()
			tableName := "servicenow_" + strcase.ToSnake(re.ReplaceAllString(name, substitution))
			plugin.Logger(ctx).Debug("servicenow.pluginTableDefinitions", "object_name", name, "table_name", tableName)
			ctx = context.WithValue(ctx, contextKey("PluginTableName"), tableName)
			ctx = context.WithValue(ctx, contextKey("ServicenowTableName"), name)
			table := generateDynamicTables(ctx, d)
			// Ignore if the requested Servicenow object is not present.
			if table != nil {
				tables[tableName] = table
			}
		}(sfTable)
	}
	wg.Wait()
	return tables, nil
}

func generateDynamicTables(ctx context.Context, d *plugin.TableMapData) *plugin.Table {
	logger := plugin.Logger(ctx)
	logger.Warn("generateDynamicTables")
	client, err := ConnectUncached(ctx, d.Connection)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "connection_error", err)
		return nil
	}

	// Get the query for the metric (required)
	servicenowTableName := ctx.Value(contextKey("ServicenowTableName")).(string)
	tableName := ctx.Value(contextKey("PluginTableName")).(string)

	// Top columns
	cols := []*plugin.Column{}
	servicenowCols := map[string]string{}
	// Key columns
	keyColumns := plugin.KeyColumnSlice{}

	servicenowObjectFields, err := getTableColumns(client, servicenowTableName)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "column_build_error", err)
	}

	fieldsDescriptions, err := getTableColumnsDescriptions(client, servicenowTableName)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "column_documentation_error", err)
	}

	for fieldName, fieldType := range servicenowObjectFields {
		columnFieldName := fieldName

		columnDescription := fieldsDescriptions[columnFieldName].Hint
		if columnDescription == "" {
			columnDescription = fieldsDescriptions[columnFieldName].Label
		}
		column := plugin.Column{
			Name:        columnFieldName,
			Description: fmt.Sprintf("%s.", columnDescription),
			Transform:   transform.FromP(getFieldFromSObjectMap, fieldName),
		}
		// Adding column type in the map to help in qual handling
		servicenowCols[columnFieldName] = fieldType

		// Set column type based on the `soapType` from servicenow schema
		switch fieldType {
		case "string", "glide_date", "date", "time":
			column.Type = proto.ColumnType_STRING
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
		case "glide_time":
			column.Type = proto.ColumnType_STRING
			column.Transform.Transform(parseGlideTime)
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "datetime":
			column.Type = proto.ColumnType_TIMESTAMP
			column.Transform.Transform(parseDateTime)
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "boolean":
			column.Type = proto.ColumnType_BOOL
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
		case "double", "decimal", "float":
			column.Type = proto.ColumnType_DOUBLE
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		case "int", "integer", "longint":
			column.Type = proto.ColumnType_INT
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
		default:
			column.Type = proto.ColumnType_JSON
		}

		cols = append(cols, &column)
	}

	serviceNowTableObject, err := getTableObject(client, servicenowTableName)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "table_documentation_error", err)
		return nil
	}

	Table := plugin.Table{
		Name: tableName,
		// Description: fmt.Sprintf("Represents Servicenow object %s.", servicenowObjectMetadata["name"]),
		Description: fmt.Sprintf("%s.", serviceNowTableObject.Label),
		List: &plugin.ListConfig{
			// KeyColumns: keyColumns,
			Hydrate: listServicenowObjectsByTable(servicenowTableName, servicenowCols),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("sys_id"),
			Hydrate:    getServicenowObjectbyID(servicenowTableName),
		},
		Columns: cols,
		// Columns: []*plugin.Column{
		// 	{Name: "raw", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromValue()},
		// },
	}

	return &Table
}

func getTableObject(client *servicenow.ServiceNow, tableName string) (*model.SysDbObject, error) {
	var returned model.SysDbObjectListResult
	err := client.NowTable.List(model.SysDbObjectTableName, 1, 0, fmt.Sprintf("name=%s", tableName), &returned)
	if err != nil {
		return nil, err
	}
	if len(returned.Result) == 0 {
		return nil, fmt.Errorf("Table %s not found on ServiceNow.", tableName)
	}
	return &returned.Result[0], nil
}

func getTableColumnsDescriptions(client *servicenow.ServiceNow, tableName string) (map[string]model.SysDocumentation, error) {
	columnsDescriptions := map[string]model.SysDocumentation{}
	limit := 1000
	offset := 0
	for {
		var returned model.SysDocumentationListResult
		err := client.NowTable.List(model.SysDocumentationTableName, limit, offset, fmt.Sprintf("name=%s", tableName), &returned)
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

func getTableColumns(client *servicenow.ServiceNow, tableName string) (map[string]string, error) {
	columns := map[string]string{}
	limit := 1000
	offset := 0
	for {
		var returned model.SysDictionaryListResult
		err := client.NowTable.List(model.SysDictionaryTableName, limit, offset, fmt.Sprintf("name=%s", tableName), &returned)
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
			var typeGlide model.SysGlideObjectListResult
			err := client.NowTable.List(model.SysGlideObjectTableName, 1, 0, fmt.Sprintf("name=%s", returnedObject.InternalType.Value), &typeGlide)
			if err != nil {
				return nil, err
			}
			columns[returnedObject.Element] = typeGlide.Result[0].ScalarType
		}

		if totalReturned < limit {
			break
		}
		offset += limit
	}

	return columns, nil
}

func parseGlideTime(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil {
		return nil, nil
	}
	timeStr := input.Value.(string)
	t, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		return nil, err
	}
	return t.Format(time.TimeOnly), nil
}
func parseDateTime(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil {
		return nil, nil
	}
	timeStr := input.Value.(string)
	return time.Parse(time.DateTime, timeStr)
}
