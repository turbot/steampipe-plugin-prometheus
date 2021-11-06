package prometheus

import (
	"context"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePrometheusMetric(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_metric",
		Description: "Query metrics in the Prometheus server.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query"},
				{Name: "step_seconds", Require: plugin.Optional},
				{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
			},
			Hydrate: listMetric,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Metric").Transform(getMetricNameFromMetric), Description: "Name of the metric."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Timestamp").Transform(transform.UnixMsToTimestamp), Description: "Timestamp of the value."},
			{Name: "value", Type: proto.ColumnType_DOUBLE, Description: "Value of the metric."},
			// Other columns
			{Name: "labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Metric"), Description: "Labels for the metric."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query used to filter the metric data."},
			{Name: "step_seconds", Type: proto.ColumnType_INT, Transform: transform.FromQual("step_seconds").Transform(getStepSeconds), Description: "Interval in seconds between metric values. Default 60 seconds."},
		},
	}
}

func listMetric(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_query.listMetric", "connection_error", err)
		return nil, err
	}

	// Default to a 1 minute step interval. Only used if the timestamps create a range.
	stepSeconds := int64(60)
	// But allow it to be changed by the query
	if d.KeyColumnQuals["step_seconds"] != nil {
		stepSeconds = d.KeyColumnQuals["step_seconds"].GetInt64Value()
	}

	// Query parameters. Default to results from the current point in time only.
	r := v1.Range{
		Start: time.Now(),
		End:   time.Now(),
		Step:  time.Second * time.Duration(stepSeconds),
	}

	// Allow the query to set a range to get values over time
	timestamp := time.Now()
	isRange := true
	if d.Quals["timestamp"] != nil {
		for _, tq := range d.Quals["timestamp"].Quals {
			ts := tq.Value.GetTimestampValue().AsTime()
			switch tq.Operator {
			case ">", ">=":
				r.Start = ts
			case "=":
				isRange = false
				timestamp = ts
				break
			case "<=", "<":
				r.End = ts
			}
		}
	} else {
		isRange = false
	}

	// Get the query (required)
	q := d.KeyColumnQuals["query"].GetStringValue()

	if isRange {

		// PRE: query is for data over time

		result, warnings, err := conn.QueryRange(ctx, q, r)
		if err != nil {
			plugin.Logger(ctx).Error("prometheus_query.listMetric", "query_error", err)
			return nil, err
		}

		// Stream the results
		for _, i := range result.(model.Matrix) {
			for _, v := range i.Values {
				row := model.Sample{
					Metric:    i.Metric,
					Timestamp: v.Timestamp,
					Value:     v.Value,
				}
				d.StreamListItem(ctx, row)
			}
		}

		// Log the warnings
		for _, i := range warnings {
			plugin.Logger(ctx).Error("prometheus_query.listMetric", "query_warning", i)
		}

	} else {

		// PRE: query is for a single point in time.

		result, warnings, err := conn.Query(ctx, q, timestamp)
		if err != nil {
			plugin.Logger(ctx).Error("prometheus_query.listMetric", "query_error", err)
			return nil, err
		}
		switch result.(type) {
		case model.Vector:
			{
				// Stream the results
				for _, i := range result.(model.Vector) {
					d.StreamListItem(ctx, i)
				}
			}
		case model.Matrix:
			{
				// Stream the results
				for _, i := range result.(model.Matrix) {
					for _, v := range i.Values {
						row := model.Sample{
							Metric:    i.Metric,
							Timestamp: v.Timestamp,
							Value:     v.Value,
						}
						d.StreamListItem(ctx, row)
					}
				}
			}
		}

		// Log the warnings
		for _, i := range warnings {
			plugin.Logger(ctx).Error("prometheus_query.listMetric", "query_warning", i)
		}

	}

	return nil, nil
}

func getMetricNameFromMetric(_ context.Context, d *transform.TransformData) (interface{}, error) {
	ls := d.Value.(model.Metric)
	return ls["__name__"], nil
}

func getStepSeconds(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return int64(60), nil
	}
	return d.Value.(int64), nil
}
