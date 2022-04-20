package prometheus

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tablePrometheusAlert(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_alert",
		Description: "Alerts currently firing on the Prometheus server.",
		List: &plugin.ListConfig{
			Hydrate: listAlert,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "active_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the alert became active."},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations set on the alert rule."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels set on the metric."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the metric, e.g. firing."},
			{Name: "value", Type: proto.ColumnType_DOUBLE, Description: "Value of the metric."},
		},
	}
}

func listAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_alert.listAlert", "connection_error", err)
		return nil, err
	}
	items, err := conn.Alerts(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_alert.listAlert", "query_error", err)
		return nil, err
	}
	for _, i := range items.Alerts {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
