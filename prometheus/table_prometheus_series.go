package prometheus

import (
	"context"
	"time"

	"github.com/prometheus/common/model"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrometheusSeries(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_series",
		Description: "Series in the Prometheus server.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", Operators: []string{"="}},
				{Name: "timestamp", Operators: []string{">", ">=", "=", "<=", "<"}, Require: plugin.Optional},
			},
			Hydrate: listSeries,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the series was found."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Metric").Transform(getMetricNameFromLabelSet), Description: "Name of the metric for the series."},
			{Name: "metric", Type: proto.ColumnType_JSON, Description: "Metric details for the series."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query used to match the series."},
		},
	}
}

type seriesRow struct {
	Timestamp time.Time
	Metric    model.LabelSet
}

func listSeries(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_target.listSeries", "connection_error", err)
		return nil, err
	}

	q := d.EqualsQuals["query"].GetStringValue()

	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	if d.Quals["timestamp"] != nil {
		for _, q := range d.Quals["timestamp"].Quals {
			ts := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case ">", ">=":
				startTime = ts
			case "=":
				startTime = ts
				endTime = ts
			case "<=", "<":
				endTime = ts
			}
		}
	}

	// We discover series from the timerange quals passed in. But, series do not
	// have an inherent timestamp associated with them, so we choose the halfway
	// point as the column return value.
	ts := startTime.Add(endTime.Sub(startTime) / 2)

	result, warnings, err := conn.Series(ctx, []string{q}, startTime, endTime)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_target.listSeries", "series_error", err)
		return nil, err
	}

	for _, i := range result {
		r := seriesRow{
			Timestamp: ts,
			Metric:    i,
		}
		d.StreamListItem(ctx, r)
	}

	// Log the warnings
	for _, i := range warnings {
		plugin.Logger(ctx).Error("prometheus_target.listSeries", "series_warning", i)
	}

	return nil, nil
}

func getMetricNameFromLabelSet(_ context.Context, d *transform.TransformData) (interface{}, error) {
	ls := d.Value.(model.LabelSet)
	return ls["__name__"], nil
}
