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
		Description:      "Tracks reported incidents.",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"additional_assignee_list", "approval_history", "approval", "business_impact", "category", "cause", "close_code", "close_notes", "comments_and_work_notes", "comments", "contact_type", "correlation_display", "correlation_id", "description", "group_list", "number", "origin_table", "short_description", "skills", "subcategory", "sys_class_name", "sys_created_by", "sys_domain_path", "sys_id", "sys_updated_by", "task_effective_number", "upon_approval", "upon_reject", "user_input", "variables", "watch_list", "work_notes_list", "work_notes"}),
			Hydrate:    listServicenowObjectsByTable(IncidentTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(IncidentTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "active", Description: "Indicates if the incident is currently active.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "active")},
			{Name: "activity_due", Description: "Due date and time for the current activity associated with the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "activity_due").Transform(parseDateTime)},
			{Name: "additional_assignee_list", Description: "List of additional users assigned to the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "additional_assignee_list")},
			{Name: "approval_history", Description: "History of approvals associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "approval_history")},
			{Name: "approval_set", Description: "Approval set associated with the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "approval_set").Transform(parseDateTime)},
			{Name: "approval", Description: "Current approval state of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "approval")},
			{Name: "assigned_to", Description: "User or group assigned to handle the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "assigned_to")},
			{Name: "assignment_group", Description: "Group assigned to handle the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "assignment_group")},
			{Name: "business_duration", Description: "Duration of the incident in business hours.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "business_duration").Transform(parseDateTime)},
			{Name: "business_impact", Description: "Impact of the incident on the business.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "business_impact")},
			{Name: "business_service", Description: "Business service associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "business_service")},
			{Name: "business_stc", Description: "Service target compliance for the incident in business hours.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "business_stc")},
			{Name: "calendar_duration", Description: "Duration of the incident in calendar time.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "calendar_duration").Transform(parseDateTime)},
			{Name: "calendar_stc", Description: "Service target compliance for the incident in calendar time.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "calendar_stc")},
			{Name: "caller_id", Description: "User who reported the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "caller_id")},
			{Name: "category", Description: "Category or classification of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "category")},
			{Name: "cause", Description: "Root cause or reason for the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "cause")},
			{Name: "caused_by", Description: "Configuration item or event that caused the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "caused_by")},
			{Name: "child_incidents", Description: "Child incidents associated with the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "child_incidents")},
			{Name: "close_code", Description: "Code indicating the reason for closing the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "close_code")},
			{Name: "close_notes", Description: "Notes or comments added when closing the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "close_notes")},
			{Name: "closed_at", Description: "Date and time when the incident was closed.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "closed_at").Transform(parseDateTime)},
			{Name: "closed_by", Description: "User who closed the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "closed_by")},
			{Name: "cmdb_ci", Description: "Configuration item associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cmdb_ci")},
			{Name: "comments_and_work_notes", Description: "Combined list of comments and work notes for the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "comments_and_work_notes")},
			{Name: "comments", Description: "Additional comments or notes on the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "comments")},
			{Name: "company", Description: "Company or organization associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "company")},
			{Name: "contact_type", Description: "Type of contact associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "contact_type")},
			{Name: "contract", Description: "Contract associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "contract")},
			{Name: "correlation_display", Description: "Display value for the correlation ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "correlation_display")},
			{Name: "correlation_id", Description: "Identifier for correlating related records.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "correlation_id")},
			{Name: "delivery_plan", Description: "Plan for delivering a resolution to the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "delivery_plan")},
			{Name: "delivery_task", Description: "Task associated with delivering a resolution to the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "delivery_task")},
			{Name: "description", Description: "Description or details of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "description")},
			{Name: "due_date", Description: "Due date and time for resolving the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "due_date").Transform(parseDateTime)},
			{Name: "escalation", Description: "Escalation details for the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "escalation")},
			{Name: "expected_start", Description: "Expected start date and time for the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "expected_start").Transform(parseDateTime)},
			{Name: "follow_up", Description: "Date and time for a follow-up action on the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "follow_up").Transform(parseDateTime)},
			{Name: "group_list", Description: "List of groups associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "group_list")},
			{Name: "hold_reason", Description: "Reason for placing the incident on hold.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "hold_reason")},
			{Name: "impact", Description: "Impact level of the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "impact")},
			{Name: "incident_state", Description: "Current state of the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "incident_state")},
			{Name: "knowledge", Description: "Knowledge article associated with the incident.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "knowledge")},
			{Name: "location", Description: "Location or place associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "location")},
			{Name: "made_sla", Description: "Indicates if the incident met the service level agreement.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "made_sla")},
			{Name: "notify", Description: "Flag indicating if notifications are sent for the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "notify")},
			{Name: "number", Description: "Unique number or identifier of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "number")},
			{Name: "opened_at", Description: "Date and time when the incident was opened.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "opened_at").Transform(parseDateTime)},
			{Name: "opened_by", Description: "User who opened the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "opened_by")},
			{Name: "order", Description: "Order associated with the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "order")},
			{Name: "origin_id", Description: "Identifier for the source of the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "origin_id")},
			{Name: "origin_table", Description: "Table or source from which the incident originated.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "origin_table")},
			{Name: "parent_incident", Description: "Parent incident associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "parent_incident")},
			{Name: "parent", Description: "Parent record associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "parent")},
			{Name: "priority", Description: "Priority level assigned to the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "priority")},
			{Name: "problem_id", Description: "Problem record associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "problem_id")},
			{Name: "reassignment_count", Description: "Number of times the incident has been reassigned.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "reassignment_count")},
			{Name: "rejection_goto", Description: "Next state in the workflow when the incident is rejected.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "rejection_goto")},
			{Name: "reopen_count", Description: "Number of times the incident has been reopened.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "reopen_count")},
			{Name: "reopened_by", Description: "User who reopened the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "reopened_by")},
			{Name: "reopened_time", Description: "Date and time when the incident was reopened.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "reopened_time").Transform(parseDateTime)},
			{Name: "resolved_at", Description: "Date and time when the incident was resolved.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "resolved_at").Transform(parseDateTime)},
			{Name: "resolved_by", Description: "User who resolved the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "resolved_by")},
			{Name: "rfc", Description: "Related Request for Change (RFC) associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "rfc")},
			{Name: "route_reason", Description: "Reason for routing or assigning the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "route_reason")},
			{Name: "service_offering", Description: "Service offering associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "service_offering")},
			{Name: "severity", Description: "Severity level assigned to the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "severity")},
			{Name: "short_description", Description: "Brief description or summary of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "short_description")},
			{Name: "skills", Description: "Skills required to handle the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "skills")},
			{Name: "sla_due", Description: "Due date and time for meeting the service level agreement (SLA).", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sla_due").Transform(parseDateTime)},
			{Name: "sn_esign_document", Description: "Document associated with electronic signature.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sn_esign_document")},
			{Name: "sn_esign_esignature_configuration", Description: "Electronic signature configuration associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sn_esign_esignature_configuration")},
			{Name: "state", Description: "Current state or status of the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "state")},
			{Name: "subcategory", Description: "Subcategory or more specific classification of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "subcategory")},
			{Name: "sys_class_name", Description: "Name of the system class or table for the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_name")},
			{Name: "sys_created_by", Description: "User who created the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the incident was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_domain_path", Description: "Path of the domain associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain_path")},
			{Name: "sys_domain", Description: "Domain associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain")},
			{Name: "sys_id", Description: "Unique system identifier for the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of times the incident has been modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "User who last updated the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time of the last update to the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "task_effective_number", Description: "Effective number of tasks associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "task_effective_number")},
			{Name: "time_worked", Description: "Total time worked on the incident.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "time_worked").Transform(parseDateTime)},
			{Name: "universal_request", Description: "Universal Request associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "universal_request")},
			{Name: "upon_approval", Description: "Actions to be performed upon approval of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "upon_approval")},
			{Name: "upon_reject", Description: "Actions to be performed upon rejection of the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "upon_reject")},
			{Name: "urgency", Description: "Urgency level assigned to the incident.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "urgency")},
			{Name: "user_input", Description: "User input or response associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user_input")},
			{Name: "variables", Description: "Variables or additional data associated with the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "variables")},
			{Name: "watch_list", Description: "List of users watching or monitoring the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "watch_list")},
			{Name: "wf_activity", Description: "Workflow activity associated with the incident.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "wf_activity")},
			{Name: "work_end", Description: "Date and time when work on the incident ends.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "work_end").Transform(parseDateTime)},
			{Name: "work_notes_list", Description: "List of work notes added to the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "work_notes_list")},
			{Name: "work_notes", Description: "Additional work notes or comments on the incident.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "work_notes")},
			{Name: "work_start", Description: "Date and time when work on the incident starts.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "work_start").Transform(parseDateTime)},
		}),
	}
}
