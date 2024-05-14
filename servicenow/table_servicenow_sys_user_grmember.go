package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysUserGroupMemberTableName = "sys_user_grmember"

//// TABLE DEFINITION

func tableServicenowSysUserGroupMember() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_user_grmember",
		Description: "User Group Member.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"sys_created_by", "sys_id", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysUserGroupMemberTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserGroupMemberTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "group", Description: "User group to which the user is a member.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "group")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "user", Description: "User who is a member of the user group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "user")},
		}),
	}
}
