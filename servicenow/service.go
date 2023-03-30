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

	client, err := servicenow.New(servicenow.Config{
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

func ConnectUncached(ctx context.Context, conn *plugin.Connection) (*servicenow.ServiceNow, error) {
	servicenowConfig := GetConfig(conn)

	client, err := servicenow.New(servicenow.Config{
		InstanceURL:  *servicenowConfig.InstanceURL,
		GrantType:    "password",
		ClientID:     *servicenowConfig.ClientID,
		ClientSecret: *servicenowConfig.ClientSecret,
		Username:     *servicenowConfig.Username,
		Password:     *servicenowConfig.Password,
	})
	return client, err
}
