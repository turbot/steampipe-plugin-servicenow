package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnChgRestChangeModel() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change_model",
		Description:      "Change Management - Change Model.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate: listServicenowSnChgRestChangeModels,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowSnChgRestChangeModels,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "sys_id", Description: "Unique identifier of the associated change model record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "active", Description: "Flag that indicates whether the associated change model record is active and available within the instance.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "active").NullIfEqual("")},
			{Name: "available_in_ui", Description: "Flag that indicates whether the associated change model record is available within the user interface.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "available_in_ui").NullIfEqual("")},
			{Name: "color", Description: "Color of the associated change model on the change request landing page.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "color").NullIfEqual("")},
			{Name: "default_change_model", Description: "Flag that indicates whether the associated change model record is the default change model.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "default_change_model").NullIfEqual("")},
			{Name: "description", Description: "Short description of the purpose of the change model.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "description").NullIfEqual("")},
			{Name: "name", Description: "Name of the change model.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "name").NullIfEqual("")},
			{Name: "record_preset", Description: "Name-value pairs of the fields that should automatically be populated, with their associated values, when a new change request record is created. Values are separated by caret symbols.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "record_preset").NullIfEqual("")},
			{Name: "state_field", Description: "Choice list field from which to collect choices, based on the provided in table_name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "state_field").NullIfEqual("")},
			{Name: "sys_class_name", Description: "Change module table name. Always Change Model/chg_model.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_class_name").NullIfEqual("")},
			{Name: "sys_created_by", Description: "Name of the user that initially created the associated change module record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Date and time that the change module record was originally created.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("")},
			{Name: "sys_domain", Description: "If using domains in the instance, the name of the domain to which the change module record is associated.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain").NullIfEqual("")},
			{Name: "sys_domain_path", Description: "If using domains in the instance, the domain path in which the associated change module record resides.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain_path").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Number of times that the associated change model record has been modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_name", Description: "Name of the change model. Always the same as the", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_name").NullIfEqual("")},
			{Name: "sys_tags", Description: "System tags associated with the change model record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_tags").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "Name of the user that last updated the associated change model record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Date and time the associated change model record was last updated.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("")},
			{Name: "table_name", Description: "Table that defines the Choice list field from which to collect choices. For change models this is always set to 'change_request'.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "table_name").NullIfEqual("")},
		}),
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChangeModels(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_model.listServicenowSnChgRestChangeModels", "connect_error", err)
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
		err := client.SnChgRestChangeModel.List(limit, offset, "", &response)
		totalReturned := len(response.Result)
		if err != nil {
			logger.Error("servicenow_sn_chg_rest_change_model.listServicenowSnChgRestChangeModels", "query_error", err)
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

//// GET FUNCTION

func getServicenowSnChgRestChangeModels(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_model.getServicenowSnChgRestChangeModels", "connect_error", err)
		return nil, err
	}

	sysId := d.EqualsQualString("sys_id")

	var response snChgRestChangeGetResult
	err = client.SnChgRestChangeModel.Get(sysId, &response)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_model.getServicenowSnChgRestChangeModels", "query_error", err)
		return nil, err
	}

	return response.Result, err
}
