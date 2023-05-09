package servicenow

import (
	"context"

	"github.com/turbot/go-servicenow/servicenow"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*servicenow.ServiceNow, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*servicenow.ServiceNow), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	servicenowConfig := GetConfig(d.Connection)

	config := servicenow.Config{
		InstanceURL: *servicenowConfig.InstanceURL,
		BasicAuth:   isBasicAuth(servicenowConfig),
		GrantType:   "password",
	}

	if servicenowConfig.ClientID != nil {
		config.ClientID = *servicenowConfig.ClientID
	}
	if servicenowConfig.ClientSecret != nil {
		config.ClientSecret = *servicenowConfig.ClientSecret
	}
	if servicenowConfig.Username != nil {
		config.Username = *servicenowConfig.Username
	}
	if servicenowConfig.Password != nil {
		config.Password = *servicenowConfig.Password
	}

	client, err := servicenow.New(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConnectUncached(ctx context.Context, conn *plugin.Connection) (*servicenow.ServiceNow, error) {
	servicenowConfig := GetConfig(conn)

	config := servicenow.Config{
		InstanceURL: *servicenowConfig.InstanceURL,
		BasicAuth:   isBasicAuth(servicenowConfig),
		GrantType:   "password",
	}

	if servicenowConfig.ClientID != nil {
		config.ClientID = *servicenowConfig.ClientID
	}
	if servicenowConfig.ClientSecret != nil {
		config.ClientSecret = *servicenowConfig.ClientSecret
	}
	if servicenowConfig.Username != nil {
		config.Username = *servicenowConfig.Username
	}
	if servicenowConfig.Password != nil {
		config.Password = *servicenowConfig.Password
	}

	client, err := servicenow.New(config)
	if err != nil {
		return nil, err
	}

	return client, err
}

func isBasicAuth(config servicenowConfig) bool {
	if config.ClientID == nil && config.ClientSecret == nil {
		return true
	}
	return false
}
