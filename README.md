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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-prometheus/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Prometheus Plugin](https://github.com/turbot/steampipe-plugin-prometheus/labels/help%20wanted)
