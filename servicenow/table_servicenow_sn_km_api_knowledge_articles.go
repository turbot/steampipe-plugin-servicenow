package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableServicenowSnKmApiKnowledgeArticle() *plugin.Table {
	return &plugin.Table{
		Name:        "servicenow_sn_km_api_knowledge_article",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listServicenowSnKmApiKnowledgeArticles,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServicenowSnKmApiKnowledgeArticles,
			KeyColumns: plugin.SingleColumn("sys_id"),
		},
		Columns: []*plugin.Column{
			{Name: "sys_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SysID")},
			{Name: "id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "number", Description: "", Type: proto.ColumnType_STRING},
			{Name: "title", Description: "", Type: proto.ColumnType_STRING},
			{Name: "short_description", Description: "", Type: proto.ColumnType_STRING},
			{Name: "rank", Description: "", Type: proto.ColumnType_INT},
			{Name: "score", Description: "", Type: proto.ColumnType_DOUBLE},
			{Name: "snippet", Description: "", Type: proto.ColumnType_STRING},
			{Name: "link", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "fields", Description: "", Type: proto.ColumnType_JSON},
			{Name: "content", Description: "", Type: proto.ColumnType_STRING},
			{Name: "template", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "display_attachments", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "embedded_content", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listServicenowSnKmApiKnowledgeArticles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_km_api_knowledge_article.listServicenowSnKmApiKnowledgeArticles", "connect_error", err)
		return nil, err
	}

	articlesResponse, err := client.GetArticles(10)
	if err != nil {
		logger.Error("servicenow_sn_km_api_knowledge_article.listServicenowSnKmApiKnowledgeArticles", "query_error", err)
		return nil, err
	}
	for _, article := range articlesResponse.Result.Articles {
		d.StreamListItem(ctx, article)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, err
}

//// GET FUNCTION

func getServicenowSnKmApiKnowledgeArticles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("servicenow_sn_km_api_knowledge_article.getServicenowSnKmApiKnowledgeArticles", "connect_error", err)
		return nil, err
	}

	sysId := d.EqualsQualString("sys_id")

	article, err := client.GetArticle(sysId)
	if err != nil {
		logger.Error("servicenow_sn_km_api_knowledge_article.getServicenowSnKmApiKnowledgeArticles", "query_error", err)
		return nil, err
	}

	return article.Result, err
}
