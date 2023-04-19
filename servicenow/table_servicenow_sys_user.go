package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const SysUserTableName = "sys_user"

//// TABLE DEFINITION

func tableServicenowSysUser() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sys_user",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowObjectsByTable(SysUserTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "accumulated_roles", Description: "Accumulated roles.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "accumulated_roles")},
			{Name: "active", Description: "Inactive users do not show in user choice lists.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "active")},
			{Name: "avatar", Description: "Avatar.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "avatar")},
			{Name: "building", Description: "Building.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "building")},
			{Name: "calendar_integration", Description: "Calendar integration.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "calendar_integration")},
			{Name: "city", Description: "City.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "city")},
			{Name: "company", Description: "Company.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "company")},
			{Name: "correlation_id", Description: "Correlation ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "correlation_id")},
			{Name: "cost_center", Description: "Cost center.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cost_center")},
			{Name: "country", Description: "Country code.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "country")},
			{Name: "date_format", Description: "Display dates with this format (blank means system default). ", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "date_format")},
			{Name: "default_perspective", Description: "Default perspective.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "default_perspective")},
			{Name: "department", Description: "Department.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "department")},
			{Name: "edu_status", Description: "EDU Status.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "edu_status")},
			{Name: "email", Description: "Email.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "email")},
			{Name: "employee_number", Description: "Employee number.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "employee_number")},
			{Name: "enable_multifactor_authn", Description: "Enable Multifactor Authentication.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "enable_multifactor_authn")},
			{Name: "failed_attempts", Description: "Failed login attempts.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "failed_attempts")},
			{Name: "first_name", Description: "First name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "first_name")},
			{Name: "gender", Description: "Gender.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "gender")},
			{Name: "hashed_user_id", Description: "Hashed User ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "hashed_user_id")},
			{Name: "home_phone", Description: "Home phone.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "home_phone")},
			{Name: "hr_integration_source", Description: "HR Integration source.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "hr_integration_source")},
			{Name: "internal_integration_user", Description: "Internal Integration User.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "internal_integration_user")},
			{Name: "introduction", Description: "Prefix.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "introduction")},
			{Name: "last_login", Description: "Last login.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_login")},
			{Name: "last_login_device", Description: "Last login device.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_login_device")},
			{Name: "last_login_time", Description: "Last login time.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "last_login_time").Transform(parseDateTime)},
			{Name: "last_name", Description: "Last name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_name")},
			{Name: "last_password", Description: "Last password.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_password")},
			{Name: "ldap_server", Description: "LDAP server.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "ldap_server")},
			{Name: "location", Description: "Location.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "location")},
			{Name: "locked_out", Description: "When checked, user cannot login.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "locked_out")},
			{Name: "manager", Description: "Manager.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "manager")},
			{Name: "middle_name", Description: "Middle name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "middle_name")},
			{Name: "mobile_phone", Description: "Mobile phone.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "mobile_phone")},
			{Name: "name", Description: "Name.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "notification", Description: "Enable or disable notifications for this user ie. email, SMS etc.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "notification")},
			{Name: "password_needs_reset", Description: "User will be prompted to change password at next login.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "password_needs_reset")},
			{Name: "phone", Description: "Business phone.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "phone")},
			{Name: "photo", Description: "Photo.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "photo")},
			{Name: "preferred_language", Description: "Language.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "preferred_language")},
			{Name: "roles", Description: "Roles.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "roles")},
			{Name: "schedule", Description: "Schedule.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "schedule")},
			{Name: "source", Description: "Source.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "source")},
			{Name: "state", Description: "State / Province.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "state")},
			{Name: "street", Description: "Street.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "street")},
			{Name: "sys_class_name", Description: "Class.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_name")},
			{Name: "sys_created_by", Description: "Created by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_domain", Description: "Domain to which the user belongs.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain")},
			{Name: "sys_domain_path", Description: "Domain Path.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain_path")},
			{Name: "sys_id", Description: "Sys ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Updates.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "Updated by.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "time_format", Description: "Display times with this format (blank means system default). ", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "time_format")},
			{Name: "time_zone", Description: "Time zone.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "time_zone")},
			{Name: "title", Description: "Title.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "title")},
			{Name: "transaction_log", Description: "Transaction log.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "transaction_log")},
			{Name: "user_name", Description: "User ID.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user_name")},
			{Name: "user_password", Description: "Password.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user_password")},
			{Name: "vip", Description: "VIP.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "vip")},
			{Name: "web_service_access_only", Description: "Prevent user from accessing UI, and require a SOAP role to make API protocol calls (such as SOAP and WSDL requests).", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "web_service_access_only")},
			{Name: "zip", Description: "Zip / Postal code.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "zip")},
		},
	}
}
