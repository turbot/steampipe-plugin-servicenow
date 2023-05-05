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
		Description: "Table change record.",
		List: &plugin.ListConfig{
			Hydrate: listServicenowObjectsByTable(SysAuditTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysAuditTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "documentkey", Description: "Document Key.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "documentkey")},
			{Name: "fieldname", Description: "Field Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "fieldname")},
			{Name: "internal_checkpoint", Description: "Record internal checkpoint.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "internal_checkpoint")},
			{Name: "newvalue", Description: "New value.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "newvalue")},
			{Name: "oldvalue", Description: "Old value.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "oldvalue")},
			{Name: "reason", Description: "Reason.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "reason")},
			{Name: "record_checkpoint", Description: "Update count.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "record_checkpoint")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "tablename", Description: "Table Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "tablename")},
			{Name: "user", Description: "User.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user")},
		},
	}
}
