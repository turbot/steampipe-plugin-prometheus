package prometheus

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePrometheusLabel(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_label",
		Description: "Labels used in metrics in the Prometheus server.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", Require: plugin.Optional},
				{Name: "timestamp", Operators: []string{">", ">=", "=", "<=", "<"}, Require: plugin.Optional},
			},
			Hydrate: listLabel,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the labels were found."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the label."},
			{Name: "values", Type: proto.ColumnType_JSON, Hydrate: listLabelValue, Transform: transform.FromValue(), Description: "Values for the label."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query used to filter the label search."},
		},
	}
}

type labelNameRow struct {
	Timestamp time.Time
	Name      string
}

func listLabel(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_label.listLabel", "connection_error", err)
		return nil, err
	}

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

	ts := startTime.Add(endTime.Sub(startTime) / 2)

	q := []string{}
	if d.KeyColumnQuals["query"] != nil {
		qs := d.KeyColumnQuals["query"].GetStringValue()
		if qs != "" {
			q = append(q, qs)
		}
	}

	result, warnings, err := conn.LabelNames(ctx, q, startTime, endTime)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_label.listLabel", "label_error", err)
		return nil, err
	}

	for _, i := range result {
		r := labelNameRow{
			Timestamp: ts,
			Name:      i,
		}
		d.StreamListItem(ctx, r)
	}

	// Log the warnings
	for _, i := range warnings {
		plugin.Logger(ctx).Error("prometheus_label.listLabel", "label_warning", i)
	}

	return nil, nil
}

func listLabelValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_label_value.listLabelValue", "connection_error", err)
		return nil, err
	}

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

	q := []string{}
	if d.KeyColumnQuals["query"] != nil {
		qs := d.KeyColumnQuals["query"].GetStringValue()
		if qs != "" {
			q = append(q, qs)
		}
	}

	result, warnings, err := conn.LabelValues(ctx, h.Item.(labelNameRow).Name, q, startTime, endTime)
	if err != nil {
		plugin.Logger(ctx).Error("prometheus_label_value.listLabelValue", "labelValue_error", err)
		return nil, err
	}

	// Log the warnings
	for _, i := range warnings {
		plugin.Logger(ctx).Error("prometheus_label_value.listLabelValue", "labelValue_warning", i)
	}

	return result, nil
}
