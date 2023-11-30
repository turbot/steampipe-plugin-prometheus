---
title: "Steampipe Table: metric_name - Query OCI Service OCI_resource using SQL"
description: "Allows users to query OCI_resource in OCI Service, providing insights into specific metrics and potential anomalies."
---

# Table: metric_name - Query OCI Service OCI_resource using SQL

OCI Service is a service within Oracle Cloud Infrastructure that allows you to monitor and respond to issues across your applications and infrastructure. It provides a centralized way to set up and manage alerts for various OCI resources. OCI Service helps you stay informed about the health and performance of your OCI resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `metric_name` table provides insights into metrics within OCI Service. As a cloud engineer, explore metric-specific details through this table, including performance, utilization, and associated metadata. Utilize it to uncover information about metrics, such as those indicating performance issues, the relationships between different metrics, and the verification of utilization rates. For more details, refer to the [schema link](https://hub.steampipe.io/plugins/turbot/prometheus/tables/metric_name).

## Examples

### Inspect the table structure
Explore the structure of a specific metric table to gain insights into its composition and the type of data it holds. This is useful for understanding what information is available for analysis and how it is organized within the table.
Assuming your connection configuration is:
```hcl
connection "prometheus" {
  plugin = "prometheus"
  address = "http://localhost:9090"
  metrics = ["prometheus_http_requests_total"]
}
```

List all tables with:

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
Explore the current metrics of total HTTP requests in your Prometheus monitoring system. This can assist in understanding the load on your system and aid in capacity planning.

```sql
select
  *
from
  prometheus_http_requests_total
```

### Get current values for a metric with specific labels
Discover the segments that are currently active for a specific metric label, which can help in assessing the performance and identifying potential issues. This can be particularly useful in monitoring and managing web traffic for specific metrics.

```sql
select
  *
from
  prometheus_http_requests_total
where
  handler = '/metrics'
```

### Get values from 24 hrs ago for a metric
Explore the performance of your web server by examining metrics from 24 hours ago. This can help identify any significant changes or issues that have occurred within the past day.

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
Analyze the frequency of HTTP requests over the past hour, segmented into 5-minute intervals. This can help in identifying patterns or anomalies in web traffic.

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