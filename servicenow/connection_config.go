package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
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
	if connection == nil || connection.GetConfig() == nil {
		return servicenowConfig{}
	}
	config, _ := connection.GetConfig().(servicenowConfig)
	return config
}
