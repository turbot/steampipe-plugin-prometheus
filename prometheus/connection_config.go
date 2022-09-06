package prometheus

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type prometheusConfig struct {
	Address *string  `cty:"address"`
	Metrics []string `cty:"metrics"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
	"metrics": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &prometheusConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) prometheusConfig {
	if connection == nil || connection.Config == nil {
		return prometheusConfig{}
	}
	config, _ := connection.Config.(prometheusConfig)
	return config
}
