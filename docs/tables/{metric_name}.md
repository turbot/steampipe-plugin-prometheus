# Table: {metric_name}

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

## Examples

### Inspect the table structure

Assuming your connection is called `prometheus` (the default), list all tables with:
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

```sql
select
  *
from
  prometheus_http_requests_total
```

### Get current values for a metric with specific labels

```sql
select
  *
from
  prometheus_http_requests_total
where
  handler = '/metrics'
```

### Get values from 24 hrs ago for a metric

```sql
select
  timestamp,
  code,
  handler,
  value
from
  prometheus_http_requests_total
where
  timestamp = now() - interval '24 hrs'
```

### Get metric values every 5 mins for the last hour

```sql
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
  timestamp
```
