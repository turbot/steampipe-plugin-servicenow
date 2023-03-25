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
		Name:        "servicenow_now_consumer",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowNowConsumers,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowNowConsumers,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "active", Description: "", Type: proto.ColumnType_STRING},
			{Name: "business_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "city", Description: "", Type: proto.ColumnType_STRING},
			{Name: "country", Description: "", Type: proto.ColumnType_STRING},
			{Name: "date_format", Description: "", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "", Type: proto.ColumnType_STRING},
			{Name: "fax", Description: "", Type: proto.ColumnType_STRING},
			{Name: "first_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "gender", Description: "", Type: proto.ColumnType_STRING},
			{Name: "home_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "household", Description: "", Type: proto.ColumnType_STRING},
			{Name: "last_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "middle_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "mobile_phone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "notes", Description: "", Type: proto.ColumnType_STRING},
			{Name: "notification", Description: "", Type: proto.ColumnType_STRING},
			{Name: "number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "photo", Description: "", Type: proto.ColumnType_STRING},
			{Name: "preferred_language", Description: "", Type: proto.ColumnType_STRING},
			{Name: "prefix", Description: "", Type: proto.ColumnType_STRING},
			{Name: "primary", Description: "", Type: proto.ColumnType_STRING},
			{Name: "state", Description: "", Type: proto.ColumnType_STRING},
			{Name: "street", Description: "", Type: proto.ColumnType_STRING},
			{Name: "suffix", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_created_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_domain", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_mod_count", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_tags", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "sys_updated_on", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_format", Description: "", Type: proto.ColumnType_STRING},
			{Name: "time_zone", Description: "", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "", Type: proto.ColumnType_STRING},
			{Name: "user", Description: "", Type: proto.ColumnType_STRING},
			{Name: "zip", Description: "", Type: proto.ColumnType_STRING},
		},
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

	returnedObject, err := client.NowContact.Read(sysId)
	if err != nil {
		logger.Error("servicenow_now_consumer.getServicenowNowConsumers", "query_error", err)
		return nil, err
	}

	return returnedObject.Result, err
}
