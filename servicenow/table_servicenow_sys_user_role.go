package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysUserRoleTableName = "sys_user_role"

//// TABLE DEFINITION

func tableServicenowSysUserRole() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_user_role",
		Description: "User Role.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"description", "includes_roles", "name", "requires_subscription", "suffix", "sys_class_name", "sys_created_by", "sys_id", "sys_name", "sys_policy", "sys_update_name", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysUserRoleTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserRoleTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "assignable_by", Description: "If the related application has Scoped administration enabled, a user needs to have the 'Assignable By' role to be able to assign this role to another user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "assignable_by")},
			{Name: "can_delegate", Description: "Can be delegated.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "can_delegate")},
			{Name: "description", Description: "Description.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "description")},
			{Name: "elevated_privilege", Description: "This role is an elevated privilege.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "elevated_privilege")},
			{Name: "grantable", Description: "Can be granted independently.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "grantable")},
			{Name: "includes_roles", Description: "Includes roles.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "includes_roles")},
			{Name: "name", Description: "Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "requires_subscription", Description: "Select Yes, if this role enables activity that requires a subscription.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "requires_subscription")},
			{Name: "scoped_admin", Description: "If this is an Application Administrator role.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "scoped_admin")},
			{Name: "suffix", Description: "Suffix.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "suffix")},
			{Name: "sys_class_name", Description: "Class.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_name")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Updates.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_name", Description: "Display name for this application file.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_name")},
			{Name: "sys_package", Description: "Package.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_package")},
			{Name: "sys_policy", Description: "Determines how application files are protected when downloaded or installed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_policy")},
			{Name: "sys_scope", Description: "Application containing this record.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_scope")},
			{Name: "sys_update_name", Description: "Update name (changed). ", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_update_name")},
			{Name: "sys_updated_by", Description: "Updated by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
		},
	}
}
