package servicenow

import (
	"context"
	"fmt"
	"regexp"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/turbot/go-servicenow/servicenow"
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

	staticTables := []string{}

	dynamicColumnsMap := map[string]dynamicMap{}
	var mapLock sync.Mutex

	// If Servicenow client was obtained, don't generate dynamic columns for
	// defined static tables
	if client != nil {
		var wgd sync.WaitGroup
		wgd.Add(len(staticTables))
		for _, st := range staticTables {
			go func(staticTable string) {
				defer wgd.Done()
				dynamicCols, dynamicKeyColumns, servicenowCols := dynamicColumns(ctx, client, staticTable, nil)
				// dynamicCols, dynamicKeyColumns, servicenowCols := dynamicColumns(ctx, client, staticTable, p)
				mapLock.Lock()
				dynamicColumnsMap[staticTable] = dynamicMap{dynamicCols, dynamicKeyColumns, servicenowCols}
				defer mapLock.Unlock()
			}(st)
		}
		wgd.Wait()
	}

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

	// // Get the query for the metric (required)
	// servicenowTableName := ctx.Value(contextKey("ServicenowTableName")).(string)
	// tableName := ctx.Value(contextKey("PluginTableName")).(string)

	// // Top columns
	// cols := []*plugin.Column{}
	// servicenowCols := map[string]string{}
	// // Key columns
	// keyColumns := plugin.KeyColumnSlice{}

	servicenowObjectFields, err := getTableColumns(client)
	if err != nil {
		logger.Error("servicenow.generateDynamicTables", "connection_error", err)
	}
	for k, v := range servicenowObjectFields {
		logger.Warn(k, v)
	}

	// for fieldName, fieldType := range servicenowObjectFields {
	// 	// Column dynamic generation
	// 	// Don't convert to snake case since field names can have underscores in
	// 	// them, so it's impossible to convert from snake case back to camel case
	// 	// to match the original field name. Also, if we convert to snake case,
	// 	// custom fields like "TestField" and "Test_Field" will result in duplicates
	// 	var columnFieldName string
	// 	if strings.HasSuffix(fieldName, "__c") {
	// 		columnFieldName = strings.ToLower(fieldName)
	// 	} else {
	// 		columnFieldName = strcase.ToSnake(fieldName)
	// 	}

	// 	column := plugin.Column{
	// 		Name:        columnFieldName,
	// 		Description: fmt.Sprintf("%s.", fieldName),
	// 		Transform:   transform.FromP(getFieldFromSObjectMap, fieldName),
	// 	}
	// 	// Adding column type in the map to help in qual handling
	// 	servicenowCols[columnFieldName] = fieldType

	// 	// Set column type based on the `soapType` from servicenow schema
	// 	switch fieldType {
	// 	case "string", "ID":
	// 		column.Type = proto.ColumnType_STRING
	// 		keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
	// 	case "date", "dateTime":
	// 		column.Type = proto.ColumnType_TIMESTAMP
	// 		keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
	// 	case "boolean":
	// 		column.Type = proto.ColumnType_BOOL
	// 		keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", "<>"}})
	// 	case "double":
	// 		column.Type = proto.ColumnType_DOUBLE
	// 		keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
	// 	case "int":
	// 		column.Type = proto.ColumnType_INT
	// 		keyColumns = append(keyColumns, &plugin.KeyColumn{Name: columnFieldName, Require: plugin.Optional, Operators: []string{"=", ">", ">=", "<=", "<"}})
	// 	default:
	// 		column.Type = proto.ColumnType_JSON
	// 	}
	// 	cols = append(cols, &column)
	// }

	// Table := plugin.Table{
	// 	Name: tableName,
	// 	// Description: fmt.Sprintf("Represents Servicenow object %s.", servicenowObjectMetadata["name"]),
	// 	List: &plugin.ListConfig{
	// 		KeyColumns: keyColumns,
	// 		Hydrate:    listServicenowObjectsByTable(servicenowTableName, servicenowCols),
	// 	},
	// 	Get: &plugin.GetConfig{
	// 		KeyColumns: plugin.SingleColumn("id"),
	// 		Hydrate:    getServicenowObjectbyID(servicenowTableName),
	// 	},
	// 	Columns: cols,
	// }

	// return &Table
	return nil
}

func getTableColumns(client *servicenow.ServiceNow) (map[string]string, error) {
	columns := map[string]string{}
	limit := 1000
	offset := 0
	for {
		var returned model.SysDictionaryListResult
		err := client.NowTable.List(model.SysDictionaryTableName, limit, offset, "name=incident", &returned)
		if err != nil {
			return nil, err
		}
		totalReturned := len(returned.Result)
		for _, returnedObject := range returned.Result {
			if returnedObject.Element == "" {
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
