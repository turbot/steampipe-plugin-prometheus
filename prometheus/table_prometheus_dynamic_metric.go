package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDynamicMetric(ctx context.Context, d *plugin.QueryData) *plugin.Table {

	conn, err := connectRaw(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_dynamic_metric.tableDynamicMetric", "connection_error", err)
		return nil
	}

	// Get the query for the metric (required)
	metricName := ctx.Value("metric_name").(string)
	metricNameBytes, _ := json.Marshal(metricName)
	q := fmt.Sprintf(`{__name__=%s}`, string(metricNameBytes))

	// Query parameters. Default to results from the current point in time only.
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute * time.Duration(5),
	}

	result, warnings, err := conn.QueryRange(ctx, q, r)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_dynamic_metric.tableDynamicMetric", "query_error", err)
		return nil
	}

	// Top columns
	cols := []*plugin.Column{
		{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Timestamp").Transform(transform.UnixMsToTimestamp), Description: "Timestamp of the value."},
		{Name: "value", Type: proto.ColumnType_DOUBLE, Description: "Value of the metric."},
	}

	// List of columns already seen
	seenCols := map[string]bool{"__name__": true, "labels": true, "step_seconds": true, "timestamp": true, "value": true}

	// Key columns for query filtering
	keyColumns := []*plugin.KeyColumn{
		{Name: "step_seconds", Require: plugin.Optional},
		{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
	}

	// If there are any results, check them for label data to build the columns
	for _, i := range result.(model.Matrix) {
		for k := range i.Metric {
			sk := string(k)
			if seenCols[sk] {
				continue
			}
			cols = append(cols, &plugin.Column{Name: sk, Type: proto.ColumnType_STRING, Transform: transform.FromField("Metric").TransformP(getMetricLabelFromMetric, sk), Description: fmt.Sprintf("The %s label.", sk)})
			keyColumns = append(keyColumns, &plugin.KeyColumn{Name: sk, Require: plugin.Optional})
		}
		// Feels silly, but a simple way to check the first result only...
		break
	}

	// Log the warnings
	for _, i := range warnings {
		plugin.Logger(ctx).Error("prometheus_dynamic_metric.tableDynamicMetric", "query_warning", i)
	}

	// If there were no labels, then return the full set in each row
	if len(cols) == 2 {
		cols = append(cols, &plugin.Column{Name: "labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Metric"), Description: "Map of all labels in the metric."})
	}

	// Other columns added at the end
	cols = append(cols, &plugin.Column{Name: "step_seconds", Type: proto.ColumnType_INT, Transform: transform.FromQual("step_seconds").Transform(getStepSeconds), Description: "Interval in seconds between metric values. Default 60 seconds."})

	return &plugin.Table{
		Name:        metricName,
		Description: fmt.Sprintf("Metrics for %s.", metricName),
		List: &plugin.ListConfig{
			KeyColumns: keyColumns,
			Hydrate:    listMetricWithName(metricName),
		},
		Columns: cols,
	}
}

func listMetricWithName(metricName string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("prometheus_query.listMetricWithName", "connection_error", err)
			return nil, err
		}

		// Query parameters. Default to results from the current point in time only.
		r := v1.Range{
			Start: time.Now(),
			End:   time.Now(),
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
				case "<=", "<":
					r.End = ts
				}
			}
			stepSeconds := (r.End.Sub(r.Start) / 1000).Round(time.Second)
			// Step has to be higher than 0 seconds
			if stepSeconds < 1 {
				stepSeconds = time.Second
			}
			r.Step = stepSeconds
		} else {
			isRange = false
		}

		// Allow user to change in the query
		if d.Quals["step_seconds"] != nil {
			r.Step = time.Second * time.Duration(d.KeyColumnQuals["step_seconds"].GetInt64Value())
		}

		// Always filter results to the specific metric
		metricNameBytes, _ := json.Marshal(metricName)
		pairs := []string{fmt.Sprintf(`__name__=%s`, string(metricNameBytes))}
		// Add any other qualifier filters
		for k, v := range d.KeyColumnQuals {
			// Skip any non-label quals
			if k == "timestamp" || k == "step_seconds" {
				continue
			}
			// A convenient way to encode the quotes etc
			valueBytes, _ := json.Marshal(v.GetStringValue())
			pairs = append(pairs, fmt.Sprintf(`%s=%s`, k, string(valueBytes)))
		}
		// Combine into a single PromQL query
		q := fmt.Sprintf("{%s}", strings.Join(pairs, ","))

		plugin.Logger(ctx).Trace("prometheus_dynamic_metric.listMetricWithName", "q", q)
		plugin.Logger(ctx).Trace("prometheus_dynamic_metric.listMetricWithName", "r", r)

		if isRange {

			// PRE: query is for data over time

			result, warnings, err := conn.QueryRange(ctx, q, r)
			if err != nil {
				plugin.Logger(ctx).Error("prometheus_query.listMetricWithName", "query_error", err)
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
				plugin.Logger(ctx).Error("prometheus_query.listMetricWithName", "query_warning", i)
			}

		} else {

			// PRE: query is for a single point in time.

			result, warnings, err := conn.Query(ctx, q, timestamp)
			if err != nil {
				plugin.Logger(ctx).Error("prometheus_query.listMetricWithName", "query_error", err)
				return nil, err
			}
			switch result := result.(type) {
			case model.Vector:
				{
					// Stream the results
					for _, i := range result {
						d.StreamListItem(ctx, i)
					}
				}
			case model.Matrix:
				{
					// Stream the results
					for _, i := range result {
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
				plugin.Logger(ctx).Error("prometheus_query.listMetricWithName", "query_warning", i)
			}

		}

		return nil, nil
	}
}

func getMetricLabelFromMetric(_ context.Context, d *transform.TransformData) (interface{}, error) {
	param := d.Param.(string)
	ls := d.Value.(model.Metric)
	return ls[model.LabelName(param)], nil
}
