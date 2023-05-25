package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysAuditTableName = "sys_audit"

//// TABLE DEFINITION

func tableServicenowSysAudit() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_audit",
		Description: "Tracks changes made to ServiceNow tables.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"documentkey", "fieldname", "internal_checkpoint", "newvalue", "oldvalue", "reason", "sys_created_by", "sys_id", "tablename", "user"}),
			Hydrate:    listServicenowObjectsByTable(SysAuditTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysAuditTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "documentkey", Description: "Key of the audited document or record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "documentkey")},
			{Name: "fieldname", Description: "Name of the field that was audited.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "fieldname")},
			{Name: "internal_checkpoint", Description: "Checkpoint value used for internal tracking.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "internal_checkpoint")},
			{Name: "newvalue", Description: "New value of the audited field.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "newvalue")},
			{Name: "oldvalue", Description: "Old value of the audited field.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "oldvalue")},
			{Name: "reason", Description: "Reason or description for the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "reason")},
			{Name: "record_checkpoint", Description: "Checkpoint value for the audited record.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "record_checkpoint")},
			{Name: "sys_created_by", Description: "User who created the audit record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the audit record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the audit record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "tablename", Description: "Name of the table that was audited.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "tablename")},
			{Name: "user", Description: "User associated with the audit record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user")},
		},
	}
}
