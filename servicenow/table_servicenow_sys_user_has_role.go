package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysUserHasRoleTableName = "sys_user_has_role"

//// TABLE DEFINITION

func tableServicenowSysUserHasRole() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_user_has_role",
		Description: "User Role.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"inh_map", "state", "sys_created_by", "sys_id", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysUserHasRoleTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserHasRoleTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "granted_by", Description: "Granted by.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "granted_by")},
			{Name: "included_in_role_instance", Description: "Included in role instance.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "included_in_role_instance")},
			{Name: "included_in_role", Description: "Included in role.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "included_in_role")},
			{Name: "inh_count", Description: "Inheritance Count.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "inh_count")},
			{Name: "inh_map", Description: "Inheritance Map.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "inh_map")},
			{Name: "inherited", Description: "Inherited.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "inherited")},
			{Name: "role", Description: "Role.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "role")},
			{Name: "state", Description: "State.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "state")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Updates.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "Updated by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "user", Description: "User.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "user")},
		},
	}
}
