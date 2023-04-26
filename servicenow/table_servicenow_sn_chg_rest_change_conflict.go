package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnChgRestChangeConflict() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change_conflict",
		Description:      "Change Management - Change Conflict.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate:       listServicenowSnChgRestChangeConflicts,
			ParentHydrate: listServicenowSnChgRestChanges,
		},
		Columns: []*plugin.Column{
			{Name: "change_sys_id", Description: "Unique identifier of the change request associated to the conflict.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_sys_id").NullIfEqual("")},
			{Name: "conflicts", Description: "List of conflicts found for the change request.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "conflicts").NullIfEqual("")},
			{Name: "job_status", Description: "Status of the actual conflict checking job.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "job_status").NullIfEqual("")},
			{Name: "last_run", Description: "Date and time the last conflict checking process started.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "last_run").NullIfEqual("")},
			{Name: "record_count", Description: "", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "record_count").NullIfEqual("")},
			{Name: "status", Description: "Outcome of the conflict checking process.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "status").NullIfEqual("")},

			{Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChangeConflicts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	changes := h.Item.(map[string]map[string]interface{})
	if changes["sys_id"]["value"] == nil {
		return nil, nil
	}
	change_sys_id := changes["sys_id"]["value"]

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_conflict.listServicenowSnChgRestChangeConflicts", "connect_error", err)
		return nil, err
	}

	var response map[string]map[string]interface{}
	err = client.SnChgRestChangeConflict.Get(change_sys_id.(string), &response)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_conflict.listServicenowSnChgRestChangeConflicts", "query_error", err)
		return nil, err
	}
	result := response["result"]
	result["change_sys_id"] = change_sys_id
	d.StreamListItem(ctx, result)
	return nil, err
}
