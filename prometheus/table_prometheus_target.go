package prometheus

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePrometheusTarget(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_target",
		Description: "Target scraped by the Prometheus server.",
		List: &plugin.ListConfig{
			Hydrate: listTarget,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "scrape_pool", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target.ScrapePool"), Description: "Name of the scrape pool this target belongs to."},
			{Name: "scrape_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target.ScrapeURL"), Description: "URL to be scraped."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the target, e.g. active, dropped."},
			{Name: "health", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target.Health"), Description: "Health of the target, e.g. up."},
			{Name: "last_scrape", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Target.LastScrape"), Description: "Time when the last scrape occurred."},
			{Name: "labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Target.Labels"), Description: "Label set after relabelling has occurred."},
			// Other columns
			{Name: "discovered_labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Target.DiscoveredLabels"), Description: "Unmodified labels retrieved during service discovery before relabelling has occurred."},
			{Name: "global_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target.GlobalURL"), Description: "Global URL to be scraped."},
			{Name: "last_error", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target.LastError"), Description: "Last error message, if any."},
			{Name: "last_scrape_duration", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Target.LastScrapeDuration"), Description: "Time in seconds the last scrape took to run."},
		},
	}
}

type targetRow struct {
	State  string
	Target interface{}
}

func listTarget(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_target.listTarget", "connection_error", err)
		return nil, err
	}
	items, err := conn.Targets(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_target.listTarget", "query_error", err)
		return nil, err
	}
	for _, i := range items.Active {
		tr := targetRow{"active", i}
		d.StreamListItem(ctx, tr)
	}
	for _, i := range items.Dropped {
		tr := targetRow{"dropped", i}
		d.StreamListItem(ctx, tr)
	}
	return nil, nil
}
