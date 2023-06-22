package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-servicenow/servicenow"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: servicenow.Plugin})
}
