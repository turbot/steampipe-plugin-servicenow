package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-servicenow"

// Plugin creates this (servicenow) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		// DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"servicenow_cmdb_ci":                     tableServicenowCmdbCi(),
			"servicenow_incident":                    tableServicenowIncident(),
			"servicenow_sys_user":                    tableServicenowSysUser(),
			"servicenow_sn_km_api_knowledge_article": tableServicenowSnKmApiKnowledgeArticle(),
		},
	}

	return p
}
