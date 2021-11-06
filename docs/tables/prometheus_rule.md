# Table: prometheus_rule

List rules in the Prometheus server.

## Examples

### List all rules

```sql
select
  *
from
  prometheus_rule
order by
  group_name,
  group_rule_num
```

### Rules with a labeled severity of high

```sql
select
  name,
  labels,
  state
from
  prometheus_rule
where
  labels ->> 'severity' = 'high'
```

### Rules that are firing

```sql
select
  name,
  labels,
  state
from
  prometheus_rule
where
  state = 'firing'
```

### Slow running rules with evaluation time > 1 sec

```sql
select
  name,
  labels,
  state
from
  prometheus_rule
where
  evaluation_time > 1
```
