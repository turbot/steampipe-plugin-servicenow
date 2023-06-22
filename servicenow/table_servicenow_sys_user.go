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
		Description: "Manages user information in the ServiceNow.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"accumulated_roles", "avatar", "city", "correlation_id", "country", "date_format", "department", "edu_status", "email", "employee_number", "first_name", "gender", "hashed_user_id", "home_phone", "introduction", "last_login", "last_login_device", "last_name", "last_password", "middle_name", "mobile_phone", "name", "phone", "photo", "preferred_language", "roles", "source", "state", "street", "sys_class_name", "sys_created_by", "sys_domain_path", "sys_id", "sys_updated_by", "time_format", "time_zone", "title", "transaction_log", "user_name", "user_password", "zip"}),
			Hydrate:    listServicenowObjectsByTable(SysUserTableName, nil),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowObjectbyID(SysUserTableName),
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "accumulated_roles", Description: "Roles accumulated by the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "accumulated_roles")},
			{Name: "active", Description: "Indicates if the user is active.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "active")},
			{Name: "avatar", Description: "Avatar image for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "avatar")},
			{Name: "building", Description: "Building where the user is located.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "building")},
			{Name: "calendar_integration", Description: "Integration settings for the user's calendar.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "calendar_integration")},
			{Name: "city", Description: "City where the user is located.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "city")},
			{Name: "company", Description: "Company associated with the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "company")},
			{Name: "correlation_id", Description: "Correlation ID for user records.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "correlation_id")},
			{Name: "cost_center", Description: "Cost center associated with the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "cost_center")},
			{Name: "country", Description: "Country where the user is located.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "country")},
			{Name: "date_format", Description: "Date format preference for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "date_format")},
			{Name: "default_perspective", Description: "Default perspective for the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "default_perspective")},
			{Name: "department", Description: "Department associated with the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "department")},
			{Name: "edu_status", Description: "Educational status of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "edu_status")},
			{Name: "email", Description: "Email address of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "email")},
			{Name: "employee_number", Description: "Employee number of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "employee_number")},
			{Name: "enable_multifactor_authn", Description: "Indicates if multi-factor authentication is enabled for the user.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "enable_multifactor_authn")},
			{Name: "failed_attempts", Description: "Number of failed login attempts for the user.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "failed_attempts")},
			{Name: "first_name", Description: "First name of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "first_name")},
			{Name: "gender", Description: "Gender of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "gender")},
			{Name: "hashed_user_id", Description: "Hashed user ID for user records.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "hashed_user_id")},
			{Name: "home_phone", Description: "Home phone number of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "home_phone")},
			{Name: "hr_integration_source", Description: "Source of HR integration for the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "hr_integration_source")},
			{Name: "internal_integration_user", Description: "Indicates if the user is an internal integration user.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "internal_integration_user")},
			{Name: "introduction", Description: "Introduction or bio of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "introduction")},
			{Name: "last_login", Description: "Last login date for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_login")},
			{Name: "last_login_device", Description: "Device used for the user's last login.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_login_device")},
			{Name: "last_login_time", Description: "Time of the user's last login.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "last_login_time").Transform(parseDateTime)},
			{Name: "last_name", Description: "Last name of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_name")},
			{Name: "last_password", Description: "Last password used by the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "last_password")},
			{Name: "ldap_server", Description: "LDAP server associated with the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "ldap_server")},
			{Name: "location", Description: "Location where the user is located.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "location")},
			{Name: "locked_out", Description: "Indicates if the user is locked out.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "locked_out")},
			{Name: "manager", Description: "Manager of the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "manager")},
			{Name: "middle_name", Description: "Middle name of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "middle_name")},
			{Name: "mobile_phone", Description: "Mobile phone number of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "mobile_phone")},
			{Name: "name", Description: "Full name of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "name")},
			{Name: "notification", Description: "Notification settings for the user.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "notification")},
			{Name: "password_needs_reset", Description: "Indicates if the user's password needs to be reset.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "password_needs_reset")},
			{Name: "phone", Description: "Phone number of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "phone")},
			{Name: "photo", Description: "Profile photo of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "photo")},
			{Name: "preferred_language", Description: "Preferred language for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "preferred_language")},
			{Name: "roles", Description: "Roles assigned to the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "roles")},
			{Name: "schedule", Description: "Schedule associated with the user.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "schedule")},
			{Name: "source", Description: "Source of the user record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "source")},
			{Name: "state", Description: "State or province where the user is located.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "state")},
			{Name: "street", Description: "Street address of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "street")},
			{Name: "sys_class_name", Description: "System class name of the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_class_name")},
			{Name: "sys_created_by", Description: "User who created the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_by")},
			{Name: "sys_created_on", Description: "Date and time when the record was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_created_on").Transform(parseDateTime)},
			{Name: "sys_domain", Description: "Domain associated with the record.", Type: proto.ColumnType_JSON, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain")},
			{Name: "sys_domain_path", Description: "Domain path associated with the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_domain_path")},
			{Name: "sys_id", Description: "Unique system identifier for the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_id")},
			{Name: "sys_mod_count", Description: "Number of times the record was modified.", Type: proto.ColumnType_INT, Transform: transform.FromP(getFieldFromSObjectMap, "sys_mod_count")},
			{Name: "sys_updated_by", Description: "User who last updated the record.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_by")},
			{Name: "sys_updated_on", Description: "Date and time when the record was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromP(getFieldFromSObjectMap, "sys_updated_on").Transform(parseDateTime)},
			{Name: "time_format", Description: "Time format preference for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "time_format")},
			{Name: "time_zone", Description: "Time zone preference for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "time_zone")},
			{Name: "title", Description: "Title or job role of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "title")},
			{Name: "transaction_log", Description: "Transaction log associated with the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "transaction_log")},
			{Name: "user_name", Description: "User name of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user_name")},
			{Name: "user_password", Description: "User password of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "user_password")},
			{Name: "vip", Description: "Indicates if the user is a VIP.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "vip")},
			{Name: "web_service_access_only", Description: "Indicates if the user has access only through web services.", Type: proto.ColumnType_BOOL, Transform: transform.FromP(getFieldFromSObjectMap, "web_service_access_only")},
			{Name: "zip", Description: "ZIP or postal code of the user.", Type: proto.ColumnType_STRING, Transform: transform.FromP(getFieldFromSObjectMap, "zip")},
		},
	}
}
