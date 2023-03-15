package rest

import "fmt"

type CmdbCI struct {
	Result []struct {
		AttestedDate      string `json:"attested_date"`
		SkipSync          string `json:"skip_sync"`
		OperationalStatus string `json:"operational_status"`
		SysUpdatedOn      string `json:"sys_updated_on"`
		AttestationScore  string `json:"attestation_score"`
		DiscoverySource   string `json:"discovery_source"`
		FirstDiscovered   string `json:"first_discovered"`
		SysUpdatedBy      string `json:"sys_updated_by"`
		DueIn             string `json:"due_in"`
		SysCreatedOn      string `json:"sys_created_on"`
		SysDomain         struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"sys_domain"`
		InstallDate        string `json:"install_date"`
		GlAccount          string `json:"gl_account"`
		InvoiceNumber      string `json:"invoice_number"`
		SysCreatedBy       string `json:"sys_created_by"`
		WarrantyExpiration string `json:"warranty_expiration"`
		AssetTag           string `json:"asset_tag"`
		Fqdn               string `json:"fqdn"`
		ChangeControl      string `json:"change_control"`
		OwnedBy            string `json:"owned_by"`
		CheckedOut         string `json:"checked_out"`
		SysDomainPath      string `json:"sys_domain_path"`
		BusinessUnit       string `json:"business_unit"`
		DeliveryDate       string `json:"delivery_date"`
		// TODO: MaintenanceSchedule is struct when filled, but string when empty
		// MaintenanceSchedule  string `json:"maintenance_schedule"`
		// MaintenanceSchedule struct {
		// 	Link  string `json:"link"`
		// 	Value string `json:"value"`
		// } `json:"maintenance_schedule"`
		InstallStatus        string `json:"install_status"`
		CostCenter           string `json:"cost_center"`
		AttestedBy           string `json:"attested_by"`
		SupportedBy          string `json:"supported_by"`
		DNSDomain            string `json:"dns_domain"`
		Name                 string `json:"name"`
		Assigned             string `json:"assigned"`
		LifeCycleStage       string `json:"life_cycle_stage"`
		PurchaseDate         string `json:"purchase_date"`
		Subcategory          string `json:"subcategory"`
		ShortDescription     string `json:"short_description"`
		AssignmentGroup      string `json:"assignment_group"`
		ManagedBy            string `json:"managed_by"`
		ManagedByGroup       string `json:"managed_by_group"`
		CanPrint             string `json:"can_print"`
		LastDiscovered       string `json:"last_discovered"`
		SysClassName         string `json:"sys_class_name"`
		Manufacturer         string `json:"manufacturer"`
		SysID                string `json:"sys_id"`
		PoNumber             string `json:"po_number"`
		CheckedIn            string `json:"checked_in"`
		SysClassPath         string `json:"sys_class_path"`
		LifeCycleStageStatus string `json:"life_cycle_stage_status"`
		MacAddress           string `json:"mac_address"`
		Vendor               string `json:"vendor"`
		Company              string `json:"company"`
		Justification        string `json:"justification"`
		ModelNumber          string `json:"model_number"`
		Department           string `json:"department"`
		AssignedTo           string `json:"assigned_to"`
		StartDate            string `json:"start_date"`
		Comments             string `json:"comments"`
		Cost                 string `json:"cost"`
		AttestationStatus    string `json:"attestation_status"`
		SysModCount          string `json:"sys_mod_count"`
		Monitor              string `json:"monitor"`
		SerialNumber         string `json:"serial_number"`
		IPAddress            string `json:"ip_address"`
		ModelID              string `json:"model_id"`
		DuplicateOf          string `json:"duplicate_of"`
		SysTags              string `json:"sys_tags"`
		CostCc               string `json:"cost_cc"`
		OrderDate            string `json:"order_date"`
		Schedule             string `json:"schedule"`
		// SupportGroup         string `json:"support_group"`
		Environment   string `json:"environment"`
		Due           string `json:"due"`
		Attested      string `json:"attested"`
		CorrelationID string `json:"correlation_id"`
		Unverified    string `json:"unverified"`
		Attributes    string `json:"attributes"`
		// TODO: Location is struct when filled, but string when empty
		// Location             string `json:"location"`
		// Asset      string `json:"asset"`
		Category   string `json:"category"`
		FaultCount string `json:"fault_count"`
		LeaseID    string `json:"lease_id"`
	} `json:"result"`
}

func (c *Client) GetCmdbCIs(limit int) (*CmdbCI, error) {
	var result CmdbCI
	err := c.getTable("cmdb_ci", limit, &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &result, nil
}
