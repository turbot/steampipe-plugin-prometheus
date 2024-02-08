package prometheus

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type ctxKey string

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-prometheus",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo(),
		SchemaMode:       plugin.SchemaModeDynamic,
		TableMapFunc:     pluginTableDefinitions,
	}
	return p
}

func pluginTableDefinitions(ctx context.Context, p *plugin.TableMapData) (map[string]*plugin.Table, error) {

	// Initialize tables
	tables := map[string]*plugin.Table{
		"prometheus_alert":      tablePrometheusAlert(ctx),
		"prometheus_label":      tablePrometheusLabel(ctx),
		"prometheus_metric":     tablePrometheusMetric(ctx),
		"prometheus_rule":       tablePrometheusRule(ctx),
		"prometheus_rule_group": tablePrometheusRuleGroup(ctx),
		"prometheus_series":     tablePrometheusSeries(ctx),
		"prometheus_target":     tablePrometheusTarget(ctx),
	}

	// Get list of metrics to create tables for from config
	prometheusConfig := GetConfig(p.Connection)
	  if prometheusConfig.Metrics == nil {
		return tables, nil
	}

	// Search for metrics to create as tables
	metricNames, err := metricNameList(ctx, p)

	if err != nil {
		// Return only the static tables when encountering an error
		return tables, err
	}

	for _, i := range metricNames {
		tableCtx := context.WithValue(ctx, ctxKey("metric_name"), i)
		base := filepath.Base(i)
		tableName := base[0 : len(base)-len(filepath.Ext(base))]
		// Add the table if it does not already exist, ensuring standard tables win
		if tables[tableName] == nil {
			tables[tableName] = tableDynamicMetric(tableCtx, p)
		} else {
			plugin.Logger(ctx).Error("prometheus.pluginTableDefinitions", "table_already_exists", tableName)
		}
	}
	
	return tables, nil
}

func metricNameList(ctx context.Context, p *plugin.TableMapData) ([]string, error) {
	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()

	// Get list of metrics to create tables for from config
	prometheusConfig := GetConfig(p.Connection)
	if prometheusConfig.Metrics == nil {
		return []string{}, nil
	}

	metrics := prometheusConfig.Metrics
	q := "{__name__=~\"" + strings.Join(metrics, "|") + "\"}"
	matches := []string{q}

	conn, err := connectRaw(ctx, p.ConnectionCache, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus.metricNameList", "connection_error", err)
		return nil, err
	}

	result, warnings, err := conn.LabelValues(ctx, "__name__", matches, startTime, endTime)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus.metricNameList", "query_error", err)
		return nil, err
	}

	names := []string{}
	for _, i := range result {
		names = append(names, string(i))
	}

	// Log the warnings
	for _, i := range warnings {
		plugin.Logger(ctx).Error("prometheus.metricNameList", "warning", i)
	}
	return names, nil
}
