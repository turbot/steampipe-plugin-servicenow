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
		Description: "Tracks assigned roles for users.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"inh_map", "state", "sys_created_by", "sys_id", "sys_updated_by"}),
			Hydrate:    listServicenowObjectsByTable(SysUserHasRoleTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserHasRoleTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "granted_by", Description: "User or role that granted this role to the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "granted_by")},
			{Name: "included_in_role_instance", Description: "Role instance in which this role is included.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "included_in_role_instance")},
			{Name: "included_in_role", Description: "Role in which this role is included.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "included_in_role")},
			{Name: "inh_count", Description: "Count of inherited roles.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "inh_count")},
			{Name: "inh_map", Description: "Mapping of inherited roles.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "inh_map")},
			{Name: "inherited", Description: "Indicates if the role is inherited.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "inherited")},
			{Name: "role", Description: "Role assigned to the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "role")},
			{Name: "state", Description: "State of the role assignment.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "state")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "user", Description: "User to whom the role is assigned.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "user")},
		}),
	}
}
