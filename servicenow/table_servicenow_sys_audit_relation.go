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
			{Name: "audit", Description: "Audit.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "audit")},
			{Name: "audit_delete", Description: "Audit delete.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "audit_delete")},
			{Name: "documentkey", Description: "Document Key.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "documentkey")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "tablename", Description: "Table Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "tablename")},
		},
	}
}
