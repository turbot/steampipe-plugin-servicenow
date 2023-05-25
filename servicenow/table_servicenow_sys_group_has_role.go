package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysGroupHasRoleTableName = "sys_group_has_role"

//// TABLE DEFINITION

func tableServicenowSysGroupHasRole() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_group_has_role",
		Description: "Represents relationships between user groups and roles.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"sys_created_by", "sys_id", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysGroupHasRoleTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysGroupHasRoleTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "granted_by", Description: "User or group who granted the role to the group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "granted_by")},
			{Name: "group", Description: "User group to which the role is assigned.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "group")},
			{Name: "inherits", Description: "Indicates whether the role is inherited.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "inherits")},
			{Name: "role", Description: "Role assigned to the group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "role")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
		},
	}
}
