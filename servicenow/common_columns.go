package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "instance_url",
			Description: "The ServiceNow instance URL.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getInstanceUrl,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getInstanceUrlMemoized = plugin.HydrateFunc(getInstanceUrlUncached).Memoize(memoize.WithCacheKeyFunction(getInstanceUrlCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getInstanceUrl(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getInstanceUrlMemoized(ctx, d, h)
}

// Build a cache key for the call to getInstanceUrlCacheKey.
func getInstanceUrlCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getInstanceUrl"
	return key, nil
}

func getInstanceUrlUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	servicenowConfig := GetConfig(d.Connection)

	return servicenowConfig.InstanceURL, nil
}
