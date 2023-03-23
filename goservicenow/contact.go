package goservicenow

type ContactGetResponse struct {
	Result Contact `json:"result"`
}
type ContactListResponse struct {
	Result []Contact `json:"result"`
}
type Contact struct {
	Country                 string `json:"country"`
	CalendarIntegration     string `json:"calendar_integration"`
	LastPositionUpdate      string `json:"last_position_update"`
	LastLoginTime           string `json:"last_login_time"`
	LastLoginDevice         string `json:"last_login_device"`
	Source                  string `json:"source"`
	SysUpdatedOn            string `json:"sys_updated_on"`
	Building                string `json:"building"`
	WebServiceAccessOnly    string `json:"web_service_access_only"`
	Notification            string `json:"notification"`
	SysUpdatedBy            string `json:"sys_updated_by"`
	EnableMultifactorAuthn  string `json:"enable_multifactor_authn"`
	SysCreatedOn            string `json:"sys_created_on"`
	SysDomain               string `json:"sys_domain"`
	AgentStatus             string `json:"agent_status"`
	State                   string `json:"state"`
	Vip                     string `json:"vip"`
	SysCreatedBy            string `json:"sys_created_by"`
	Longitude               string `json:"longitude"`
	Zip                     string `json:"zip"`
	HomePhone               string `json:"home_phone"`
	TimeFormat              string `json:"time_format"`
	LastLogin               string `json:"last_login"`
	DefaultPerspective      string `json:"default_perspective"`
	GeolocationTracked      string `json:"geolocation_tracked"`
	Active                  string `json:"active"`
	TimeSheetPolicy         string `json:"time_sheet_policy"`
	SysDomainPath           string `json:"sys_domain_path"`
	Phone                   string `json:"phone"`
	CostCenter              string `json:"cost_center"`
	Name                    string `json:"name"`
	EmployeeNumber          string `json:"employee_number"`
	Gender                  string `json:"gender"`
	City                    string `json:"city"`
	UserName                string `json:"user_name"`
	FailedAttempts          string `json:"failed_attempts"`
	EduStatus               string `json:"edu_status"`
	Latitude                string `json:"latitude"`
	Roles                   string `json:"roles"`
	Title                   string `json:"title"`
	SysClassName            string `json:"sys_class_name"`
	SysID                   string `json:"sys_id"`
	InternalIntegrationUser string `json:"internal_integration_user"`
	LdapServer              string `json:"ldap_server"`
	MobilePhone             string `json:"mobile_phone"`
	Street                  string `json:"street"`
	Company                 string `json:"company"`
	Department              string `json:"department"`
	FirstName               string `json:"first_name"`
	PreferredLanguage       string `json:"preferred_language"`
	Introduction            string `json:"introduction"`
	Email                   string `json:"email"`
	Manager                 string `json:"manager"`
	LockedOut               string `json:"locked_out"`
	SysModCount             string `json:"sys_mod_count"`
	LastName                string `json:"last_name"`
	Photo                   string `json:"photo"`
	SysTags                 string `json:"sys_tags"`
	MiddleName              string `json:"middle_name"`
	TimeZone                string `json:"time_zone"`
	Schedule                string `json:"schedule"`
	OnSchedule              string `json:"on_schedule"`
	DateFormat              string `json:"date_format"`
	Location                string `json:"location"`
	Account                 string `json:"account"`
}

func (c *Client) GetContacts(limit, offset int) (*ContactListResponse, error) {
	var result ContactListResponse
	err := c.listContacts(limit, offset, &result)
	return &result, err
}

func (c *Client) GetContact(sysId string) (*ContactGetResponse, error) {
	var result ContactGetResponse
	err := c.getContact(sysId, &result)
	return &result, err
}
