package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysAuditRelationTableName = "sys_audit_relation"

//// TABLE DEFINITION

func tableServicenowSysAuditRelation() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_audit_relation",
		Description: "Table relationship audit record.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"documentkey", "sys_created_by", "sys_id", "tablename"}),
			Hydrate:    listServicenowObjectsByTable(SysAuditRelationTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysAuditRelationTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "audit", Description: "Reference to the audit record in the `sys_audit` table.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "audit")},
			{Name: "audit_delete", Description: "Indicates if the relation is deleted during an audit deletion.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "audit_delete")},
			{Name: "documentkey", Description: "Key of the related audited document or record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "documentkey")},
			{Name: "sys_created_by", Description: "User who created the audit relation record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the audit relation record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Unique system identifier for the audit relation record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "tablename", Description: "Name of the table that is related to the audit.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "tablename")},
		},
	}
}
