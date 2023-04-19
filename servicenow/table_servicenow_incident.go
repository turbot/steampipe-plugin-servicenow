package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const IncidentTableName = "incident"

//// TABLE DEFINITION

func tableServicenowIncident() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_incident",
		Description:      "Incident.",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: listServicenowObjectsByTable(IncidentTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(IncidentTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "business_stc", Description: "Business time elapsed (stored in seconds) before Incident was resolved.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "business_stc")},
			{Name: "calendar_stc", Description: "Time elapsed (stored in seconds) before Incident was Resolved.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "calendar_stc")},
			{Name: "child_incidents", Description: "Number of child Incidents related to this Problem.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "child_incidents")},
			{Name: "hold_reason", Description: "On hold reason.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "hold_reason")},
			{Name: "incident_state", Description: "Workflow state of the incident.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "incident_state")},
			{Name: "notify", Description: "Notify.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "notify")},
			{Name: "reopen_count", Description: "Number of times Incident state has changed from Resolved or Closed to another state.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "reopen_count")},
			{Name: "severity", Description: "Severity.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "severity")},
			{Name: "caller_id", Description: "Person who reported or is affected by this incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "caller_id")},
			{Name: "caused_by", Description: "Change request that caused the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "caused_by")},
			{Name: "origin_id", Description: "Origin.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "origin_id")},
			{Name: "parent_incident", Description: "Can be used to collect Incidents for the same root issue.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "parent_incident")},
			{Name: "problem_id", Description: "Related problem, if one exists.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "problem_id")},
			{Name: "reopened_by", Description: "Last reopened by.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "reopened_by")},
			{Name: "resolved_by", Description: "Resolved by.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "resolved_by")},
			{Name: "rfc", Description: "Related change, if one exists.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "rfc")},
			{Name: "business_impact", Description: "Business impact.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "business_impact")},
			{Name: "category", Description: "Category.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "category")},
			{Name: "cause", Description: "Probable cause.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "cause")},
			{Name: "close_code", Description: "For use in reporting.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "close_code")},
			{Name: "origin_table", Description: "Table of the Origin record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "origin_table")},
			{Name: "subcategory", Description: "Subcategory.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "subcategory")},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "reopened_time", Description: "Last reopened at.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "reopened_timee").Transform(parseDateTime)},
			{Name: "resolved_at", Description: "Resolved.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "resolved_ate").Transform(parseDateTime)},
		},
	}
}
