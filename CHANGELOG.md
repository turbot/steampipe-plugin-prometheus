
## v0.1.0 [2022-05-25]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#5](https://github.com/turbot/steampipe-plugin-prometheus/pull/5))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#4](https://github.com/turbot/steampipe-plugin-prometheus/pull/4))

## v0.0.1 [2022-01-11]

_What's new?_

- New tables added
  - [prometheus_alert](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_alert)
  - [prometheus_label](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_label)
  - [prometheus_metric](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_metric)
  - [prometheus_rule](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_rule)
  - [prometheus_rule_group](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_rule_group)
  - [prometheus_series](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_series)
  - [prometheus_target](https://hub.steampipe.io/plugins/turbot/prometheus/tables/prometheus_target)
  - [{metric_name}](https://hub.steampipe.io/plugins/turbot/prometheus/tables/{metric_name})
