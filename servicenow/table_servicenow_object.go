package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// LIST HYDRATE FUNCTION

type tableListResult struct {
	Result []map[string]interface{} `json:"result"`
}

func listServicenowObjectsByTable(tableName string, servicenowCols map[string]string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		logger := plugin.Logger(ctx)
		client, err := ConnectUncached(ctx, d.Connection)
		if err != nil {
			logger.Error("servicenow.listServicenowObjectsByTable", "connect_error", err)
			return nil, err
		}

		var response tableListResult
		err = client.NowTable.List(tableName, 10, 0, "", &response)
		if err != nil {
			logger.Error("servicenow_incident.listServicenowIncidents", "query_error", err)
			return nil, err
		}

		for _, element := range response.Result {
			// Check if the value is an empty string, if it is, replace it with nil
			for key, value := range element {
				if str, ok := value.(string); ok && str == "" {
					element[key] = nil
				}
			}

			d.StreamListItem(ctx, element)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		return nil, nil
	}
}

//// GET HYDRATE FUNCTION

func getServicenowObjectbyID(tableName string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		// plugin.Logger(ctx).Info("servicenow.getServicenowObjectbyID", "Table_Name", d.Table.Name)
		// id := d.KeyColumnQualString("id")
		// if strings.TrimSpace(id) == "" {
		// 	return nil, nil
		// }

		// client, err := connect(ctx, d)
		// if err != nil {
		// 	plugin.Logger(ctx).Error("servicenow.getServicenowObjectbyID", "connection error", err)
		// 	return nil, err
		// }
		// if client == nil {
		// 	plugin.Logger(ctx).Error("servicenow.getServicenowObjectbyID", "client_not_found: unable to generate dynamic tables because of invalid steampipe servicenow configuration", err)
		// 	return nil, fmt.Errorf("servicenow.getServicenowObjectbyID: client_not_found, unable to query table %s because of invalid steampipe servicenow configuration", d.Table.Name)
		// }

		// obj := client.SObject(tableName).Get(id)
		// if obj == nil {
		// 	// Object doesn't exist, handle the error
		// 	plugin.Logger(ctx).Warn("servicenow.getServicenowObjectbyID", fmt.Sprintf("%s with id \"%s\" not found", tableName, id))
		// 	return nil, nil
		// }

		// object := new(map[string]interface{})
		// err = decodeQueryResult(ctx, obj, object)
		// if err != nil {
		// 	plugin.Logger(ctx).Error("servicenow.getServicenowObjectbyID", "result decoding error", err)
		// 	return nil, err
		// }

		// return *object, nil
		return nil, nil
	}
}

//// TRANSFORM FUNCTION

func getFieldFromSObjectMap(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	param := d.Param.(string)
	ls := d.HydrateItem.(map[string]interface{})
	return ls[param], nil
}

func getFieldFromSObjectMapByColumnName(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	servicenowColumnName := getServicenowColumnName(d.ColumnName)
	ls := d.HydrateItem.(map[string]interface{})
	return ls[servicenowColumnName], nil
}
