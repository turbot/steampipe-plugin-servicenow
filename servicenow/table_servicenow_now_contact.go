package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowNowContact() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_now_contact",
		Description:      "Customer Service Management (CSM) contact records.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate: listServicenowNowContacts,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowNowContacts,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "Unique identifier for the associated contact record.", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "account", Description: "Sys_id of the account record to which the contact is associated; Account [customer_account] table.", Type: proto.ColumnType_STRING},
			{Name: "active", Description: "Flag that indicates whether the contact is active within the system.", Type: proto.ColumnType_BOOL},
			{Name: "agent_status", Description: "Status of the agent.", Type: proto.ColumnType_STRING},
			{Name: "building", Description: "Sys_id of the record that describes the building in which the contact resides; Building [cmn_building] table.", Type: proto.ColumnType_STRING},
			{Name: "calendar_integration", Description: "Calendar application that the contact uses.", Type: proto.ColumnType_INT},
			{Name: "city", Description: "City in which the contact resides.", Type: proto.ColumnType_STRING},
			{Name: "company", Description: "Sys_id of the company record to which the contact is associated; Company [core_company] table.", Type: proto.ColumnType_STRING},
			{Name: "cost_center", Description: "Sys_id of the cost center associated with the contact; Cost Center [cmn_cost_center] table.", Type: proto.ColumnType_STRING},
			{Name: "country", Description: "Country code of the country in which the contact resides.", Type: proto.ColumnType_STRING},
			{Name: "date_format", Description: "Format in which to display dates to contacts.", Type: proto.ColumnType_STRING},
			{Name: "default_perspective", Description: "Sys_id of the default perspective for the contact. Located in the Menu List [sys_perspective] table.", Type: proto.ColumnType_STRING},
			{Name: "department", Description: "Sys_id of the department associated with the contact. Located in the Department [cmn_department] table.", Type: proto.ColumnType_STRING},
			{Name: "edu_status", Description: "Education status of the associated contact.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "Contact email address.", Type: proto.ColumnType_STRING},
			{Name: "employee_number", Description: "Contact employee number.", Type: proto.ColumnType_STRING},
			{Name: "enable_multifactor_authn", Description: "Flag that indicates whether multifactor authorization is required for the contact to log in to the service portal.", Type: proto.ColumnType_BOOL},
			{Name: "failed_attempts", Description: "Number of failed log in attempts.", Type: proto.ColumnType_INT},
			{Name: "first_name", Description: "Contact first name.", Type: proto.ColumnType_STRING},
			{Name: "gender", Description: "Contact gender.", Type: proto.ColumnType_STRING},
			{Name: "geolocation_tracked", Description: "Flag that indicates whether the contact location is obtained through geotracking.", Type: proto.ColumnType_BOOL},
			{Name: "home_phone", Description: "Contact home phone number.", Type: proto.ColumnType_STRING},
			{Name: "internal_integration_user", Description: "Flag that indicates whether the contact is an internal integration user.", Type: proto.ColumnType_BOOL},
			{Name: "introduction", Description: "Introduction", Type: proto.ColumnType_STRING},
			{Name: "last_login_device", Description: "Date on which the contact last logged into the system.", Type: proto.ColumnType_STRING},
			{Name: "last_login_time", Description: "Date and time the contact logged in to the system.", Type: proto.ColumnType_STRING},
			{Name: "last_login", Description: "Date on which the contact last logged into the system.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "last_name", Description: "Contact last name.", Type: proto.ColumnType_STRING},
			{Name: "last_position_update", Description: "Date and time the last position was updated.", Type: proto.ColumnType_STRING},
			{Name: "latitude", Description: "Latitude coordinate of the contact.", Type: proto.ColumnType_DOUBLE},
			{Name: "ldap_server", Description: "Sys_id of the LDAP server used by the contact to last log in to the system; LDAP Server [ldap_server_config] table.", Type: proto.ColumnType_STRING},
			{Name: "location", Description: "Sys_id of the record that describes the location of the contact; Location [cmn_location] table.", Type: proto.ColumnType_STRING},
			{Name: "locked_out", Description: "Flag that indicates if the contact is locked-out.", Type: proto.ColumnType_BOOL},
			{Name: "longitude", Description: "Longitude coordinate of the contact.", Type: proto.ColumnType_DOUBLE},
			{Name: "manager", Description: "Sys_id of the record that describes the direct supervisor of the contact; User [sys_user] table.", Type: proto.ColumnType_STRING},
			{Name: "middle_name", Description: "Contact middle name.", Type: proto.ColumnType_STRING},
			{Name: "mobile_phone", Description: "Contact mobile phone number.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Contact full name.", Type: proto.ColumnType_STRING},
			{Name: "notification", Description: "Indicates whether the contact should receive notifications.", Type: proto.ColumnType_INT},
			{Name: "on_schedule", Description: "Indicates the timeliness of dispatched service personnel.", Type: proto.ColumnType_STRING},
			{Name: "phone", Description: "Contact business phone number.", Type: proto.ColumnType_STRING},
			{Name: "photo", Description: "Photo image of the contact.", Type: proto.ColumnType_STRING},
			{Name: "preferred_language", Description: "Country code of the contact primary language.", Type: proto.ColumnType_STRING},
			{Name: "roles", Description: "List of user roles associated with the contact.", Type: proto.ColumnType_STRING},
			{Name: "schedule", Description: "Sys_id of the record that describes the work schedule for the associated contact; Schedule [cmn_schedule] table.", Type: proto.ColumnType_STRING},
			{Name: "source", Description: "Source of the contact.", Type: proto.ColumnType_STRING},
			{Name: "state", Description: "State in which the contact resides.", Type: proto.ColumnType_STRING},
			{Name: "street", Description: "Contact street address.", Type: proto.ColumnType_STRING},
			{Name: "sys_class_name", Description: "Table that contains the contact record.", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "User that originally created the associated contact record.", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "Data and time the associated contact was originally created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromGo().Transform(parseDateTime)},
			{Name: "sys_domain_path", Description: "Contact record domain path.", Type: proto.ColumnType_STRING},
			{Name: "sys_domain", Description: "ServiceNow instance domain of the associated contact record.", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "Number of times that the associated contact record has been modified.", Type: proto.ColumnType_INT},
			{Name: "sys_tags", Description: "System tags.", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "User that last updated the associated contact information.", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "Data and time the associated contact information was updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromGo().Transform(parseDateTime)},
			{Name: "time_format", Description: "Format in which to display time.", Type: proto.ColumnType_STRING},
			{Name: "time_sheet_policy", Description: "Sys_id of the record that contains the time sheet policy for the associated contact; Time Sheet Policy [time_sheet_policy] table.", Type: proto.ColumnType_STRING},
			{Name: "time_zone", Description: "Time zone in which the contact resides, such as Canada/Central or US/Eastern.", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "Contact business title such as Manager, Software Developer, or Contractor.", Type: proto.ColumnType_STRING},
			{Name: "user_name", Description: "Contact user ID.", Type: proto.ColumnType_STRING},
			{Name: "vip", Description: "Flag that indicates whether the associated contact has VIP status.", Type: proto.ColumnType_BOOL},
			{Name: "web_service_access_only", Description: "Flag that indicates whether the contact can only access services through the web.", Type: proto.ColumnType_BOOL},
			{Name: "zip", Description: "Contact zip code.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listServicenowNowContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_now_contact.listServicenowNowContacts", "connect_error", err)
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
		returnedObject, err := client.NowContact.List(limit, offset)
		totalReturned := len(returnedObject.Result)
		if err != nil {
			logger.Error("servicenow_now_contact.listServicenowNowContacts", "query_error", err)
			return nil, err
		}
		for _, object := range returnedObject.Result {
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

func getServicenowNowContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_now_contact.getServicenowNowContacts", "connect_error", err)
		return nil, err
	}

	sysId := d.EqualsQualString("sys_id")

	returnedObject, err := client.NowContact.Read(sysId)
	if err != nil {
		logger.Error("servicenow_now_contact.getServicenowNowContacts", "query_error", err)
		return nil, err
	}

	return returnedObject.Result, err
}
