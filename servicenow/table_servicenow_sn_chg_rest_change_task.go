package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnChgRestChangeTask() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change_task",
		Description:      "Change Management - Change Task.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate:       listServicenowSnChgRestChangeTasks,
			ParentHydrate: listServicenowSnChgRestChanges,
		},
		Columns: []*plugin.Column{
			{Name: "active", Description: "Indicates whether the change task is active.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "active").NullIfEqual("")},
			{Name: "activity_due", Description: "Due date for the activity associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "activity_due").NullIfEqual("")},
			{Name: "additional_assignee_list", Description: "Additional users assigned to the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "additional_assignee_list").NullIfEqual("")},
			{Name: "approval_history", Description: "Approval history of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval_history").NullIfEqual("")},
			{Name: "approval_set", Description: "Approval set associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval_set").NullIfEqual("")},
			{Name: "approval", Description: "Approval associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval").NullIfEqual("")},
			{Name: "assigned_to_name", Description: "Name of the user assigned to the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "assigned_to").NullIfEqual("")},
			{Name: "assigned_to", Description: "User assigned to the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "assigned_to").NullIfEqual("")},
			{Name: "assignment_group_name", Description: "Name of the assignment group for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "assignment_group").NullIfEqual("")},
			{Name: "assignment_group", Description: "Assignment group for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "assignment_group").NullIfEqual("")},
			{Name: "business_duration", Description: "Business duration of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "business_duration").NullIfEqual("")},
			{Name: "business_service", Description: "Business service associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "business_service").NullIfEqual("")},
			{Name: "calendar_duration", Description: "Calendar duration of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "calendar_duration").NullIfEqual("")},
			{Name: "change_request_number", Description: "Number of the change request associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "change_request").NullIfEqual("")},
			{Name: "change_request", Description: "Change request associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_request").NullIfEqual("")},
			{Name: "change_task_type", Description: "Type of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_task_type").NullIfEqual("")},
			{Name: "close_code", Description: "Code indicating the reason for closing the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "close_code").NullIfEqual("")},
			{Name: "close_notes", Description: "Notes indicating the details of the closure of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "close_notes").NullIfEqual("")},
			{Name: "closed_at", Description: "Date and time when the change task was closed.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "closed_at").NullIfEqual("").Transform(parseDateTime)},
			{Name: "closed_by", Description: "User who closed the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "closed_by").NullIfEqual("")},
			{Name: "cmdb_ci_name", Description: "Name of the configuration item associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "cmdb_ci").NullIfEqual("")},
			{Name: "cmdb_ci", Description: "Configuration item associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cmdb_ci").NullIfEqual("")},
			{Name: "comments_and_work_notes", Description: "Comments and work notes added to the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "comments_and_work_notes").NullIfEqual("")},
			{Name: "comments", Description: "Comments added to the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "comments").NullIfEqual("")},
			{Name: "company", Description: "Company associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "company").NullIfEqual("")},
			{Name: "contact_type", Description: "Type of contact associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "contact_type").NullIfEqual("")},
			{Name: "contract", Description: "Contract associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "contract").NullIfEqual("")},
			{Name: "correlation_display", Description: "Display value of the correlation ID for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "correlation_display").NullIfEqual("")},
			{Name: "correlation_id", Description: "Correlation ID for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "correlation_id").NullIfEqual("")},
			{Name: "created_from", Description: "Source from which the change task was created.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "created_from").NullIfEqual("")},
			{Name: "delivery_plan", Description: "Delivery plan associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "delivery_plan").NullIfEqual("")},
			{Name: "delivery_task", Description: "Delivery task associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "delivery_task").NullIfEqual("")},
			{Name: "description", Description: "Description of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "description").NullIfEqual("")},
			{Name: "due_date", Description: "Due date for the change task.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "due_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "escalation_name", Description: "Name of the escalation associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "escalation").NullIfEqual("")},
			{Name: "escalation", Description: "Escalation associated with the change task.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "escalation").NullIfEqual("")},
			{Name: "expected_start", Description: "Expected start date for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "expected_start").NullIfEqual("")},
			{Name: "follow_up", Description: "Follow-up details for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "follow_up").NullIfEqual("")},
			{Name: "group_list", Description: "List of groups associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "group_list").NullIfEqual("")},
			{Name: "impact", Description: "Impact of the change task.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "impact").NullIfEqual("")},
			{Name: "knowledge", Description: "Indicates whether the change task is based on knowledge.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "knowledge").NullIfEqual("")},
			{Name: "location", Description: "Location associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "location").NullIfEqual("")},
			{Name: "made_sla", Description: "Indicates whether the change task met the service level agreement (SLA).", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "made_sla").NullIfEqual("")},
			{Name: "number", Description: "Number of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "number").NullIfEqual("")},
			{Name: "on_hold_reason", Description: "Reason for placing the change task on hold.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "on_hold_reason").NullIfEqual("")},
			{Name: "on_hold", Description: "Indicates whether the change task is on hold.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "on_hold").NullIfEqual("")},
			{Name: "opened_at", Description: "Date and time when the change task was opened.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "opened_at").NullIfEqual("").Transform(parseDateTime)},
			{Name: "opened_by", Description: "User who opened the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "opened_by").NullIfEqual("")},
			{Name: "order", Description: "Order value of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "order").NullIfEqual("")},
			{Name: "parent", Description: "Parent change task of the current change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "parent").NullIfEqual("")},
			{Name: "planned_end_date", Description: "Planned end date for the change task.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "planned_end_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "planned_start_date", Description: "Planned start date for the change task.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "planned_start_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "priority", Description: "Priority of the change task.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "priority").NullIfEqual("")},
			{Name: "reassignment_count", Description: "Number of times the change task was reassigned.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "reassignment_count").NullIfEqual("")},
			{Name: "rejection_goto", Description: "Destination for rejection in the change task workflow.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "rejection_goto").NullIfEqual("")},
			{Name: "route_reason", Description: "Reason for routing the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "route_reason").NullIfEqual("")},
			{Name: "service_offering", Description: "Service offering associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "service_offering").NullIfEqual("")},
			{Name: "short_description", Description: "Short description of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "short_description").NullIfEqual("")},
			{Name: "skills", Description: "Skills associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "skills").NullIfEqual("")},
			{Name: "sla_due", Description: "Due date for the service level agreement (SLA) associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sla_due").NullIfEqual("")},
			{Name: "sn_esign_document", Description: "Electronic signature document associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sn_esign_document").NullIfEqual("")},
			{Name: "sn_esign_esignature_configuration", Description: "Electronic signature configuration associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sn_esign_esignature_configuration").NullIfEqual("")},
			{Name: "state_name", Description: "Name of the state of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "state").NullIfEqual("")},
			{Name: "state", Description: "State of the change task.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "state").NullIfEqual("")},
			{Name: "sys_class_name", Description: "Class name of the change task record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_class_name").NullIfEqual("")},
			{Name: "sys_created_by", Description: "User who created the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Date and time when the change task was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "sys_domain_path", Description: "Path of the domain associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain_path").NullIfEqual("")},
			{Name: "sys_domain", Description: "Domain associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain").NullIfEqual("")},
			{Name: "sys_id", Description: "Unique system identifier for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Number of times the change task was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_tags", Description: "Tags associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_tags").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "User who last updated the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Date and time when the change task was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "task_effective_number", Description: "Effective number of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task_effective_number").NullIfEqual("")},
			{Name: "time_worked", Description: "Amount of time worked on the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "time_worked").NullIfEqual("")},
			{Name: "universal_request", Description: "Universal request associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "universal_request").NullIfEqual("")},
			{Name: "upon_approval", Description: "Actions to be performed upon approval of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "upon_approval").NullIfEqual("")},
			{Name: "upon_reject", Description: "Actions to be performed upon rejection of the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "upon_reject").NullIfEqual("")},
			{Name: "urgency", Description: "Urgency of the change task.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "urgency").NullIfEqual("")},
			{Name: "user_input", Description: "User input associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "user_input").NullIfEqual("")},
			{Name: "variables", Description: "Variables associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "variables").NullIfEqual("")},
			{Name: "watch_list", Description: "Users on the watch list for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "watch_list").NullIfEqual("")},
			{Name: "wf_activity", Description: "Workflow activity associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "wf_activity").NullIfEqual("")},
			{Name: "work_end", Description: "End time of the work performed for the change task.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "work_end").NullIfEqual("").Transform(parseDateTime)},
			{Name: "work_notes_list", Description: "Work notes associated with the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_notes_list").NullIfEqual("")},
			{Name: "work_notes", Description: "Work notes added to the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "work_notes").NullIfEqual("")},
			{Name: "work_start", Description: "Start time of the work performed for the change task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_start").NullIfEqual("")},
		},
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChangeTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	changes := h.Item.(map[string]map[string]interface{})
	if changes["sys_id"]["value"] == nil {
		return nil, nil
	}
	change_sys_id := changes["sys_id"]["value"]

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change_task.listServicenowSnChgRestChangeTasks", "connect_error", err)
		return nil, err
	}

	offset := 0
	limit := 30
	if d.QueryContext.Limit != nil {
		pgLimit := int(*d.QueryContext.Limit)
		if pgLimit < limit {
			limit = pgLimit
		}
	}

	for {
		var response snChgRestChangeListResult
		err := client.SnChgRestChangeTask.List(change_sys_id.(string), limit, offset, "", &response)
		totalReturned := len(response.Result)
		if err != nil {
			logger.Error("servicenow_sn_chg_rest_change_task.listServicenowSnChgRestChangeTasks", "query_error", err)
			return nil, err
		}
		for _, object := range response.Result {
			d.StreamListItem(ctx, object)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if totalReturned < limit {
			break
		}
		offset += limit
	}
	return nil, err
}
