package servicenow

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-servicenow/rest"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*rest.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*rest.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	servicenowConfig := GetConfig(d.Connection)

	client, err := rest.New(rest.Config{
		InstanceURL:  *servicenowConfig.InstanceURL,
		GrantType:    "password",
		ClientID:     *servicenowConfig.ClientID,
		ClientSecret: *servicenowConfig.ClientSecret,
		Username:     *servicenowConfig.Username,
		Password:     *servicenowConfig.Password,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
