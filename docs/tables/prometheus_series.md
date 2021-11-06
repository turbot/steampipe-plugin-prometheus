# Table: prometheus_series

List series being scraped by the Prometheus server.

## Examples

### List all current prometheus_http_requests_total series

```sql
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
```

### List all prometheus_http_requests_total series present 24 hours ago

```sql
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
  and timestamp = now() - interval '24 hrs'
```

### List all prometheus_http_requests_total series for /metrics present 24 hours ago

```sql
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total{handler="/metrics"}'
  and timestamp = now() - interval '24 hrs'
```

### List all prometheus_http_requests_total series on 31st Oct 2021

```sql
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
  and timestamp > '2021-10-31'
  and timestamp < '2021-11-01'
```
