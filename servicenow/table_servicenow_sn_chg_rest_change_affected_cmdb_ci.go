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
		List: &plugin.ListConfig{
			Hydrate:       listServicenowSnChgRestChangeAffectedCmdbCis,
			ParentHydrate: listServicenowSnChgRestChanges,
		},
		Columns: []*plugin.Column{
			{Name: "applied_date", Description: "Date when the change was applied to the CMDB CI item.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "applied_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "ci_item_sys_id", Description: "System ID of the CMDB CI item affected by the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "ci_item").NullIfEqual("")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_tags", Description: "Tags associated with the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_tags").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "task_name", Description: "Name of the task associated with the affected CMDB CI item.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "task").NullIfEqual("")},
			{Name: "task", Description: "Task associated with the affected CMDB CI item.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task").NullIfEqual("")},
			{Name: "applied", Description: "Indicates whether the change was applied to the CI item.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "applied").NullIfEqual("")},
			{Name: "ci_item_name", Description: "Name of the CI item affected by the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "ci_item").NullIfEqual("")},
			{Name: "manual_proposed_change", Description: "Indicates whether the proposed change was manual.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "manual_proposed_change").NullIfEqual("")},
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
