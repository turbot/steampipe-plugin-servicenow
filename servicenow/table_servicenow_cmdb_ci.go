package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowCmdbCi() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_cmdb_ci",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowCmdbCis,
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "attested_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "skip_sync", Description: "", Type: proto.ColumnType_STRING},
			{Name: "operational_status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "attestation_score", Description: "", Type: proto.ColumnType_STRING},
			{Name: "discovery_source", Description: "", Type: proto.ColumnType_STRING},
			{Name: "first_discovered", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "due_in", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "install_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "gl_account", Description: "", Type: proto.ColumnType_STRING},
			{Name: "invoice_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "warranty_expiration", Description: "", Type: proto.ColumnType_STRING},
			{Name: "asset_tag", Description: "", Type: proto.ColumnType_STRING},
			{Name: "fqdn", Description: "", Type: proto.ColumnType_STRING},
			{Name: "change_control", Description: "", Type: proto.ColumnType_STRING},
			{Name: "owned_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "checked_out", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain_path", Description: "", Type: proto.ColumnType_STRING},
			{Name: "business_unit", Description: "", Type: proto.ColumnType_STRING},
			{Name: "delivery_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "install_status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "cost_center", Description: "", Type: proto.ColumnType_STRING},
			{Name: "attested_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "supported_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "assigned", Description: "", Type: proto.ColumnType_STRING},
			{Name: "life_cycle_stage", Description: "", Type: proto.ColumnType_STRING},
			{Name: "purchase_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "subcategory", Description: "", Type: proto.ColumnType_STRING},
			{Name: "short_description", Description: "", Type: proto.ColumnType_STRING},
			{Name: "assignment_group", Description: "", Type: proto.ColumnType_STRING},
			{Name: "managed_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "managed_by_group", Description: "", Type: proto.ColumnType_STRING},
			{Name: "can_print", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_discovered", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_class_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "manufacturer", Description: "", Type: proto.ColumnType_STRING},
			{Name: "po_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "checked_in", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_class_path", Description: "", Type: proto.ColumnType_STRING},
			{Name: "life_cycle_stage_status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "mac_address", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vendor", Description: "", Type: proto.ColumnType_STRING},
			{Name: "company", Description: "", Type: proto.ColumnType_STRING},
			{Name: "justification", Description: "", Type: proto.ColumnType_STRING},
			{Name: "model_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "department", Description: "", Type: proto.ColumnType_STRING},
			{Name: "assigned_to", Description: "", Type: proto.ColumnType_STRING},
			{Name: "start_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "comments", Description: "", Type: proto.ColumnType_STRING},
			{Name: "cost", Description: "", Type: proto.ColumnType_STRING},
			{Name: "attestation_status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "monitor", Description: "", Type: proto.ColumnType_STRING},
			{Name: "serial_number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "duplicate_of", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_tags", Description: "", Type: proto.ColumnType_STRING},
			{Name: "cost_cc", Description: "", Type: proto.ColumnType_STRING},
			{Name: "order_date", Description: "", Type: proto.ColumnType_STRING},
			{Name: "schedule", Description: "", Type: proto.ColumnType_STRING},
			{Name: "environment", Description: "", Type: proto.ColumnType_STRING},
			{Name: "due", Description: "", Type: proto.ColumnType_STRING},
			{Name: "attested", Description: "", Type: proto.ColumnType_STRING},
			{Name: "unverified", Description: "", Type: proto.ColumnType_STRING},
			{Name: "attributes", Description: "", Type: proto.ColumnType_STRING},
			{Name: "category", Description: "", Type: proto.ColumnType_STRING},
			{Name: "fault_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "lease_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("LeaseID")},
			{Name: "correlation_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("CorrelationID")},
			{Name: "model_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("ModelID")},
			{Name: "ip_address", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("IPAddress")},
			{Name: "dns_domain", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("DNSDomain")},
			{Name: "sys_domain", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listServicenowCmdbCis(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_cmdb_ci.listServicenowCmdbCis", "connect_error", err)
		return nil, err
	}

	cmdb_cisResponse, err := client.GetCmdbCIs(10)
	if err != nil {
		logger.Error("servicenow_cmdb_ci.listServicenowCmdbCis", "query_error", err)
		return nil, err
	}
	for _, cmdb_ci := range cmdb_cisResponse.Result {
		d.StreamListItem(ctx, cmdb_ci)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
