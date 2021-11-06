# Table: prometheus_label

List labels and their values discovered by the Prometheus server.

## Examples

### List all labels names

```sql
select
  name
from
  prometheus_label
```

### List all labels with their values

```sql
select
  ln.name as name,
  value
from
  prometheus_label as ln,
  jsonb_array_elements_text(ln.values) as value
order by
  name,
  value
```
