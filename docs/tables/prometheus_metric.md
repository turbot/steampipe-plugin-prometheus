# Table: prometheus_metric

List metric values for a query and time range.

Notes:

* A `query` must be provided in all queries to this table.

## Examples

### Get current values for a metric

```sql
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
```

### Get current values for a metric with specific labels

```sql
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total{handler="/metrics"}'
```

### Get values from 24 hrs ago for a metric

```sql
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
  and timestamp = now() - interval '24 hrs'
```

### Get metric values every 5 mins for the last hour

```sql
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
  and timestamp > now() - interval '1 hrs'
  and step_seconds = 300
order by
  timestamp
```
