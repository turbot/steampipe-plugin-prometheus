---
title: "Steampipe Table: prometheus_alert - Query Prometheus Alerts using SQL"
description: "Allows users to query Alerts in Prometheus, specifically the active alerts and their details, providing insights into system performance and potential issues."
---

# Table: prometheus_alert - Query Prometheus Alerts using SQL

Prometheus is an open-source systems monitoring and alerting toolkit. It provides a multi-dimensional data model with time series data identified by metric name and key/value pairs. Alerts in Prometheus sends notifications to external systems when certain conditions are observed in the monitored system.

## Table Usage Guide

The `prometheus_alert` table provides insights into active alerts within Prometheus. As a system administrator or DevOps engineer, explore alert-specific details through this table, including alert name, severity, and associated metadata. Utilize it to uncover information about active alerts, such as those with high severity, the conditions that triggered the alerts, and the duration for which the alerts have been active.

## Examples

### List all alerts
Gain insights into all the alerts currently active in your system. This can be useful for monitoring system health and responding quickly to any issues.

```sql+postgres
select
  *
from
  prometheus_alert;
```

```sql+sqlite
select
  *
from
  prometheus_alert;
```

### Alerts with a labeled severity of high
Identify instances where alerts have been marked with a high severity level. This can be useful in prioritizing responses to potential issues within your system.

```sql+postgres
select
  *
from
  prometheus_alert
where
  labels ->> 'severity' = 'high';
```

```sql+sqlite
select
  *
from
  prometheus_alert
where
  json_extract(labels, '$.severity') = 'high';
```

### Alerts that became active in the last 5 mins
Discover the alerts that have been activated recently, allowing you to respond promptly to any potential issues or threats. This can be particularly beneficial in real-time monitoring and incident response scenarios.

```sql+postgres
select
  *
from
  prometheus_alert
where
  active_at > now() - interval '5 min';
```

```sql+sqlite
select
  *
from
  prometheus_alert
where
  active_at > datetime('now', '-5 minutes');
```