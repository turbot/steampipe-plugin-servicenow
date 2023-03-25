package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowNowContact() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_now_contact",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowNowContacts,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowNowContacts,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "account", Description: "", Type: proto.ColumnType_STRING},
			{Name: "active", Description: "", Type: proto.ColumnType_STRING},
			{Name: "agent_status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "building", Description: "", Type: proto.ColumnType_STRING},
			{Name: "calendar_integration", Description: "", Type: proto.ColumnType_STRING},
			{Name: "city", Description: "", Type: proto.ColumnType_STRING},
			{Name: "company", Description: "", Type: proto.ColumnType_STRING},
			{Name: "cost_center", Description: "", Type: proto.ColumnType_STRING},
			{Name: "country", Description: "", Type: proto.ColumnType_STRING},
			{Name: "date_format", Description: "", Type: proto.ColumnType_STRING},
			{Name: "default_perspective", Description: "", Type: proto.ColumnType_STRING},
			{Name: "department", Description: "", Type: proto.ColumnType_STRING},
			{Name: "edu_status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "", Type: proto.ColumnType_STRING},
			{Name: "employee_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "enable_multifactor_authn", Description: "", Type: proto.ColumnType_STRING},
			{Name: "failed_attempts", Description: "", Type: proto.ColumnType_STRING},
			{Name: "first_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "gender", Description: "", Type: proto.ColumnType_STRING},
			{Name: "geolocation_tracked", Description: "", Type: proto.ColumnType_STRING},
			{Name: "home_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "internal_integration_user", Description: "", Type: proto.ColumnType_STRING},
			{Name: "introduction", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_login_device", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_login_time", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_login", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_position_update", Description: "", Type: proto.ColumnType_STRING},
			{Name: "latitude", Description: "", Type: proto.ColumnType_STRING},
			{Name: "ldap_server", Description: "", Type: proto.ColumnType_STRING},
			{Name: "location", Description: "", Type: proto.ColumnType_STRING},
			{Name: "locked_out", Description: "", Type: proto.ColumnType_STRING},
			{Name: "longitude", Description: "", Type: proto.ColumnType_STRING},
			{Name: "manager", Description: "", Type: proto.ColumnType_STRING},
			{Name: "middle_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "mobile_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "notification", Description: "", Type: proto.ColumnType_STRING},
			{Name: "on_schedule", Description: "", Type: proto.ColumnType_STRING},
			{Name: "phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "photo", Description: "", Type: proto.ColumnType_STRING},
			{Name: "preferred_language", Description: "", Type: proto.ColumnType_STRING},
			{Name: "roles", Description: "", Type: proto.ColumnType_STRING},
			{Name: "schedule", Description: "", Type: proto.ColumnType_STRING},
			{Name: "source", Description: "", Type: proto.ColumnType_STRING},
			{Name: "state", Description: "", Type: proto.ColumnType_STRING},
			{Name: "street", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_class_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain_path", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_tags", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_format", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_sheet_policy", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_zone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "", Type: proto.ColumnType_STRING},
			{Name: "user_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vip", Description: "", Type: proto.ColumnType_STRING},
			{Name: "web_service_access_only", Description: "", Type: proto.ColumnType_STRING},
			{Name: "zip", Description: "", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listServicenowNowContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_now_contact.listServicenowNowContacts", "connect_error", err)
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
		returnedObject, err := client.NowContact.List(limit, offset)
		totalReturned := len(returnedObject.Result)
		if err != nil {
			logger.Error("servicenow_now_contact.listServicenowNowContacts", "query_error", err)
			return nil, err
		}
		for _, object := range returnedObject.Result {
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

func getServicenowNowContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_now_contact.getServicenowNowContacts", "connect_error", err)
		return nil, err
	}

	sysId := d.EqualsQualString("sys_id")

	returnedObject, err := client.NowContact.Read(sysId)
	if err != nil {
		logger.Error("servicenow_now_contact.getServicenowNowContacts", "query_error", err)
		return nil, err
	}

	return returnedObject.Result, err
}
