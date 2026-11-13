package prometheus

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type prometheusConfig struct {
	Address *string  `hcl:"address"`
	Metrics []string `hcl:"metrics,optional"`
}

func ConfigInstance() interface{} {
	return &prometheusConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) prometheusConfig {
	if connection == nil || connection.GetConfig() == nil {
		return prometheusConfig{}
	}
	config, _ := connection.GetConfig().(prometheusConfig)
	return config
}
