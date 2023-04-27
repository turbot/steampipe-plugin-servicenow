package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnChgRestChangeAffectedCmdbCi() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change_affected_cmdb_ci",
		Description:      "Change Management - CMDB CI affected by change.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: ignoreError([]string{"response 400", "400"}),
		},
		List: &plugin.ListConfig{
			Hydrate:       listServicenowSnChgRestChangeAffectedCmdbCis,
			ParentHydrate: listServicenowSnChgRestChanges,
		},
		Columns: []*plugin.Column{
			{Name: "change_sys_id", Description: "Unique identifier of the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task").NullIfEqual("")},
			{Name: "cmdb_ci_sys_id", Description: "Unique identifier of the affected CMDB CI.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cmdb_ci_item").NullIfEqual("")},
			{Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChangeAffectedCmdbCis(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	changes := h.Item.(map[string]map[string]interface{})
	if changes["sys_id"]["value"] == nil {
		return nil, nil
	}
	change_sys_id := changes["sys_id"]["value"]

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_affected_cmdb_ci.listServicenowSnChgRestChangeAffectedCmdbCis", "connect_error", err)
		return nil, err
	}

	offset := 0
	limit := 30
	if d.QueryContext.Limit != nil {
		pgLimit := int(*d.QueryContext.Limit)
		if pgLimit < limit {
			limit = pgLimit
		}
	}

	for {
		var response snChgRestChangeListResult
		err := client.SnChgRestChangeCi.List(change_sys_id.(string), "affected", limit, offset, "", &response)
		totalReturned := len(response.Result)
		if err != nil {
			logger.Error("servicenow_sn_chg_rest_change_affected_cmdb_ci.listServicenowSnChgRestChangeAffectedCmdbCis", "query_error", err)
			return nil, err
		}
		for _, object := range response.Result {
			d.StreamListItem(ctx, object)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if totalReturned < limit {
			break
		}
		offset += limit
	}
	return nil, err
}
