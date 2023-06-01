package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnChgRestChange() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_sn_chg_rest_change",
		Description:      "Change Management - Change request.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate: listServicenowSnChgRestChanges,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowSnChgRestChanges,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "Unique identifier of the associated change request record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_id").NullIfEqual("")},
			{Name: "action_status", Description: "Current action status of the associated change request.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "action_status").NullIfEqual("")},
			{Name: "active", Description: "Flag that indicates whether the change request is active.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "active").NullIfEqual("")},
			{Name: "activity_due", Description: "Date and time for which the associated case is expected to be completed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "activity_due").NullIfEqual("")},
			{Name: "additional_assignee_list", Description: "List of sys_ids of additional persons assigned to work on the change request.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "additional_assignee_list").NullIfEqual("")},
			{Name: "approval", Description: "Type of approval process required.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval").NullIfEqual("")},
			{Name: "approval_history", Description: "Most recent approval history journal entry.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval_history").NullIfEqual("")},
			{Name: "approval_set", Description: "Date and time that the associated action was approved.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "approval_set").NullIfEqual("")},
			{Name: "assigned_to", Description: "Sys_id of the user assigned to the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "assigned_to").NullIfEqual("")},
			{Name: "assignment_group", Description: "Sys_id of the group assigned to the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "assignment_group").NullIfEqual("")},
			{Name: "backout_plan", Description: "Description of the plan to execute if the change must be reversed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "backout_plan").NullIfEqual("")},
			{Name: "business_duration", Description: "Length in scheduled work hours, work days, and work weeks that it took to complete the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "business_duration").NullIfEqual("")},
			{Name: "business_service", Description: "Sys_id of the business service associated with the change request. Located in the Service [cmdb_ci_service] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "business_service").NullIfEqual("")},
			{Name: "cab_date", Description: "Date on which the Change Advisory Board (CAB) meets.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cab_date").NullIfEqual("")},
			{Name: "cab_delegate", Description: "Sys_id of the user that can substitute for the CAB manager during a CAB meeting. Located in the User [sys_user] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cab_delegate").NullIfEqual("")},
			{Name: "cab_recommendation", Description: "Description of the CAB recommendations for the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cab_recommendation").NullIfEqual("")},
			{Name: "cab_required", Description: "Flag that indicates whether the CAB is required.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "cab_required").NullIfEqual("")},
			{Name: "calendar_duration", Description: "Not currently used by Change Management.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "calendar_duration").NullIfEqual("")},
			{Name: "category", Description: "Category of the change, for example hardware, network, or software.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "category").NullIfEqual("")},
			{Name: "change_plan", Description: "Activities and roles for managing and controlling the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "change_plan").NullIfEqual("")},
			{Name: "chg_model", Description: "Sys_id of the change model that the associated change request was based on. Located in the Change Model [chg_model] table. The Change Model defines the state flow, transitions, and process activities that must be completed for the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "chg_model").NullIfEqual("")},
			{Name: "closed_at", Description: "Date and time that the associated change request was closed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "closed_at").NullIfEqual("")},
			{Name: "closed_by", Description: "Sys_id of the person that closed the change request. Located in the User [sys_user] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "closed_by").NullIfEqual("")},
			{Name: "close_code", Description: "Code assigned to the change request when it was closed. For example, Successful, Successful with issues, and Unsuccessful.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "close_code").NullIfEqual("")},
			{Name: "close_notes", Description: "Notes that the person entered when closing the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "close_notes").NullIfEqual("")},
			{Name: "cmdb_ci", Description: "Sys_id of the configuration item associated with the change request. Located in the Configuration Item [cmdb_ci] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "cmdb_ci").NullIfEqual("")},
			{Name: "comments", Description: "List of customer-facing work notes entered in the associated change request.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "comments").NullIfEqual("")},
			{Name: "comments_and_work_notes", Description: "List of both internal and customer-facing work notes entered for the associated change request.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "comments_and_work_notes").NullIfEqual("")},
			{Name: "company", Description: "Sys_id of the company associated with the change request. Located in the Company [core_company] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "company").NullIfEqual("")},
			{Name: "conflict_last_run", Description: "Date and time that the conflict detection script was last run on the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "conflict_last_run").NullIfEqual("")},
			{Name: "conflict_status", Description: "Current conflict status as detected by the conflict detection script, such as Conflict and Not Run.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "conflict_status").NullIfEqual("")},
			{Name: "contact_type", Description: "Method in which the change request was initially requested.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "contact_type").NullIfEqual("")},
			{Name: "contract", Description: "Sys_id of the contract associated with the change request. Located in the Contract [ast_contract] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "contract").NullIfEqual("")},
			{Name: "correlation_display", Description: "User-friendly name for the correlation_id.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "correlation_display").NullIfEqual("")},
			{Name: "correlation_id", Description: "Globally unique ID (GUID) of a matching change request record in a third-party system.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "correlation_id").NullIfEqual("")},
			{Name: "delivery_plan", Description: "No longer in use. Sys_id of the delivery plan associated with the change request. Located in the Execution Plan [sc_cat_item_delivery_plan] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "delivery_plan").NullIfEqual("")},
			{Name: "delivery_task", Description: "No longer in use. Sys_id of the delivery task associated with the change request. Located in the Execution Plan Task [sc_cat_item_delivery_task] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "delivery_task").NullIfEqual("")},
			{Name: "description", Description: "Detailed description of the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "description").NullIfEqual("")},
			{Name: "due_date", Description: "Task due date. Not used by change request process.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "due_date").NullIfEqual("")},
			{Name: "end_date", Description: "Date and time when the change request is to be completed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "end_date").NullIfEqual("")},
			{Name: "escalation", Description: "Current escalation level.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "escalation").NullIfEqual("")},
			{Name: "expected_start", Description: "Date and time when the task is to start. Not used by the change request process.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "expected_start").NullIfEqual("")},
			{Name: "follow_up", Description: "Date and time when a user followed-up with the person requesting the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "follow_up").NullIfEqual("")},
			{Name: "group_list", Description: "List of sys_ids and names of the groups associated with the change request.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "group_list").NullIfEqual("")},
			{Name: "impact", Description: "Impact on the change request will have on the customer.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "impact").NullIfEqual("")},
			{Name: "implementation_plan", Description: "Sequential steps to execute to implement this change. It also contains any dependencies between steps and assignee details for each step.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "implementation_plan").NullIfEqual("")},
			{Name: "justification", Description: "Benefits of implementing this change and the impact if this change is not implemented.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "justification").NullIfEqual("")},
			{Name: "knowledge", Description: "Flag that indicates whether there are any knowledge base ()KB) articles associated with the change request.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "knowledge").NullIfEqual("")},
			{Name: "location", Description: "Sys_id and name of the location of the equipment referenced in the change request. Located in the Location Located in the Location [cmn_location] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "location").NullIfEqual("")},
			{Name: "made_sla", Description: "No longer used. Flag that indicates whether the change request was implemented in alignment with the associated service level agreement.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "made_sla").NullIfEqual("")},
			{Name: "needs_attention", Description: "Flag that indicates whether the change request needs attention.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "needs_attention").NullIfEqual("")},
			{Name: "number", Description: "Change number assigned to the change request by the system, such as CHG0040007.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "number").NullIfEqual("")},
			{Name: "on_hold", Description: "Flag that indicates whether the change request is currently on hold.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "on_hold").NullIfEqual("")},
			{Name: "on_hold_reason", Description: "If the on_hold parameter is 'true', description of the reason why the change request is being held up.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "on_hold_reason").NullIfEqual("")},
			{Name: "on_hold_task", Description: "If the on_hold parameter is 'true', list of the sys_ids of the tasks that must be completed before the hold is released.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "on_hold_task").NullIfEqual("")},
			{Name: "opened_at", Description: "Date and time that the change release was created.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "opened_at").NullIfEqual("")},
			{Name: "opened_by", Description: "Sys_id and name of the user that created the change release. Located in the User [sys_user] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "opened_by").NullIfEqual("")},
			{Name: "order", Description: "Not used by Change Management. Optional numeric field by which to order records, such as when retrieving them from a database.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "order").NullIfEqual("")},
			{Name: "outside_maintenance_schedule", Description: "Flag that indicates whether maintenance by an outside company has been schedule for the change request.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "outside_maintenance_schedule").NullIfEqual("")},
			{Name: "parent", Description: "Sys_id and name of the parent task to this change request, if any. Located in the Task [task] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "parent").NullIfEqual("")},
			{Name: "phase", Description: "Current phase of the change request. This defines what the change is doing in greater detail.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "phase").NullIfEqual("")},
			{Name: "phase_state", Description: "Change_phase records that should be created for a change. They are dependent on the category, such that each type of change may have different change_phase records. The change_phase records provide an opportunity to control the approval process as each change_phase can have a schedule and a set of approvers.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "phase_state").NullIfEqual("")},
			{Name: "priority", Description: "Priority of the change request.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "priority").NullIfEqual("")},
			{Name: "production_system", Description: "Flag that indicates whether the change request is for a ServiceNow instance that is in a production environment.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "production_system").NullIfEqual("")},
			{Name: "reason", Description: "Description of why the change request was initiated.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "reason").NullIfEqual("")},
			{Name: "reassignment_count", Description: "Number of times that the change request has been reassigned to a new owner.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "reassignment_count").NullIfEqual("")},
			{Name: "rejection_goto", Description: "Sys_id of the task to perform if the change request is rejected. Located in the Task [table].", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "rejection_goto").NullIfEqual("")},
			{Name: "requested_by", Description: "Sys_id of the user that requested the change. Located in the User [sys_user] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "requested_by").NullIfEqual("")},
			{Name: "requested_by_date", Description: "Date and time when the change is requested to be implemented by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "requested_by_date").NullIfEqual("")},
			{Name: "review_comments", Description: "Comments entered when the change request was reviewed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "review_comments").NullIfEqual("")},
			{Name: "review_date", Description: "Date that the change request was reviewed.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "review_date").NullIfEqual("")},
			{Name: "review_status", Description: "Current status of the requested change request review.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "review_status").NullIfEqual("")},
			{Name: "risk", Description: "Level of risk associated with the change request.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "risk").NullIfEqual("")},
			{Name: "risk_impact_analysis", Description: "Description of the risk and analysis of implementing the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "risk_impact_analysis").NullIfEqual("")},
			{Name: "route_reason", Description: "Not currently used by Change Management. Reason that the change request was transferred.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "route_reason").NullIfEqual("")},
			{Name: "scope", Description: "Size of the change request.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "scope").NullIfEqual("")},
			{Name: "service_offering", Description: "Sys_id of the service offering associated with the change request. Service offerings uniquely define the level of service in terms of availability, scope, pricing, and packaging options. Located in the Offering [service_offering] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "service_offering").NullIfEqual("")},
			{Name: "short_description", Description: "Description of the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "short_description").NullIfEqual("")},
			{Name: "skills", Description: "List of the sys_ids of all of the skills required to implement the change request. Located in the Skill [cmn_skill] table.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "skills").NullIfEqual("")},
			{Name: "sla_due", Description: "No longer in use. Date and time that the change request must be completed based on the associated service level agreement.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sla_due").NullIfEqual("")},
			{Name: "sn_esign_document", Description: "Sys_id of any E-signed document attached to the change request. Located in the Attachment [sys_attachment] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sn_esign_document").NullIfEqual("")},
			{Name: "sn_esign_esignature_configuration", Description: "Sys_id of the E-signed signature template used for the associated document. Located in the E-signature Template [sn_esign_configuration] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sn_esign_esignature_configuration").NullIfEqual("")},
			{Name: "start_date", Description: "Date and time that the change request is planned to start implementation.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "start_date").NullIfEqual("")},
			{Name: "state", Description: "Current state of the change request. Possible values are defined in the change model.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "state").NullIfEqual("")},
			{Name: "std_change_producer_version", Description: "Sys_id of the record producer and change proposal associated with the change request. It also includes the number and percentage of successful and unsuccessful change requests created from the proposal. Located in the Standard Change Template Version [std_change_producer_version] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "std_change_producer_version").NullIfEqual("")},
			{Name: "sys_class_name", Description: "Name of the table in which the change request is located.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_class_name").NullIfEqual("")},
			{Name: "sys_created_by", Description: "Name of the user that initially created the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_by").NullIfEqual("")},
			{Name: "sys_created_on", Description: "Date and time that the associated change request record was originally created.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_created_on").NullIfEqual("")},
			{Name: "sys_domain", Description: "If using domains in the instance, the name of the domain to which the change module record is associated.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain").NullIfEqual("")},
			{Name: "sys_domain_path", Description: "If using domains in the instance, the domain path in which the associated change module record resides.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_domain_path").NullIfEqual("")},
			{Name: "sys_mod_count", Description: "Number of updates to the case since it was initially created.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "sys_mod_count").NullIfEqual("")},
			{Name: "sys_updated_by", Description: "Person that last updated the case.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_by").NullIfEqual("")},
			{Name: "sys_updated_on", Description: "Date and time when the case was last updated.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "sys_updated_on").NullIfEqual("")},
			{Name: "task_effective_number", Description: "Universal request number.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task_effective_number").NullIfEqual("")},
			{Name: "task_for", Description: "Not used by Change Management. Sys_id of the user that the task was created for. Located in the User [sys_user] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "task_for").NullIfEqual("")},
			{Name: "test_plan", Description: "Description of the associated test plan for the change.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "test_plan").NullIfEqual("")},
			{Name: "time_worked", Description: "Total amount of time worked on the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "time_worked").NullIfEqual("")},
			{Name: "type", Description: "Change request type.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "type").NullIfEqual("")},
			{Name: "unauthorized", Description: "Flag that indicates whether the change request is unauthorized", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldValue, "unauthorized").NullIfEqual("")},
			{Name: "universal_request", Description: "Sys_id of the Parent Universal request to which this change request is a part of. Located in the Task [task] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "universal_request").NullIfEqual("")},
			{Name: "upon_approval", Description: "Action to take if the change request is approved.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "upon_approval").NullIfEqual("")},
			{Name: "upon_reject", Description: "Action to take if the change request is rejected.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "upon_reject").NullIfEqual("")},
			{Name: "urgency", Description: "Urgency of the change request.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldValue, "urgency").NullIfEqual("")},
			{Name: "user_input", Description: "Additional user input.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "user_input").NullIfEqual("")},
			{Name: "variables", Description: "Name-value pairs of variables associated with the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "variables").NullIfEqual("")},
			{Name: "watch_list", Description: "List of sys_ids of the users who receive notifications about this change request when additional comments are added or if the state of a change request is changed to Resolved or Closed. Located in the User [sys_user] table.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "watch_list").NullIfEqual("")},
			{Name: "wf_activity", Description: "Sys_id of the workflow activity record associated with the change request. Located in the Workflow Activity [wf_activity] table.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "wf_activity").NullIfEqual("")},
			{Name: "work_end", Description: "Date and time work ended on the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_end").NullIfEqual("")},
			{Name: "work_notes", Description: "Information about how to resolve the change request, or steps taken to resolve it.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_notes").NullIfEqual("")},
			{Name: "work_notes_list", Description: "List of sys_ids of the internal users who receive notifications about this change request when work notes are added. Located in the User [sys_user] table.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldValue, "work_notes_list").NullIfEqual("")},
			{Name: "work_start", Description: "Date and time that work started on the change request.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldValue, "work_start").NullIfEqual("")},
		},
	}
}

//// LIST FUNCTION

func listServicenowSnChgRestChanges(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change.listServicenowSnChgRestChanges", "connect_error", err)
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
		err := client.SnChgRestChange.List(limit, offset, "", &response)
		totalReturned := len(response.Result)
		if err != nil {
			logger.Error("servicenow_sn_chg_rest_change.listServicenowSnChgRestChanges", "query_error", err)
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

//// GET FUNCTION

func getServicenowSnChgRestChanges(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change.getServicenowSnChgRestChanges", "connect_error", err)
		return nil, err
	}

	sysId := d.EqualsQualString("sys_id")

	var response snChgRestChangeGetResult
	err = client.SnChgRestChange.Get(sysId, &response)
	if err != nil {
		logger.Error("servicenow_sn_chg_rest_change.getServicenowSnChgRestChanges", "query_error", err)
		return nil, err
	}

	return response.Result, err
}

//// TRANSFORM FUNCTION

type snChgRestChangeListResult struct {
	Result []map[string]map[string]interface{} `json:"result"`
}

type snChgRestChangeGetResult struct {
	Result map[string]map[string]interface{} `json:"result"`
}

func getFieldValue(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	param := d.Param.(string)
	ls := d.HydrateItem.(map[string]map[string]interface{})
	return ls[param]["value"], nil
}

func getFieldDisplayValue(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	param := d.Param.(string)
	ls := d.HydrateItem.(map[string]map[string]interface{})
	return ls[param]["display_value"], nil
}
