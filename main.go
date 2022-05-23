package main

import (
	"github.com/turbot/steampipe-plugin-prometheus/prometheus"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: prometheus.Plugin})
}
