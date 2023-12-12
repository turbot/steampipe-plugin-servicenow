package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type servicenowConfig struct {
	InstanceURL  *string   `hcl:"instance_url"`
	ClientID     *string   `hcl:"client_id"`
	ClientSecret *string   `hcl:"client_secret"`
	Username     *string   `hcl:"username"`
	Password     *string   `hcl:"password"`
	Objects      *[]string `hcl:"objects"`
}

func ConfigInstance() interface{} {
	return &servicenowConfig{}
}

func GetConfig(connection *plugin.Connection) servicenowConfig {
	if connection == nil || connection.Config == nil {
		return servicenowConfig{}
	}
	config, _ := connection.Config.(servicenowConfig)
	return config
}
