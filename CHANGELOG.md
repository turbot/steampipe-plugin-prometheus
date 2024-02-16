## v0.6.2 [2024-02-13]

_Bug fixes_

- Fixed the plugin initialization error by returning only the static tables when invalid config parameters were set for dynamic tables. ([#39](https://github.com/turbot/steampipe-plugin-prometheus/pull/39))

## v0.6.1 [2023-12-12]

_Bug fixes_

- Fixed the missing optional tag on the `Metrics` connection config parameter. [#36](https://github.com/turbot/steampipe-plugin-prometheus/pull/36)

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#34](https://github.com/turbot/steampipe-plugin-prometheus/pull/34))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#34](https://github.com/turbot/steampipe-plugin-prometheus/pull/34))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-prometheus/blob/main/docs/LICENSE). ([#34](https://github.com/turbot/steampipe-plugin-prometheus/pull/34))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#33](https://github.com/turbot/steampipe-plugin-prometheus/pull/33))

## v0.5.0 [2023-10-20]

_What's new?_

- The Prometheus address (`address`) can now be set with the `PROMETHEUS_URL` environment variable. ([#23](https://github.com/turbot/steampipe-plugin-prometheus/pull/23)) (Thanks [@beudbeud](https://github.com/beudbeud) for the contribution!)

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#24](https://github.com/turbot/steampipe-plugin-prometheus/pull/24))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#19](https://github.com/turbot/steampipe-plugin-prometheus/pull/19))
- Recompiled plugin with Go version `1.21`. ([#19](https://github.com/turbot/steampipe-plugin-prometheus/pull/19))

## v0.3.0 [2023-06-20]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. This update significantly lowers the plugin initialization time of dynamic plugins by avoiding recursing into child folders when not necessary. ([#13](https://github.com/turbot/steampipe-plugin-prometheus/pull/13))

## v0.2.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#9](https://github.com/turbot/steampipe-plugin-prometheus/pull/9))
- Recompiled plugin with Go version `1.19`. ([#9](https://github.com/turbot/steampipe-plugin-prometheus/pull/9))

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
