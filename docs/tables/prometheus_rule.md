---
title: "Steampipe Table: prometheus_rule - Query Prometheus Rules using SQL"
description: "Allows users to query Prometheus Rules, specifically the rule configurations, providing insights into monitoring and alerting rules set in the Prometheus service."
---

# Table: prometheus_rule - Query Prometheus Rules using SQL

Prometheus is an open-source systems monitoring and alerting toolkit. It collects metrics from configured targets at given intervals, evaluates rule expressions, displays the results, and can trigger alerts if some condition is observed to be true. Prometheus's main features are a multi-dimensional data model with time series data identified by metric name and key/value pairs, a flexible query language to leverage this dimensionality, and no reliance on distributed storage.

## Table Usage Guide

The `prometheus_rule` table provides insights into rules within the Prometheus service. As a DevOps engineer, explore rule-specific details through this table, including rule configurations, alerting rules, and associated metadata. Utilize it to uncover information about rules, such as those triggering certain alerts, the conditions set for each rule, and the verification of rule expressions.

## Examples

### List all rules
Explore all the rules in your Prometheus setup, organized by group name and rule number. This can help you understand the structure and hierarchy of your rules for better management and troubleshooting.

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
This query is used to identify any rules within the Prometheus system that have been marked with a high severity label. This could be useful in prioritizing responses to system alerts or issues, by focusing on the most critical rules first.

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
Explore which rules are currently active or 'firing' within your Prometheus monitoring system. This can aid in identifying potential issues or anomalies in your network or system that require immediate attention.

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
Analyze your system's performance by pinpointing rules that are running slowly, taking more than one second for evaluation. This can help you identify potential bottlenecks or areas for optimization to improve overall system efficiency.

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