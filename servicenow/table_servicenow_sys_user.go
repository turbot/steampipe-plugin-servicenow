package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSysUser() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_user",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowSysUsers,
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "first_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "", Type: proto.ColumnType_STRING},
			{Name: "department", Description: "", Type: proto.ColumnType_STRING},
			{Name: "introduction", Description: "", Type: proto.ColumnType_STRING},
			{Name: "calendar_integration", Description: "", Type: proto.ColumnType_STRING},
			{Name: "country", Description: "", Type: proto.ColumnType_STRING},
			{Name: "user_password", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_login_time", Description: "", Type: proto.ColumnType_STRING},
			{Name: "source", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "building", Description: "", Type: proto.ColumnType_STRING},
			{Name: "web_service_access_only", Description: "", Type: proto.ColumnType_STRING},
			{Name: "notification", Description: "", Type: proto.ColumnType_STRING},
			{Name: "enable_multifactor_authn", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "state", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vip", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "zip", Description: "", Type: proto.ColumnType_STRING},
			{Name: "home_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_format", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_login", Description: "", Type: proto.ColumnType_STRING},
			{Name: "default_perspective", Description: "", Type: proto.ColumnType_STRING},
			{Name: "active", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain_path", Description: "", Type: proto.ColumnType_STRING},
			{Name: "cost_center", Description: "", Type: proto.ColumnType_STRING},
			{Name: "phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "employee_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "password_needs_reset", Description: "", Type: proto.ColumnType_STRING},
			{Name: "gender", Description: "", Type: proto.ColumnType_STRING},
			{Name: "city", Description: "", Type: proto.ColumnType_STRING},
			{Name: "failed_attempts", Description: "", Type: proto.ColumnType_STRING},
			{Name: "user_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "roles", Description: "", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_class_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "internal_integration_user", Description: "", Type: proto.ColumnType_STRING},
			{Name: "ldap_server", Description: "", Type: proto.ColumnType_STRING},
			{Name: "mobile_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "street", Description: "", Type: proto.ColumnType_STRING},
			{Name: "preferred_language", Description: "", Type: proto.ColumnType_STRING},
			{Name: "manager", Description: "", Type: proto.ColumnType_STRING},
			{Name: "locked_out", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "photo", Description: "", Type: proto.ColumnType_STRING},
			{Name: "avatar", Description: "", Type: proto.ColumnType_STRING},
			{Name: "middle_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_tags", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_zone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "schedule", Description: "", Type: proto.ColumnType_STRING},
			{Name: "date_format", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listServicenowSysUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sys_user.listServicenowSysUsers", "connect_error", err)
		return nil, err
	}

	sysUsersResponse, err := client.GetSysUsers(10)
	if err != nil {
		logger.Error("servicenow_sys_user.listServicenowSysUsers", "query_error", err)
		return nil, err
	}
	for _, sysUser := range sysUsersResponse.Result {
		d.StreamListItem(ctx, sysUser)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
