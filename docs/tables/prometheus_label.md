---
title: "Steampipe Table: prometheus_label - Query Prometheus Labels using SQL"
description: "Allows users to query Labels in Prometheus, specifically the metadata attached to timeseries data, providing insights into the metrics data and its dimensions."
---

# Table: prometheus_label - Query Prometheus Labels using SQL

Prometheus Labels are a type of metadata attached to timeseries data in Prometheus, a widely used open-source monitoring and alerting toolkit. Labels enable the identification of the metrics data and its dimensions, such as instance, job, etc. They play a crucial role in data querying, visualization, and aggregation.

## Table Usage Guide

The `prometheus_label` table provides insights into the labels used in Prometheus. As a DevOps engineer or a system administrator, you can explore label-specific details through this table, including the key-value pairs that identify the timeseries data. Utilize it to uncover information about the metrics data, such as the instance it belongs to, the job it is associated with, and other dimensions that help in effective data querying and visualization.

## Examples

### List all labels names
Explore all existing labels in your Prometheus monitoring system to understand the various classifications and groupings within your data. This can help in organizing and managing your system more effectively.

```sql
select
  name
from
  prometheus_label
```

### List all labels with their values
Explore which labels are associated with specific values in your Prometheus data. This can help you categorize and better understand your data for more effective management and analysis.

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