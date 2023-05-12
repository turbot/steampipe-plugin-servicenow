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
		Description: "Group Role.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"sys_created_by", "sys_id", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysGroupHasRoleTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysGroupHasRoleTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "granted_by", Description: "Granted by.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "granted_by")},
			{Name: "group", Description: "Group.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "group")},
			{Name: "inherits", Description: "Members of this group automatically are granted this roles.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "inherits")},
			{Name: "role", Description: "Role.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "role")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Updates.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "Updated by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
		},
	}
}
