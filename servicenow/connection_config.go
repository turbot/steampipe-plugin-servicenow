package servicenow

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type servicenowConfig struct {
	InstanceURL  *string   `cty:"instance_url"`
	ClientID     *string   `cty:"client_id"`
	ClientSecret *string   `cty:"client_secret"`
	Username     *string   `cty:"username"`
	Password     *string   `cty:"password"`
	Objects      *[]string `cty:"objects"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"instance_url": {
		Type: schema.TypeString,
	},
	"client_id": {
		Type: schema.TypeString,
	},
	"client_secret": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"objects": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{
			Type: schema.TypeString,
		},
	},
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
