package prometheus

import (
	"context"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrometheusRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prometheus_rule",
		Description: "Rules in the Prometheus server.",
		List: &plugin.ListConfig{
			ParentHydrate: listRuleGroup,
			Hydrate:       listRule,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rule.Name"), Description: "Name of the alert."},
			{Name: "last_evaluation", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Rule.LastEvaluation"), Description: "Time when the rule was last evaluated."},
			{Name: "health", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rule.Health"), Description: "Health of the rule, e.g. ok."},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rule.State"), Description: "State of the alert for this rule, e.g. firing, inactive."},
			// Other columns
			{Name: "alerts", Type: proto.ColumnType_JSON, Transform: transform.FromField("Rule.Alerts"), Description: "Alerts for the rule."},
			{Name: "annotations", Type: proto.ColumnType_JSON, Transform: transform.FromField("Rule.Annotations"), Description: "Annotations to add to each alert."},
			{Name: "duration", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Rule.Duration"), Description: "Alerts are considered firing once they have been returned for this long. Alerts which have not yet fired for long enough are considered pending."},
			{Name: "evaluation_time", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Rule.EvaluationTime"), Description: "Time taken in seconds to run the rule evaluation."},
			{Name: "group_file", Type: proto.ColumnType_STRING, Description: "Path to the file that defines the rule group for this rule."},
			{Name: "group_interval", Type: proto.ColumnType_STRING, Description: "Interval between evaluations for rules in this rule group."},
			{Name: "group_name", Type: proto.ColumnType_STRING, Description: "Name of the rule group that contains the rule."},
			{Name: "group_rule_num", Type: proto.ColumnType_INT, Description: "Rule number within the group. Starts at 1."},
			{Name: "labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Rule.Labels"), Description: "Labels to add or overwrite for each alert."},
			{Name: "last_error", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rule.LastError"), Description: "Last error message from the rule, if any."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rule.Query"), Description: "The PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts."},
		},
	}
}

type ruleRow struct {
	GroupName     string
	GroupFile     string
	GroupInterval float64
	GroupRuleNum  int
	Rule          interface{}
}

func listRule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	ruleGroup := h.Item.(v1.RuleGroup)
	for idx, i := range ruleGroup.Rules {
		r := ruleRow{
			GroupName:     ruleGroup.Name,
			GroupFile:     ruleGroup.File,
			GroupInterval: ruleGroup.Interval,
			Rule:          i,
			GroupRuleNum:  idx + 1,
		}
		d.StreamListItem(ctx, r)
	}
	return nil, nil
}
