package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const CmdbCIServerTableName = "cmdb_ci_server"

//// TABLE DEFINITION

func tableServicenowCmdbCiServer() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_cmdb_ci_server",
		DefaultTransform: transform.FromCamel(),
		Description:      "Manages server configuration item data in the CMDB.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"asset_tag", "attestation_status", "attributes", "category", "chassis_type", "classification", "comments", "correlation_id", "cost_cc", "cpu_name", "cpu_type", "default_gateway", "discovery_source", "dns_domain", "due_in", "environment", "firewall_status", "floppy", "form_factor", "fqdn", "gl_account", "hardware_status", "hardware_substatus", "host_name", "invoice_number", "ip_address", "justification", "lease_id", "mac_address", "model_number", "name", "object_id", "os_domain", "os_service_pack", "os_version", "os", "po_number", "purchase_date", "serial_number", "short_description", "subcategory", "sys_class_name", "sys_class_path", "sys_created_by", "sys_domain_path", "sys_id", "sys_updated_by", "used_for", "warranty_expiration"}),
			Hydrate:    listServicenowObjectsByTable(CmdbCIServerTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(CmdbCIServerTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "asset_tag", Description: "Tag or identifier associated with the server asset.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "asset_tag")},
			{Name: "asset", Description: "Server asset information.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "asset")},
			{Name: "assigned_to", Description: "User or group assigned to the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "assigned_to")},
			{Name: "assigned", Description: "Flag indicating if the server is assigned.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "assigned").Transform(parseDateTime)},
			{Name: "assignment_group", Description: "Group responsible for managing the server assignment.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "assignment_group")},
			{Name: "attestation_score", Description: "Score representing the level of attestation for the server.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "attestation_score")},
			{Name: "attestation_status", Description: "Status of attestation for the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "attestation_status")},
			{Name: "attested_by", Description: "User who attested or confirmed the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "attested_by")},
			{Name: "attested_date", Description: "Date when the server was attested or confirmed.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "attested_date").Transform(parseDateTime)},
			{Name: "attested", Description: "Flag indicating if the server is attested or confirmed.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "attested")},
			{Name: "attributes", Description: "Additional attributes or details of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "attributes")},
			{Name: "business_unit", Description: "Business unit associated with the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "business_unit")},
			{Name: "can_print", Description: "Flag indicating if printing is allowed for the server.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "can_print")},
			{Name: "category", Description: "Categorization or classification of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "category")},
			{Name: "cd_rom", Description: "CD-ROM information associated with the server.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "cd_rom")},
			{Name: "cd_speed", Description: "Speed of the CD-ROM drive.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "cd_speed")},
			{Name: "change_control", Description: "Change control process or record related to the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "change_control")},
			{Name: "chassis_type", Description: "Type or form factor of the server chassis.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "chassis_type")},
			{Name: "checked_in", Description: "Flag indicating if the server is checked in.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "checked_in").Transform(parseDateTime)},
			{Name: "checked_out", Description: "Flag indicating if the server is checked out.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "checked_out").Transform(parseDateTime)},
			{Name: "classification", Description: "Classification or categorization of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "classification")},
			{Name: "comments", Description: "Comments or additional notes about the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "comments")},
			{Name: "company", Description: "Company or organization associated with the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "company")},
			{Name: "correlation_id", Description: "ID or reference used for correlation purposes.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "correlation_id")},
			{Name: "cost_cc", Description: "Cost center associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "cost_cc")},
			{Name: "cost_center", Description: "Cost center associated with the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cost_center")},
			{Name: "cost", Description: "Cost or financial information related to the server.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "cost")},
			{Name: "cpu_core_count", Description: "Number of CPU cores in the server.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_core_count")},
			{Name: "cpu_core_thread", Description: "Number of threads per CPU core.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_core_thread")},
			{Name: "cpu_count", Description: "Number of CPUs in the server.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_count")},
			{Name: "cpu_manufacturer", Description: "Manufacturer of the CPU.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_manufacturer")},
			{Name: "cpu_name", Description: "Name or model of the CPU.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_name")},
			{Name: "cpu_speed", Description: "Speed of the CPU in GHz.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_speed")},
			{Name: "cpu_type", Description: "Type or architecture of the CPU.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "cpu_type")},
			{Name: "default_gateway", Description: "Default gateway configuration for the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "default_gateway")},
			{Name: "delivery_date", Description: "Date when the server was delivered or received.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "delivery_date").Transform(parseDateTime)},
			{Name: "department", Description: "Department or organizational unit associated with the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "department")},
			{Name: "discovery_source", Description: "Source or method of server discovery.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "discovery_source")},
			{Name: "disk_space", Description: "Available disk space on the server.", Type: proto.ColumnType_DOUBLE, Transform: transform.FromP(getFieldFromSObjectMap, "disk_space")},
			{Name: "dns_domain", Description: "DNS domain associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "dns_domain")},
			{Name: "dr_backup", Description: "Flag indicating if the server is a backup for disaster recovery.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "dr_backup")},
			{Name: "due_in", Description: "Remaining time until the server is due.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "due_in")},
			{Name: "due", Description: "Due date or deadline for the server.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "due").Transform(parseDateTime)},
			{Name: "duplicate_of", Description: "Reference to a duplicate server record.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "duplicate_of")},
			{Name: "environment", Description: "Environment or deployment stage of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "environment")},
			{Name: "fault_count", Description: "Number of faults or issues associated with the server.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "fault_count")},
			{Name: "firewall_status", Description: "Status of the firewall configuration for the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "firewall_status")},
			{Name: "first_discovered", Description: "Date when the server was first discovered.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "first_discovered").Transform(parseDateTime)},
			{Name: "floppy", Description: "Floppy drive information associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "floppy")},
			{Name: "form_factor", Description: "Form factor or physical size of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "form_factor")},
			{Name: "fqdn", Description: "Fully Qualified Domain Name (FQDN) of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "fqdn")},
			{Name: "gl_account", Description: "General ledger account associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "gl_account")},
			{Name: "hardware_status", Description: "Status of the hardware components of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "hardware_status")},
			{Name: "hardware_substatus", Description: "Substatus or additional status information for hardware components.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "hardware_substatus")},
			{Name: "host_name", Description: "Host name or identifier of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "host_name")},
			{Name: "install_date", Description: "Date when the server was installed.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "install_date").Transform(parseDateTime)},
			{Name: "install_status", Description: "Status of the server installation.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "install_status")},
			{Name: "internet_facing", Description: "Flag indicating if the server is internet-facing.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "internet_facing")},
			{Name: "invoice_number", Description: "Invoice number related to the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "invoice_number")},
			{Name: "ip_address", Description: "IP address assigned to the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "ip_address")},
			{Name: "justification", Description: "Reason or justification for the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "justification")},
			{Name: "last_discovered", Description: "Date when the server", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "last_discovered").Transform(parseDateTime)},
			{Name: "lease_id", Description: "Identifier or reference for the server lease.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "lease_id")},
			{Name: "life_cycle_stage_status", Description: "Status of the life cycle stage for the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "life_cycle_stage_status")},
			{Name: "life_cycle_stage", Description: "Current stage in the life cycle of the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "life_cycle_stage")},
			{Name: "location", Description: "Location or physical placement of the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "location")},
			{Name: "mac_address", Description: "MAC address of the server's network interface.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "mac_address")},
			{Name: "maintenance_schedule", Description: "Schedule or plan for server maintenance activities.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "maintenance_schedule")},
			{Name: "managed_by_group", Description: "Group responsible for managing the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "managed_by_group")},
			{Name: "managed_by", Description: "User or entity responsible for managing the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "managed_by")},
			{Name: "manufacturer", Description: "Manufacturer of the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "manufacturer")},
			{Name: "model_id", Description: "Identifier or reference for the server model.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "model_id")},
			{Name: "model_number", Description: "Model number or identifier of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "model_number")},
			{Name: "monitor", Description: "Flag indicating if the server is monitored.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "monitor")},
			{Name: "most_frequent_user", Description: "User who frequently uses the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "most_frequent_user")},
			{Name: "name", Description: "Name or identifier of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "object_id", Description: "Identifier or reference for the server object.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "object_id")},
			{Name: "operational_status", Description: "Operational status or condition of the server.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "operational_status")},
			{Name: "order_date", Description: "Date when the server was ordered.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "order_date").Transform(parseDateTime)},
			{Name: "os_address_width", Description: "Address width or architecture of the server's operating system.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "os_address_width")},
			{Name: "os_domain", Description: "Domain or network domain associated with the server's operating system.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "os_domain")},
			{Name: "os_service_pack", Description: "Service pack version of the server's operating system.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "os_service_pack")},
			{Name: "os_version", Description: "Version or release of the server's operating system.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "os_version")},
			{Name: "os", Description: "Operating system installed on the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "os")},
			{Name: "owned_by", Description: "User or entity that owns or possesses the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "owned_by")},
			{Name: "po_number", Description: "Purchase order number related to the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "po_number")},
			{Name: "purchase_date", Description: "Date when the server was purchased.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "purchase_date")},
			{Name: "ram", Description: "Amount of RAM (Random Access Memory) in the server.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "ram")},
			{Name: "schedule", Description: "Schedule or plan associated with the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "schedule")},
			{Name: "serial_number", Description: "Serial number or unique identifier of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "serial_number")},
			{Name: "short_description", Description: "Brief description or summary of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "short_description")},
			{Name: "skip_sync", Description: "Flag indicating if synchronization is skipped for the server.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "skip_sync")},
			{Name: "start_date", Description: "Date when the server was started or activated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "start_date").Transform(parseDateTime)},
			{Name: "subcategory", Description: "Subcategory or further classification of the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "subcategory")},
			{Name: "support_group", Description: "Group responsible for providing support for the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "support_group")},
			{Name: "supported_by", Description: "User or entity responsible for supporting the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "supported_by")},
			{Name: "sys_class_name", Description: "Name of the system class associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_name")},
			{Name: "sys_class_path", Description: "Path or location of the system class associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_path")},
			{Name: "sys_created_by", Description: "User who created the server record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the server record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_domain_path", Description: "Path or location of the system domain associated with the server.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain_path")},
			{Name: "sys_domain", Description: "System domain associated with the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain")},
			{Name: "sys_id", Description: "Unique identifier of the server record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of modifications or updates made to the server record.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "User who last updated the server record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the server record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "unverified", Description: "Flag indicating if the server is unverified or not confirmed.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "unverified")},
			{Name: "used_for", Description: "Purpose or function for which the server is used.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "used_for")},
			{Name: "vendor", Description: "Vendor or supplier of the server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "vendor")},
			{Name: "virtual", Description: "Flag indicating if the server is a virtual machine.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "virtual")},
			{Name: "warranty_expiration", Description: "Expiration date of the server warranty.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "warranty_expiration")},
		}),
	}
}
