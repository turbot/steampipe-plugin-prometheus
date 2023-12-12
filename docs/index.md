---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/prometheus.svg"
brand_color: "#E6522C"
display_name: "Prometheus"
short_name: "prometheus"
description: "Steampipe plugin to query metrics, labels, alerts and more from Prometheus."
og_description: "Query Prometheus with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/prometheus-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Prometheus + Steampipe

[Prometheus](https://prometheus.io) is an open-source monitoring system with a dimensional data model, flexible query language, efficient time series database and modern alerting approach.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/prometheus/tables)**

## Get started

### Install

Download and install the latest Prometheus plugin:

```bash
steampipe plugin install prometheus
```

### Configuration

Installing the latest prometheus plugin will create a config file (`~/.steampipe/config/prometheus.spc`) with a single connection named `prometheus`:

```hcl
connection "prometheus" {
  plugin = "prometheus"


  # The address of your Prometheus.
  # Can also be set with the PROMETHEUS_URL environment variable
  address = "http://localhost:9090"
  metrics = ["prometheus_http_requests_.*", ".*error.*"]
}
```

- `address` - HTTP address of your prometheus server. Can also be set with the PROMETHEUS_URL environment variable.
- `metrics` - List of metric expressions to be matched against while creating dynamic metric tables.


