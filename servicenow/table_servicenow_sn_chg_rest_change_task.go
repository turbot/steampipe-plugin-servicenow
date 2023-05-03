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
			{Name: "active", Description: "Active.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "active").NullIfEqual("")},
			{Name: "activity_due", Description: "Activity Due.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "activity_due").NullIfEqual("")},
			{Name: "additional_assignee_list", Description: "Additional Assignee List.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "additional_assignee_list").NullIfEqual("")},
			{Name: "approval_history", Description: "Approval History.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval_history").NullIfEqual("")},
			{Name: "approval_set", Description: "Approval Set.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval_set").NullIfEqual("")},
			{Name: "approval", Description: "Approval.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval").NullIfEqual("")},
			{Name: "assigned_to_name", Description: "Assigned To (Name).", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "assigned_to").NullIfEqual("")},
			{Name: "assigned_to", Description: "Assigned To.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "assigned_to").NullIfEqual("")},
			{Name: "assignment_group_name", Description: "Assignment Group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "assignment_group").NullIfEqual("")},
			{Name: "assignment_group", Description: "Assignment Group.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "assignment_group").NullIfEqual("")},
			{Name: "business_duration", Description: "Business Duration.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "business_duration").NullIfEqual("")},
			{Name: "business_service", Description: "Business Service.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "business_service").NullIfEqual("")},
			{Name: "calendar_duration", Description: "Calendar Duration.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "calendar_duration").NullIfEqual("")},
			{Name: "change_request_number", Description: "Change Request Number.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "change_request").NullIfEqual("")},
			{Name: "change_request", Description: "Change Request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_request").NullIfEqual("")},
			{Name: "change_task_type", Description: "Change Task Type.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_task_type").NullIfEqual("")},
			{Name: "close_code", Description: "Close Code.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "close_code").NullIfEqual("")},
			{Name: "close_notes", Description: "Close Notes.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "close_notes").NullIfEqual("")},
			{Name: "closed_at", Description: "Closed At.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "closed_at").NullIfEqual("").Transform(parseDateTime)},
			{Name: "closed_by", Description: "Closed By.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "closed_by").NullIfEqual("")},
			{Name: "cmdb_ci_name", Description: "Cmdb Ci Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "cmdb_ci").NullIfEqual("")},
			{Name: "cmdb_ci", Description: "Cmdb Ci.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cmdb_ci").NullIfEqual("")},
			{Name: "comments_and_work_notes", Description: "Comments And Work Notes.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "comments_and_work_notes").NullIfEqual("")},
			{Name: "comments", Description: "Comments.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "comments").NullIfEqual("")},
			{Name: "company", Description: "Company.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "company").NullIfEqual("")},
			{Name: "contact_type", Description: "Contact Type.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "contact_type").NullIfEqual("")},
			{Name: "contract", Description: "Contract.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "contract").NullIfEqual("")},
			{Name: "correlation_display", Description: "Correlation Display.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "correlation_display").NullIfEqual("")},
			{Name: "correlation_id", Description: "Correlation Id.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "correlation_id").NullIfEqual("")},
			{Name: "created_from", Description: "Created From.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "created_from").NullIfEqual("")},
			{Name: "delivery_plan", Description: "Delivery Plan.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "delivery_plan").NullIfEqual("")},
			{Name: "delivery_task", Description: "Delivery Task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "delivery_task").NullIfEqual("")},
			{Name: "description", Description: "Description.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "description").NullIfEqual("")},
			{Name: "due_date", Description: "Due Date.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "due_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "escalation_name", Description: "Escalation Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "escalation").NullIfEqual("")},
			{Name: "escalation", Description: "Escalation.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "escalation").NullIfEqual("")},
			{Name: "expected_start", Description: "Expected Start.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "expected_start").NullIfEqual("")},
			{Name: "follow_up", Description: "Follow Up.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "follow_up").NullIfEqual("")},
			{Name: "group_list", Description: "Group List.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "group_list").NullIfEqual("")},
			{Name: "impact", Description: "Impact.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "impact").NullIfEqual("")},
			{Name: "knowledge", Description: "Knowledge.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "knowledge").NullIfEqual("")},
			{Name: "location", Description: "Location.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "location").NullIfEqual("")},
			{Name: "made_sla", Description: "Made Sla.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "made_sla").NullIfEqual("")},
			{Name: "number", Description: "Number.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "number").NullIfEqual("")},
			{Name: "on_hold_reason", Description: "On Hold Reason.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "on_hold_reason").NullIfEqual("")},
			{Name: "on_hold", Description: "On Hold.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "on_hold").NullIfEqual("")},
			{Name: "opened_at", Description: "Opened At.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "opened_at").NullIfEqual("").Transform(parseDateTime)},
			{Name: "opened_by", Description: "Opened By.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "opened_by").NullIfEqual("")},
			{Name: "order", Description: "Order.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "order").NullIfEqual("")},
			{Name: "parent", Description: "Parent Change Task or Change Request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "parent").NullIfEqual("")},
			{Name: "planned_end_date", Description: "Planned End Date.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "planned_end_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "planned_start_date", Description: "Planned Start Date.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "planned_start_date").NullIfEqual("").Transform(parseDateTime)},
			{Name: "priority", Description: "Priority.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "priority").NullIfEqual("")},
			{Name: "reassignment_count", Description: "Reassignment Count.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "reassignment_count").NullIfEqual("")},
			{Name: "rejection_goto", Description: "Rejection Goto.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "rejection_goto").NullIfEqual("")},
			{Name: "route_reason", Description: "Route Reason.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "route_reason").NullIfEqual("")},
			{Name: "service_offering", Description: "Service Offering.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "service_offering").NullIfEqual("")},
			{Name: "short_description", Description: "Short Description.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "short_description").NullIfEqual("")},
			{Name: "skills", Description: "Skills.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "skills").NullIfEqual("")},
			{Name: "sla_due", Description: "Sla Due.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sla_due").NullIfEqual("")},
			{Name: "sn_esign_document", Description: "Sn Esign Document.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sn_esign_document").NullIfEqual("")},
			{Name: "sn_esign_esignature_configuration", Description: "Sn Esign Esignature Configuration.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sn_esign_esignature_configuration").NullIfEqual("")},
			{Name: "state_name", Description: "State Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "state").NullIfEqual("")},
			{Name: "state", Description: "State.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "state").NullIfEqual("")},
			{Name: "sys_class_name", Description: "Sys Class Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_class_name").NullIfEqual("")},
			{Name: "sys_created_by", Description: "Sys Created By.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Sys Created On.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "sys_domain_path", Description: "Sys Domain Path.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain_path").NullIfEqual("")},
			{Name: "sys_domain", Description: "Sys Domain.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain").NullIfEqual("")},
			{Name: "sys_id", Description: "Unique identifier of the Change Task.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Sys Mod Count.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_tags", Description: "Sys Tags.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_tags").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "Sys Updated By.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Sys Updated On.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("").Transform(parseDateTime)},
			{Name: "task_effective_number", Description: "Task Effective Number.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task_effective_number").NullIfEqual("")},
			{Name: "time_worked", Description: "Time Worked.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "time_worked").NullIfEqual("")},
			{Name: "universal_request", Description: "Universal Request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "universal_request").NullIfEqual("")},
			{Name: "upon_approval", Description: "Upon Approval.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "upon_approval").NullIfEqual("")},
			{Name: "upon_reject", Description: "Upon Reject.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "upon_reject").NullIfEqual("")},
			{Name: "urgency", Description: "Urgency.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "urgency").NullIfEqual("")},
			{Name: "user_input", Description: "User Input.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "user_input").NullIfEqual("")},
			{Name: "variables", Description: "Variables.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "variables").NullIfEqual("")},
			{Name: "watch_list", Description: "Watch List.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "watch_list").NullIfEqual("")},
			{Name: "wf_activity", Description: "Wf Activity.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "wf_activity").NullIfEqual("")},
			{Name: "work_end", Description: "Work End.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldValue, "work_end").NullIfEqual("").Transform(parseDateTime)},
			{Name: "work_notes_list", Description: "Work Notes List.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_notes_list").NullIfEqual("")},
			{Name: "work_notes", Description: "Work Notes.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldDisplayValue, "work_notes").NullIfEqual("")},
			{Name: "work_start", Description: "Work Start.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_start").NullIfEqual("")},
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
