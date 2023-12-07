---
title: "Steampipe Table: {metric_name} - Query Prometheus metrics using SQL"
description: "Allows users to query Proemtheus metrics, providing insights into specific metrics and potential anomalies."
---

# Table: {metric_name} - Query Prometheus metrics using SQL

Query data from the metric called `{metric_name}`. A table is automatically created to represent each metric.

For instance, given the metric:

```
{
  "__name__": "prometheus_http_requests_total",
  "code": "302",
  "handler": "/",
  "instance": "localhost:9090",
  "job":"prometheus"
}
```

And the connection configuration:
```hcl
connection "prometheus" {
  plugin = "prometheus"
  address = "http://localhost:9090"
  metrics = ["prometheus_http_requests_total"]
}
```

This plugin will automatically create a table called `prometheus_http_requests_total`:
```
> select * from prometheus_http_requests_total;
+----------------------+-------+------+----------------------------+----------------+------------+--------------+
| timestamp            | value | code | handler                    | instance       | job        | step_seconds |
+----------------------+-------+------+----------------------------+----------------+------------+--------------+
| 2021-11-06T02:05:41Z | 308   | 200  | /api/v1/label/:name/values | localhost:9090 | prometheus | 60           |
| 2021-11-06T01:43:41Z | 1     | 200  | /-/ready                   | localhost:9090 | prometheus | 60           |
| 2021-11-06T01:56:41Z | 12    | 200  | /api/v1/labels             | localhost:9090 | prometheus | 60           |
+----------------------+-------+------+----------------------------+----------------+------------+--------------+
```

Regular expressions can also be used to match metric names. For instance, if
you want to create tables for all metrics starting with
`prometheus_target_`, use the following configuration:

```hcl
connection "prometheus" {
  plugin = "prometheus"
  address = "http://localhost:9090"
  metrics = ["prometheus_target_.*"]
}
```

If you want to create tables for all metrics, use:

```hcl
connection "prometheus" {
  plugin = "prometheus"
  address = "http://localhost:9090"
  metrics = [".+"]
}
```

However, please note that this could be slow depending on how many metrics are
in your environment.

## Examples

### Inspect the table structure

Assuming your connection configuration is:
```hcl
connection "prometheus" {
  plugin = "prometheus"
  address = "http://localhost:9090"
  metrics = ["prometheus_http_requests_total"]
}
```

List all tables with:

```sql
.inspect prometheus
+--------------------------------+---------------------------------------------------+
| table                          | description                                       |
+--------------------------------+---------------------------------------------------+
| prometheus_alert               | Alerts currently firing on the Prometheus server. |
| ...                            | ...                                               |
| prometheus_http_requests_total | Metrics for prometheus_http_requests_total.       |
| ...                            | ...                                               |
+--------------------------------+---------------------------------------------------+
```

To get details of a specific metric table, inspect it by name:

```sql
> .inspect prometheus.prometheus_http_requests_total
+--------------+--------------------------+----------------------------------------------------------------+
| column       | type                     | description                                                    |
+--------------+--------------------------+----------------------------------------------------------------+
| code         | text                     | The code label.                                                |
| handler      | text                     | The handler label.                                             |
| instance     | text                     | The instance label.                                            |
| job          | text                     | The job label.                                                 |
| step_seconds | bigint                   | Interval in seconds between metric values. Default 60 seconds. |
| timestamp    | timestamp with time zone | Timestamp of the value.                                        |
| value        | double precision         | Value of the metric.                                           |
+--------------+--------------------------+----------------------------------------------------------------+
```

### Get current values for prometheus_http_requests_total

```sql+postgres
select
  *
from
  prometheus_http_requests_total;
```

```sql+sqlite
select
  *
from
  prometheus_http_requests_total;
```

### Get current values for a metric with specific labels

```sql+postgres
select
  *
from
  prometheus_http_requests_total
where
  handler = '/metrics';
```

```sql+sqlite
select
  *
from
  prometheus_http_requests_total
where
  handler = '/metrics';
```

### Get values from 24 hrs ago for a metric

```sql+postgres
select
  timestamp,
  code,
  handler,
  value
from
  prometheus_http_requests_total
where
  timestamp = now() - interval '24 hrs';
```

```sql+sqlite
select
  timestamp,
  code,
  handler,
  value
from
  prometheus_http_requests_total
where
  timestamp = datetime('now','-24 hours');
```

### Get metric values every 5 mins for the last hour

```sql+postgres
select
  timestamp,
  code,
  handler,
  value
from
  prometheus_http_requests_total
where
  timestamp > now() - interval '1 hrs'
  and step_seconds = 300
order by
  timestamp;
```

```sql+sqlite
select
  timestamp,
  code,
  handler,
  value
from
  prometheus_http_requests_total
where
  timestamp > datetime('now','-1 hours')
  and step_seconds = 300
order by
  timestamp;
```