package servicenow

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func ignoreError(errors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range errors {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}

func tableServicenowSnChgRestChangeSchedule() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change_schedule",
		Description:      "Change Management - Change Schedule.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: ignoreError([]string{"response 400", "400"}),
		},
		List: &plugin.ListConfig{
			Hydrate:       listServicenowSnChgRestChangeSchedules,
			ParentHydrate: listServicenowSnChgRestChanges,
		},
		Columns: []*plugin.Column{
			{Name: "change_sys_id", Description: "Unique identifier of the change request associated to the conflict.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_sys_id").NullIfEqual("")},
			{Name: "state_description", Description: "Information on the current state of the worker.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "state.display_value").NullIfEqual("")},
			{Name: "state", Description: "Numeric information on the current state of the worker.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "state.value").NullIfEqual("")},
			{Name: "worker_link", Description: "	Indicates the type of request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "worker.link").NullIfEqual("")},
			{Name: "worker_sys_id", Description: "Unique identifier of the change request associated to the conflict.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "worker.sysId").NullIfEqual("")},
			// {Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChangeSchedules(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	changes := h.Item.(map[string]map[string]interface{})
	if changes["sys_id"]["value"] == nil {
		return nil, nil
	}
	change_sys_id := changes["sys_id"]["value"]

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_schedule.listServicenowSnChgRestChangeSchedules", "connect_error", err)
		return nil, err
	}

	var response map[string]interface{}
	err = client.SnChgRestChangeSchedule.Get(change_sys_id.(string), &response)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_schedule.listServicenowSnChgRestChangeSchedules", "query_error", err)
		return nil, err
	}
	if response["error"] != nil {
		return nil, nil
	}
	response["change_sys_id"] = change_sys_id
	d.StreamListItem(ctx, response)
	return nil, err
}
