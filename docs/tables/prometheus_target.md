# Table: prometheus_target

List targets being scraped by the Prometheus server.

## Examples

### List all targets

```sql
select
  *
from
  prometheus_target
```

### Targets that are not up

```sql
select
  scrape_pool,
  scrape_url,
  health,
  last_scrape,
  last_error
from
  prometheus_target
where
  health != 'up'
```

### Targets whose last scrape was more than 24 hrs ago

```sql
select
  scrape_pool,
  scrape_url,
  health,
  last_scrape,
  last_error
from
  prometheus_target
where
  last_scrape < now() - interval '24 hrs'
```
