package rest

type OAuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type Incident struct {
	Result []struct {
		Parent              string `json:"parent"`
		MadeSLA             string `json:"made_sla"`
		CausedBy            string `json:"caused_by"`
		WatchList           string `json:"watch_list"`
		UponReject          string `json:"upon_reject"`
		SysUpdatedOn        string `json:"sys_updated_on"`
		ChildIncidents      string `json:"child_incidents"`
		HoldReason          string `json:"hold_reason"`
		OriginTable         string `json:"origin_table"`
		TaskEffectiveNumber string `json:"task_effective_number"`
		ApprovalHistory     string `json:"approval_history"`
		Number              string `json:"number"`
		ResolvedBy          struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"resolved_by"`
		SysUpdatedBy string `json:"sys_updated_by"`
		OpenedBy     struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"opened_by"`
		UserInput    string `json:"user_input"`
		SysCreatedOn string `json:"sys_created_on"`
		SysDomain    struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"sys_domain"`
		State        string `json:"state"`
		RouteReason  string `json:"route_reason"`
		SysCreatedBy string `json:"sys_created_by"`
		Knowledge    string `json:"knowledge"`
		Order        string `json:"order"`
		CalendarStc  string `json:"calendar_stc"`
		ClosedAt     string `json:"closed_at"`
		// CmdbCi       struct {
		// 	Link  string `json:"link"`
		// 	Value string `json:"value"`
		// } `json:"cmdb_ci"`
		DeliveryPlan  string `json:"delivery_plan"`
		Contract      string `json:"contract"`
		Impact        string `json:"impact"`
		Active        string `json:"active"`
		WorkNotesList string `json:"work_notes_list"`
		// BusinessService struct {
		// 	Link  string `json:"link"`
		// 	Value string `json:"value"`
		// } `json:"business_service"`
		BusinessImpact   string `json:"business_impact"`
		Priority         string `json:"priority"`
		SysDomainPath    string `json:"sys_domain_path"`
		Rfc              string `json:"rfc"`
		TimeWorked       string `json:"time_worked"`
		ExpectedStart    string `json:"expected_start"`
		OpenedAt         string `json:"opened_at"`
		BusinessDuration string `json:"business_duration"`
		GroupList        string `json:"group_list"`
		WorkEnd          string `json:"work_end"`
		CallerID         struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"caller_id"`
		ReopenedTime       string `json:"reopened_time"`
		ResolvedAt         string `json:"resolved_at"`
		ApprovalSet        string `json:"approval_set"`
		Subcategory        string `json:"subcategory"`
		WorkNotes          string `json:"work_notes"`
		UniversalRequest   string `json:"universal_request"`
		ShortDescription   string `json:"short_description"`
		CloseCode          string `json:"close_code"`
		CorrelationDisplay string `json:"correlation_display"`
		DeliveryTask       string `json:"delivery_task"`
		WorkStart          string `json:"work_start"`
		// AssignmentGroup    struct {
		// 	Link  string `json:"link"`
		// 	Value string `json:"value"`
		// } `json:"assignment_group"`
		AdditionalAssigneeList string `json:"additional_assignee_list"`
		BusinessStc            string `json:"business_stc"`
		Cause                  string `json:"cause"`
		Description            string `json:"description"`
		OriginID               string `json:"origin_id"`
		CalendarDuration       string `json:"calendar_duration"`
		CloseNotes             string `json:"close_notes"`
		Notify                 string `json:"notify"`
		ServiceOffering        string `json:"service_offering"`
		SysClassName           string `json:"sys_class_name"`
		ClosedBy               struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"closed_by"`
		FollowUp       string `json:"follow_up"`
		ParentIncident string `json:"parent_incident"`
		SysID          string `json:"sys_id"`
		ContactType    string `json:"contact_type"`
		ReopenedBy     string `json:"reopened_by"`
		IncidentState  string `json:"incident_state"`
		Urgency        string `json:"urgency"`
		ProblemID      string `json:"problem_id"`
		// Company        struct {
		// 	Link  string `json:"link"`
		// 	Value string `json:"value"`
		// } `json:"company"`
		ReassignmentCount string `json:"reassignment_count"`
		ActivityDue       string `json:"activity_due"`
		// AssignedTo        struct {
		// 	Link  string `json:"link"`
		// 	Value string `json:"value"`
		// } `json:"assigned_to"`
		Severity             string `json:"severity"`
		Comments             string `json:"comments"`
		Approval             string `json:"approval"`
		SLADue               string `json:"sla_due"`
		CommentsAndWorkNotes string `json:"comments_and_work_notes"`
		DueDate              string `json:"due_date"`
		SysModCount          string `json:"sys_mod_count"`
		ReopenCount          string `json:"reopen_count"`
		SysTags              string `json:"sys_tags"`
		Escalation           string `json:"escalation"`
		UponApproval         string `json:"upon_approval"`
		CorrelationID        string `json:"correlation_id"`
		// Location             string `json:"location"`
		Category string `json:"category"`
	} `json:"result"`
}

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
