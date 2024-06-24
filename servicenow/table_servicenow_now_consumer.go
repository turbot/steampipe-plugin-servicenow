package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowNowConsumer() *plugin.Table {
	return &plugin.Table{
		Name:             "servicenow_now_consumer",
		Description:      "Customer Service Management (CSM) consumer records.",
		DefaultTransform: transform.FromCamel().NullIfEqual(""),
		List: &plugin.ListConfig{
			Hydrate: listServicenowNowConsumers,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowNowConsumers,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "sys_id", Description: "Unique identifier for the consumer.", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "active", Description: "Flag that indicates whether the consumer is active.", Type: proto.ColumnType_BOOL},
			{Name: "business_phone", Description: "Business phone number of the consumer.", Type: proto.ColumnType_STRING},
			{Name: "city", Description: "City in which the consumer resides.", Type: proto.ColumnType_STRING},
			{Name: "country", Description: "Country in which the consumer resides.", Type: proto.ColumnType_STRING},
			{Name: "date_format", Description: "Format in which to display dates.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "Email address of the consumer.", Type: proto.ColumnType_STRING},
			{Name: "fax", Description: "Fax number of the consumer.", Type: proto.ColumnType_STRING},
			{Name: "first_name", Description: "Consumer first name.", Type: proto.ColumnType_STRING},
			{Name: "gender", Description: "Gender of the consumer.", Type: proto.ColumnType_STRING},
			{Name: "home_phone", Description: "Home phone number of the consumer.", Type: proto.ColumnType_STRING},
			{Name: "household", Description: "Sys_id of the record that describes the household characteristics. Located in the Household [servicenow_csm_household] table.", Type: proto.ColumnType_STRING},
			{Name: "last_name", Description: "Consumer last name.", Type: proto.ColumnType_STRING},
			{Name: "middle_name", Description: "Consumer middle name.", Type: proto.ColumnType_STRING},
			{Name: "mobile_phone", Description: "Consumer mobile phone number.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "Consumer full name; first_name+middle_name+last_name.", Type: proto.ColumnType_STRING},
			{Name: "notes", Description: "Notes on consumer.", Type: proto.ColumnType_STRING},
			{Name: "notification", Description: "Indicates whether the consumer should receive notifications.", Type: proto.ColumnType_INT},
			{Name: "number", Description: "Unique number associated with the consumer.", Type: proto.ColumnType_STRING},
			{Name: "photo", Description: "Photo of the consumer.", Type: proto.ColumnType_STRING},
			{Name: "preferred_language", Description: "Consumer primary language.", Type: proto.ColumnType_STRING},
			{Name: "prefix", Description: "Consumer name prefix such as, Dr., Mr., Mrs., or Ms.", Type: proto.ColumnType_STRING},
			{Name: "primary", Description: "Flag that indicates whether this is the primary consumer.", Type: proto.ColumnType_BOOL},
			{Name: "state", Description: "State in which the consumer resides.", Type: proto.ColumnType_STRING},
			{Name: "street", Description: "Consumer street address.", Type: proto.ColumnType_STRING},
			{Name: "suffix", Description: "Consumer name suffix such as Jr., Sr., or II.", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "User that created the consumer record.", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "Date and time the consumer record was originally created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromGo().Transform(parseDateTime)},
			{Name: "sys_domain", Description: "ServiceNow domain in which the consumer information resides.", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "Number of times that the associated consumer information has been modified.", Type: proto.ColumnType_INT},
			{Name: "sys_tags", Description: "System tags.", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "User that last updated the consumer information.", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "Date and time when the consumer information was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromGo().Transform(parseDateTime)},
			{Name: "time_format", Description: "Format in which to display time.", Type: proto.ColumnType_STRING},
			{Name: "time_zone", Description: "Consumer time zone, such as Canada/Central or US/Eastern.", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "Consumer business title such as Manager, Software Developer, or Contractor.", Type: proto.ColumnType_STRING},
			{Name: "user", Description: "Sys_id of the consumer user. Located in the Consumer User [csm_consumer_user] table.", Type: proto.ColumnType_STRING},
			{Name: "zip", Description: "Consumer zip code.", Type: proto.ColumnType_STRING},
		}),
	}
}

//// LIST FUNCTION

func listServicenowNowConsumers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_now_consumer.listServicenowNowConsumers", "connect_error", err)
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
		returnedObject, err := client.NowConsumer.List(limit, offset)
		totalReturned := len(returnedObject.Result)
		if err != nil {
			logger.Error("servicenow_now_consumer.listServicenowNowConsumers", "query_error", err)
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

func getServicenowNowConsumers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_now_consumer.getServicenowNowConsumers", "connect_error", err)
		return nil, err
	}

	sysId := d.EqualsQualString("sys_id")

	returnedObject, err := client.NowContact.Get(sysId)
	if err != nil {
		logger.Error("servicenow_now_consumer.getServicenowNowConsumers", "query_error", err)
		return nil, err
	}

	return returnedObject.Result, err
}
