package prometheus

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-prometheus",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		TableMapFunc:     pluginTableDefinitions,
	}
	return p
}

func pluginTableDefinitions(ctx context.Context, p *plugin.Plugin) (map[string]*plugin.Table, error) {

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

	// Search for CSV files to create as tables
	metricNames, err := metricNameList(ctx, p)
	if err != nil {
		return nil, err
	}
	for _, i := range metricNames {
		tableCtx := context.WithValue(ctx, "metric_name", i)
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

func metricNameList(ctx context.Context, p *plugin.Plugin) ([]string, error) {

	conn, err := connectRaw(ctx, p.ConnectionManager, p.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus.metricNameList", "connection_error", err)
		return nil, err
	}

	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	q := []string{}

	// Access config to get the list of metrics to be queried
	prometheusConfig := GetConfig(p.Connection)
	if &prometheusConfig != nil {
		if prometheusConfig.Metrics != nil {
			metrics := prometheusConfig.Metrics
			for _, metric := range metrics {
				if metric == "" {
					plugin.Logger(ctx).Error("prometheus.metricNameList", "config_error", "A metric name must have at least one non-empty")
					return nil, errors.New("A metric name must have at least one non-empty matcher")
				}
				q = append(q, fmt.Sprintf("{__name__=~\"%s\"}", metric))
			}
		}
	}

	result, warnings, err := conn.LabelValues(ctx, "__name__", q, startTime, endTime)
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
