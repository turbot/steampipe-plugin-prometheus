# Table: prometheus_alert

List alerts in the Prometheus server.

## Examples

### List all alerts

```sql
select
  *
from
  prometheus_alert
```

### Alerts with a labeled severity of high

```sql
select
  *
from
  prometheus_alert
where
  labels ->> 'severity' = 'high'
```

### Alerts that became active in the last 5 mins

```sql
select
  *
from
  prometheus_alert
where
  active_at > now() - interval '5 min'
```
