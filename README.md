![image](https://hub.steampipe.io/images/plugins/turbot/prometheus-social-graphic.png)

# Prometheus Plugin for Steampipe

Use SQL to query instances, domains and more from Prometheus.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/prometheus)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/prometheus/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-prometheus/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install prometheus
```

Configure the server address in `~/.steampipe/config/prometheus.spc`:

```hcl
connection "prometheus" {
  plugin  = "prometheus"
  address = "http://localhost:9090"
  username = "optional_username"
  password = "optional_password"
}
```

Run steampipe:

```shell
steampipe query
```

Query all the labels in your prometheus metrics:

```sql
select
  name,
  values
from
  prometheus_label
```

```
> select name, values from prometheus_label
+---------------+----------------------------------------------+
| name          | values                                       |
+---------------+----------------------------------------------+
| alertname     | ["TotalRequests"]                            |
| alertstate    | ["firing"]                                   |
| reason        | ["refused","resolution","timeout","unknown"] |
| interval      | ["10s"]                                      |
| version       | ["2.30.3","go1.17.1"]                        |
| code          | ["200","302","400","500","503"]              |
+---------------+----------------------------------------------+
```

Query data for a given metric (tables are dynamically created):

```sql
select
  code,
  handler,
  value
from
  prometheus_http_requests_total
```

```
+------+----------------------------+-------+
| code | handler                    | value |
+------+----------------------------+-------+
| 302  | /                          | 1     |
| 200  | /-/ready                   | 1     |
| 200  | /api/v1/alerts             | 1     |
| 200  | /api/v1/label/:name/values | 421   |
| 200  | /api/v1/labels             | 16    |
| 200  | /graph                     | 1     |
| 200  | /static/*filepath          | 4     |
+------+----------------------------+-------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-prometheus.git
cd steampipe-plugin-prometheus
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/prometheus.spc
```

Try it!

```
steampipe query
> .inspect prometheus
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Prometheus Plugin](https://github.com/turbot/steampipe-plugin-prometheus/labels/help%20wanted)
