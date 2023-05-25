package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysUserGroupTableName = "sys_user_group"

//// TABLE DEFINITION

func tableServicenowSysUserGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_user_group",
		Description: "User Group.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"description", "email", "name", "roles", "source", "sys_created_by", "sys_id", "sys_updated_by", "type"}),
			Hydrate:    listServicenowObjectsByTable(SysUserGroupTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserGroupTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "active", Description: "Indicates whether the user group is active or inactive.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "active")},
			{Name: "cost_center", Description: "Cost center associated with the user group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cost_center")},
			{Name: "default_assignee", Description: "Default user assigned to tasks assigned to the user group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "default_assignee")},
			{Name: "description", Description: "Description or additional information about the user group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "description")},
			{Name: "email", Description: "Email address associated with the user group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "email")},
			{Name: "exclude_manager", Description: "Determines whether the manager is excluded from the user group.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "exclude_manager")},
			{Name: "include_members", Description: "Determines whether members are included in the user group.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "include_members")},
			{Name: "manager", Description: "Manager assigned to the user group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "manager")},
			{Name: "name", Description: "Name of the user group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "parent", Description: "Parent user group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "parent")},
			{Name: "roles", Description: "Roles assigned to the user group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "roles")},
			{Name: "source", Description: "Source of the user group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "source")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "type", Description: "Type of the user group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "type")},
		},
	}
}
