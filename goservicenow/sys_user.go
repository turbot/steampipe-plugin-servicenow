package goservicenow

import "fmt"

type SysUser struct {
	Result []Result `json:"result"`
}

type SysDomain struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type Result struct {
	CalendarIntegration     string    `json:"calendar_integration"`
	Country                 string    `json:"country"`
	UserPassword            string    `json:"user_password"`
	LastLoginTime           string    `json:"last_login_time"`
	Source                  string    `json:"source"`
	SysUpdatedOn            string    `json:"sys_updated_on"`
	Building                string    `json:"building"`
	WebServiceAccessOnly    string    `json:"web_service_access_only"`
	Notification            string    `json:"notification"`
	EnableMultifactorAuthn  string    `json:"enable_multifactor_authn"`
	SysUpdatedBy            string    `json:"sys_updated_by"`
	SysCreatedOn            string    `json:"sys_created_on"`
	SysDomain               SysDomain `json:"sys_domain"`
	State                   string    `json:"state"`
	Vip                     string    `json:"vip"`
	SysCreatedBy            string    `json:"sys_created_by"`
	Zip                     string    `json:"zip"`
	HomePhone               string    `json:"home_phone"`
	TimeFormat              string    `json:"time_format"`
	LastLogin               string    `json:"last_login"`
	DefaultPerspective      string    `json:"default_perspective"`
	Active                  string    `json:"active"`
	SysDomainPath           string    `json:"sys_domain_path"`
	CostCenter              string    `json:"cost_center"`
	Phone                   string    `json:"phone"`
	Name                    string    `json:"name"`
	EmployeeNumber          string    `json:"employee_number"`
	PasswordNeedsReset      string    `json:"password_needs_reset"`
	Gender                  string    `json:"gender"`
	City                    string    `json:"city"`
	FailedAttempts          string    `json:"failed_attempts"`
	UserName                string    `json:"user_name"`
	Roles                   string    `json:"roles"`
	Title                   string    `json:"title"`
	SysClassName            string    `json:"sys_class_name"`
	SysID                   string    `json:"sys_id"`
	InternalIntegrationUser string    `json:"internal_integration_user"`
	LdapServer              string    `json:"ldap_server"`
	MobilePhone             string    `json:"mobile_phone"`
	Street                  string    `json:"street"`
	// Company                 string    `json:"company"`
	Department        string `json:"department"`
	FirstName         string `json:"first_name"`
	Email             string `json:"email"`
	Introduction      string `json:"introduction"`
	PreferredLanguage string `json:"preferred_language"`
	Manager           string `json:"manager"`
	LockedOut         string `json:"locked_out"`
	SysModCount       string `json:"sys_mod_count"`
	LastName          string `json:"last_name"`
	Photo             string `json:"photo"`
	Avatar            string `json:"avatar"`
	MiddleName        string `json:"middle_name"`
	SysTags           string `json:"sys_tags"`
	TimeZone          string `json:"time_zone"`
	Schedule          string `json:"schedule"`
	DateFormat        string `json:"date_format"`
	// Location          string `json:"location"`
}

func (c *Client) GetSysUsers(limit int) (*SysUser, error) {
	var result SysUser
	err := c.getTable("incident", limit, &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &result, nil
}
