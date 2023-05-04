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
			{Name: "applied_date", Description: "Applied Date.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "applied_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "ci_item_sys_id", Description: "Unique identifier of the impacted CMDB CI service..", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "ci_item").NullIfEqual("")},
			{Name: "sys_created_by", Description: "Sys Created By.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Sys Created On.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys Id.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Sys Mod Count.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_tags", Description: "Sys Tags.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_tags").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "Sys Updated By.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Sys Updated On.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "task_name", Description: "Task Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "task").NullIfEqual("")},
			{Name: "task", Description: "Task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task").NullIfEqual("")},
			{Name: "applied", Description: "Applied.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "applied").NullIfEqual("")},
			{Name: "ci_item_name", Description: "Cmdb Ci Service Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "ci_item").NullIfEqual("")},
			{Name: "manual_proposed_change", Description: "Manual Proposed Change.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "manual_proposed_change").NullIfEqual("")},
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
