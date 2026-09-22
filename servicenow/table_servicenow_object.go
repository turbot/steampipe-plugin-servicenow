package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

type tableListResult struct {
	Result []map[string]interface{} `json:"result"`
}

type tableGetResult struct {
	Result map[string]interface{} `json:"result"`
}

// snowPageSize is how many rows are requested per API call when paging through a table.
const snowPageSize = 30

// pageSizeFor returns the number of rows to ask the API for. Only the first page is trimmed to the
// SQL limit, which is all an exactly filtered query needs. Reaching a second page means rows are
// being dropped, either by the qual recheck or by Postgres, and a page the size of the limit would
// then walk the rest of the table a few rows per request.
func pageSizeFor(limit *int64, firstPage bool) int {
	if !firstPage || limit == nil || *limit >= snowPageSize {
		return snowPageSize
	}
	// ServiceNow answers 400 to a sysparm_limit of zero or less, which would fail the query
	// rather than return nothing, so never ask for fewer than one row.
	if *limit < 1 {
		return 1
	}
	return int(*limit)
}

//// LIST HYDRATE FUNCTION

func listServicenowObjectsByTable(tableName string, servicenowCols map[string]string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		logger := plugin.Logger(ctx)
		client, err := Connect(ctx, d)
		if err != nil {
			logger.Error("servicenow.listServicenowObjectsByTable", "connect_error", err)
			return nil, err
		}

		query := buildQueryFromQuals(d.Quals, d.Table.Columns, servicenowCols)
		if query != "" {
			plugin.Logger(ctx).Debug("servicenow.listServicenowObjectsByTable", "table_name", d.Table.Name, "query_condition", query)
		}

		offset := 0
		firstPage := true

		for {
			pageSize := pageSizeFor(d.QueryContext.Limit, firstPage)

			var response tableListResult
			err = client.NowTable.List(tableName, pageSize, offset, query, false, &response)
			if err != nil {
				logger.Error("servicenow.listServicenowObjectsByTable", "query_error", err)
				return nil, err
			}
			totalReturned := len(response.Result)

			for _, element := range response.Result {
				sanitizeTableObject(element)

				// Skip rows that don't match the exact quals, so they don't count towards the limit
				if !rowMatchesQuals(element, d.Quals, d.Table.Columns) {
					continue
				}

				d.StreamListItem(ctx, element)
				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}

			if totalReturned < pageSize {
				break
			}
			offset += totalReturned
			firstPage = false
		}
		return nil, err
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
		err = client.NowTable.Get(tableName, sysId, false, &response)
		if err != nil {
			logger.Error("servicenow.getServicenowObjectbyID", "query_error", err)
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
	ls := d.HydrateItem.(map[string]interface{})
	return ls[d.ColumnName], nil
}

func sanitizeTableObject(tableObject map[string]interface{}) {
	// Check if the value is an empty string, if it is, replace it with nil
	for key, value := range tableObject {
		if str, ok := value.(string); ok && str == "" {
			tableObject[key] = nil
		}
	}
}
