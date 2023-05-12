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
			{Name: "active", Description: "Active.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "active")},
			{Name: "cost_center", Description: "Cost center.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cost_center")},
			{Name: "default_assignee", Description: "Default assignee for this assignment group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "default_assignee")},
			{Name: "description", Description: "Description.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "description")},
			{Name: "email", Description: "By default the group email address overrides individual email addresses of members.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "email")},
			{Name: "exclude_manager", Description: "Manager will not get email notifications sent to this group.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "exclude_manager")},
			{Name: "include_members", Description: "Should members also get email when a group email is supplied?. ", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "include_members")},
			{Name: "manager", Description: "Manager.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "manager")},
			{Name: "name", Description: "Descriptive name, e.g. DBAs, Network, etc.. ", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "parent", Description: "Parent.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "parent")},
			{Name: "roles", Description: "Roles.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "roles")},
			{Name: "source", Description: "Source.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "source")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Updates.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "Updated by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "type", Description: "Types of this group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "type")},
		},
	}
}
