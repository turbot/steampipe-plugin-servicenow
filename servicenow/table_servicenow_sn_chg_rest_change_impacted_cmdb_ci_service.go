package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnChgRestChangeImpactedCmdbCiService() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change_impacted_cmdb_ci_service",
		Description:      "Change Management - CMDB CI service impacted by change.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate:       listServicenowSnChgRestChangeImpactedCmdbCiServices,
			ParentHydrate: listServicenowSnChgRestChanges,
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "cmdb_ci_service_name", Description: "Name of the CMDB CI service impacted by the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "cmdb_ci_service").NullIfEqual("")},
			{Name: "cmdb_ci_service_sys_id", Description: "System ID of the CMDB CI service impacted by the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cmdb_ci_service").NullIfEqual("")},
			{Name: "manually_added", Description: "Indicates whether the CMDB CI service was manually added to the change.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "manually_added").NullIfEqual("")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_tags", Description: "Tags associated with the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_tags").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "task_name", Description: "Name of the task associated with the impacted CMDB CI service.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "task").NullIfEqual("")},
			{Name: "task", Description: "Task associated with the impacted CMDB CI service.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task").NullIfEqual("")},
		}),
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChangeImpactedCmdbCiServices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	changes := h.Item.(map[string]map[string]interface{})
	if changes["sys_id"]["value"] == nil {
		return nil, nil
	}
	change_sys_id := changes["sys_id"]["value"]

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_impacted_cmdb_ci_service.listServicenowSnChgRestChangeImpactedCmdbCiServices", "connect_error", err)
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
		err := client.SnChgRestChangeCi.List(change_sys_id.(string), "impacted", limit, offset, "", &response)
		totalReturned := len(response.Result)
		if err != nil {
			logger.Error("servicenow_sn_chg_rest_change_impacted_cmdb_ci_service.listServicenowSnChgRestChangeImpactedCmdbCiServices", "query_error", err)
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
