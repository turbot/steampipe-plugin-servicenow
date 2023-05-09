package servicenow

import (
	"context"
	"os"

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

	// Overwrite config variables with any environments variables
	instanceUrl := os.Getenv("SERVICENOW_INSTANCE_URL")
	if instanceUrl != "" {
		servicenowConfig.InstanceURL = &instanceUrl
	}
	clientId := os.Getenv("SERVICENOW_CLIENT_ID")
	if clientId != "" {
		servicenowConfig.ClientID = &clientId
	}
	clientSecret := os.Getenv("SERVICENOW_CLIENT_SECRET")
	if clientSecret != "" {
		servicenowConfig.ClientSecret = &clientSecret
	}
	username := os.Getenv("SERVICENOW_USERNAME")
	if username != "" {
		servicenowConfig.Username = &username
	}
	password := os.Getenv("SERVICENOW_PASSWORD")
	if password != "" {
		servicenowConfig.Password = &password
	}

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

	// Overwrite config variables with any environments variables
	instanceUrl := os.Getenv("SERVICENOW_INSTANCE_URL")
	if instanceUrl != "" {
		servicenowConfig.InstanceURL = &instanceUrl
	}
	clientId := os.Getenv("SERVICENOW_CLIENT_ID")
	if clientId != "" {
		servicenowConfig.ClientID = &clientId
	}
	clientSecret := os.Getenv("SERVICENOW_CLIENT_SECRET")
	if clientSecret != "" {
		servicenowConfig.ClientSecret = &clientSecret
	}
	username := os.Getenv("SERVICENOW_USERNAME")
	if username != "" {
		servicenowConfig.Username = &username
	}
	password := os.Getenv("SERVICENOW_PASSWORD")
	if password != "" {
		servicenowConfig.Password = &password
	}

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
