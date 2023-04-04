package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type tableListResult struct {
	Result []map[string]interface{} `json:"result"`
}

type tableGetResult struct {
	Result map[string]interface{} `json:"result"`
}

//// LIST HYDRATE FUNCTION

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
		logger := plugin.Logger(ctx)
		logger.Info("servicenow.getServicenowObjectbyID", "Table_Name", d.Table.Name)
		sysId := d.EqualsQualString("sys_id")

		client, err := Connect(ctx, d)
		if err != nil {
			logger.Error("servicenow.getServicenowObjectbyID", "connect_error", err)
			return nil, err
		}

		var response tableGetResult
		err = client.NowTable.Read(tableName, sysId, &response)
		if err != nil {
			logger.Error("servicenow_incident.getServicenowIncident", "query_error", err)
			return nil, err
		}

		sanitizeTableObject(response.Result)
		return response.Result, nil
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

func sanitizeTableObject(tableObject map[string]interface{}) {
	// Check if the value is an empty string, if it is, replace it with nil
	for key, value := range tableObject {
		if str, ok := value.(string); ok && str == "" {
			tableObject[key] = nil
		}
	}
}
