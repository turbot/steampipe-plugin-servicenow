package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowIncident() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_incident",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowIncidents,
		},
		Columns: []*plugin.Column{
			// {Name: "raw", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromValue()},
			{Name: "parent", Description: "", Type: proto.ColumnType_STRING},
			{Name: "caused_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "watch_list", Description: "", Type: proto.ColumnType_STRING},
			{Name: "upon_reject", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "child_incidents", Description: "", Type: proto.ColumnType_STRING},
			{Name: "hold_reason", Description: "", Type: proto.ColumnType_STRING},
			{Name: "origin_table", Description: "", Type: proto.ColumnType_STRING},
			{Name: "task_effective_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "approval_history", Description: "", Type: proto.ColumnType_STRING},
			{Name: "number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "user_input", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "state", Description: "", Type: proto.ColumnType_STRING},
			{Name: "route_reason", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "knowledge", Description: "", Type: proto.ColumnType_STRING},
			{Name: "order", Description: "", Type: proto.ColumnType_STRING},
			{Name: "calendar_stc", Description: "", Type: proto.ColumnType_STRING},
			{Name: "closed_at", Description: "", Type: proto.ColumnType_STRING},
			{Name: "delivery_plan", Description: "", Type: proto.ColumnType_STRING},
			{Name: "contract", Description: "", Type: proto.ColumnType_STRING},
			{Name: "impact", Description: "", Type: proto.ColumnType_STRING},
			{Name: "active", Description: "", Type: proto.ColumnType_STRING},
			{Name: "work_notes_list", Description: "", Type: proto.ColumnType_STRING},
			{Name: "business_impact", Description: "", Type: proto.ColumnType_STRING},
			{Name: "priority", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain_path", Description: "", Type: proto.ColumnType_STRING},
			{Name: "rfc", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_worked", Description: "", Type: proto.ColumnType_STRING},
			{Name: "expected_start", Description: "", Type: proto.ColumnType_STRING},
			{Name: "opened_at", Description: "", Type: proto.ColumnType_STRING},
			{Name: "business_duration", Description: "", Type: proto.ColumnType_STRING},
			{Name: "group_list", Description: "", Type: proto.ColumnType_STRING},
			{Name: "work_end", Description: "", Type: proto.ColumnType_STRING},
			{Name: "reopened_time", Description: "", Type: proto.ColumnType_STRING},
			{Name: "resolved_at", Description: "", Type: proto.ColumnType_STRING},
			{Name: "approval_set", Description: "", Type: proto.ColumnType_STRING},
			{Name: "subcategory", Description: "", Type: proto.ColumnType_STRING},
			{Name: "work_notes", Description: "", Type: proto.ColumnType_STRING},
			{Name: "universal_request", Description: "", Type: proto.ColumnType_STRING},
			{Name: "short_description", Description: "", Type: proto.ColumnType_STRING},
			{Name: "close_code", Description: "", Type: proto.ColumnType_STRING},
			{Name: "correlation_display", Description: "", Type: proto.ColumnType_STRING},
			{Name: "delivery_task", Description: "", Type: proto.ColumnType_STRING},
			{Name: "work_start", Description: "", Type: proto.ColumnType_STRING},
			{Name: "additional_assignee_list", Description: "", Type: proto.ColumnType_STRING},
			{Name: "business_stc", Description: "", Type: proto.ColumnType_STRING},
			{Name: "cause", Description: "", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "", Type: proto.ColumnType_STRING},
			{Name: "calendar_duration", Description: "", Type: proto.ColumnType_STRING},
			{Name: "close_notes", Description: "", Type: proto.ColumnType_STRING},
			{Name: "notify", Description: "", Type: proto.ColumnType_STRING},
			{Name: "service_offering", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_class_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "follow_up", Description: "", Type: proto.ColumnType_STRING},
			{Name: "parent_incident", Description: "", Type: proto.ColumnType_STRING},
			{Name: "contact_type", Description: "", Type: proto.ColumnType_STRING},
			{Name: "reopened_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "incident_state", Description: "", Type: proto.ColumnType_STRING},
			{Name: "urgency", Description: "", Type: proto.ColumnType_STRING},
			{Name: "reassignment_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "activity_due", Description: "", Type: proto.ColumnType_STRING},
			{Name: "severity", Description: "", Type: proto.ColumnType_STRING},
			{Name: "comments", Description: "", Type: proto.ColumnType_STRING},
			{Name: "approval", Description: "", Type: proto.ColumnType_STRING},
			{Name: "comments_and_work_notes", Description: "", Type: proto.ColumnType_STRING},
			{Name: "due_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "reopen_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_tags", Description: "", Type: proto.ColumnType_STRING},
			{Name: "escalation", Description: "", Type: proto.ColumnType_STRING},
			{Name: "upon_approval", Description: "", Type: proto.ColumnType_STRING},
			{Name: "category", Description: "", Type: proto.ColumnType_STRING},
			{Name: "origin_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("OriginID")},
			{Name: "made_sla", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("MadeSLA")},
			{Name: "sys_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "problem_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("ProblemID")},
			{Name: "sla_due", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SLADue")},
			{Name: "correlation_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("CorrelationID")},
			{Name: "resolved_by", Description: "", Type: proto.ColumnType_JSON},
			{Name: "opened_by", Description: "", Type: proto.ColumnType_JSON},
			{Name: "sys_domain", Description: "", Type: proto.ColumnType_JSON},
			{Name: "caller_id", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromField("CorrelatCallerIDionID")},
			{Name: "closed_by", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listServicenowIncidents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_incident.listServicenowIncidents", "connect_error", err)
		return nil, err
	}

	incidentsResponse, err := client.GetIncidents(10)
	if err != nil {
		logger.Error("servicenow_incident.listServicenowIncidents", "query_error", err)
		return nil, err
	}
	for _, incident := range incidentsResponse.Result {
		d.StreamListItem(ctx, incident)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
