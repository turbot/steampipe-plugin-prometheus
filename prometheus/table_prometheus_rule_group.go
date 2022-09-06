package prometheus

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tablePrometheusRuleGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_rule_group",
		Description: "Rule groups in the Prometheus server.",
		List: &plugin.ListConfig{
			Hydrate: listRuleGroup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the group. Must be unique within a file."},
			{Name: "file", Type: proto.ColumnType_STRING, Description: "Path to the rule group file definition."},
			{Name: "interval", Type: proto.ColumnType_DOUBLE, Description: "How often rules in the group are evaluated in seconds."},
		},
	}
}

func listRuleGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_rule_group.listRule", "connection_error", err)
		return nil, err
	}
	items, err := conn.Rules(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_rule_group.listRule", "query_error", err)
		return nil, err
	}
	for _, i := range items.Groups {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
