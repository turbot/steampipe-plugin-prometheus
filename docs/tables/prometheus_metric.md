---
title: "Steampipe Table: prometheus_metric - Query Prometheus Metrics using SQL"
description: "Allows users to query Metrics in Prometheus, specifically the numerical data about the system's state, providing insights into system performance and potential anomalies."
---

# Table: prometheus_metric - Query Prometheus Metrics using SQL

Prometheus is an open-source systems monitoring and alerting toolkit. It collects numerical data about the state of a system at any point in time. This data is stored as a series of metrics, which can be queried and visualized to gain insights into system performance.

## Table Usage Guide

The `prometheus_metric` table provides insights into the numerical data about the state of a system in Prometheus. As a system administrator or DevOps engineer, explore metric-specific details through this table, including metric names, labels, and values. Utilize it to monitor system performance, identify potential issues, and make data-driven decisions about system improvements.

**Important Notes**
- A `query` must be provided in all queries to this table.

## Examples

### Get current values for a metric
Explore the current values for a specific metric to monitor the performance and health of your system. This could be particularly useful in identifying potential issues or bottlenecks in your system's operation.

```sql+postgres
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total';
```

```sql+sqlite
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total';
```

### Get current values for a metric with specific labels
Explore the current values of a specific metric by identifying its unique labels. This can be beneficial in monitoring and analyzing the performance of your system based on certain parameters.

```sql+postgres
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total{handler="/metrics"}';
```

```sql+sqlite
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total{handler="/metrics"}';
```

### Get values from 24 hrs ago for a metric
Analyze the metrics to understand the changes in HTTP requests over the past 24 hours. This is particularly useful for monitoring server performance and identifying potential issues or anomalies.

```sql+postgres
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
  and timestamp = now() - interval '24 hrs';
```

```sql+sqlite
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
  and timestamp = datetime('now', '-24 hours');
```

### Get metric values every 5 mins for the last hour
Analyze the frequency of HTTP requests in the last hour, by obtaining metrics at 5-minute intervals. This can help monitor web traffic patterns and identify potential surges or dips in usage.

```sql+postgres
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
  and timestamp > now() - interval '1 hrs'
  and step_seconds = 300
order by
  timestamp;
```

```sql+sqlite
select
  *
from
  prometheus_metric
where
  query = 'prometheus_http_requests_total'
  and timestamp > datetime('now', '-1 hours')
  and step_seconds = 300
order by
  timestamp;
```