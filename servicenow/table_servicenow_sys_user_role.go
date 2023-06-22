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
		Description: "Defines available roles in the ServiceNow.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"description", "includes_roles", "name", "requires_subscription", "suffix", "sys_class_name", "sys_created_by", "sys_id", "sys_name", "sys_policy", "sys_update_name", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysUserRoleTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserRoleTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "assignable_by", Description: "Roles that can assign this user role.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "assignable_by")},
			{Name: "can_delegate", Description: "Indicates if users with this role can delegate tasks.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "can_delegate")},
			{Name: "description", Description: "Description or details of the user role.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "description")},
			{Name: "elevated_privilege", Description: "Indicates if the user role has elevated privileges.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "elevated_privilege")},
			{Name: "grantable", Description: "Indicates if the user role can be granted to other users.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "grantable")},
			{Name: "includes_roles", Description: "Roles included within this user role.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "includes_roles")},
			{Name: "name", Description: "Name or label of the user role.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "requires_subscription", Description: "Indicates if the user role requires a subscription.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "requires_subscription")},
			{Name: "scoped_admin", Description: "Indicates if the user role has scoped administration privileges.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "scoped_admin")},
			{Name: "suffix", Description: "Suffix or postfix for the user role.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "suffix")},
			{Name: "sys_class_name", Description: "System class name of the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_name")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_name", Description: "Name of the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_name")},
			{Name: "sys_package", Description: "Package associated with the record.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_package")},
			{Name: "sys_policy", Description: "Policy associated with the user role.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_policy")},
			{Name: "sys_scope", Description: "Scope associated with the user role.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_scope")},
			{Name: "sys_update_name", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_update_name")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
		},
	}
}
