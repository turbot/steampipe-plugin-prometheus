---
title: "Steampipe Table: prometheus_series - Query Prometheus Time Series Data using SQL"
description: "Allows users to query Time Series Data in Prometheus, specifically the series of data points indexed by time, providing insights into system metrics over a period of time."
---

# Table: prometheus_series - Query Prometheus Time Series Data using SQL

Prometheus is an open-source system monitoring and alerting toolkit that collects and stores its metrics as time series data, i.e., metrics information that is output at a regular interval. This includes various system metrics such as CPU usage, memory utilization, disk I/O, network traffic, etc. It provides a multidimensional data model with time series data identified by metric name and key-value pairs.

## Table Usage Guide

The `prometheus_series` table provides insights into time series data within Prometheus. As a System Administrator or DevOps engineer, you can explore specific metrics details through this table, including the metric name, labels, and timestamp. Utilize it to uncover information about system performance over time, identify trends, and aid in capacity planning.

## Examples

### List all current prometheus_http_requests_total series
Explore the current series of Prometheus HTTP requests to understand the volume and nature of web traffic. This can be useful in monitoring server performance and identifying potential issues or areas for optimization.

```sql+postgres
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total';
```

```sql+sqlite
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total';
```

### List all prometheus_http_requests_total series present 24 hours ago
Explore the total number of HTTP requests made to your Prometheus server 24 hours ago. This can help in identifying any unusual spikes or drops in traffic, assisting in network monitoring and troubleshooting.

```sql+postgres
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
  and timestamp = now() - interval '24 hrs';
```

```sql+sqlite
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
  and timestamp = datetime('now', '-24 hours');
```

### List all prometheus_http_requests_total series for /metrics present 24 hours ago
Analyze the settings to understand the total number of HTTP requests made to the '/metrics' handler in Prometheus, exactly 24 hours ago. This can help in monitoring traffic patterns and identifying possible issues or anomalies in request volume.

```sql+postgres
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total{handler="/metrics"}'
  and timestamp = now() - interval '24 hrs';
```

```sql+sqlite
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total{handler="/metrics"}'
  and timestamp = datetime('now','-24 hours');
```

### List all prometheus_http_requests_total series on 31st Oct 2021
Explore the total number of HTTP requests recorded by Prometheus on October 31st, 2021. This can help in analyzing the web traffic patterns and server load on that specific day.

```sql+postgres
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
  and timestamp > '2021-10-31'
  and timestamp < '2021-11-01';
```

```sql+sqlite
select
  *
from
  prometheus_series
where
  query = 'prometheus_http_requests_total'
  and timestamp > '2021-10-31'
  and timestamp < '2021-11-01';
```