---
title: "Steampipe Table: prometheus_rule_group - Query Prometheus Rule Groups using SQL"
description: "Allows users to query Prometheus Rule Groups, specifically the configuration and status of each rule group, providing insights into monitoring and alerting rules."
---

# Table: prometheus_rule_group - Query Prometheus Rule Groups using SQL

Prometheus Rule Groups are a set of Prometheus rules that are evaluated together in a specific order. These rules can be either recording rules, which precompute frequently needed or computationally expensive expressions and save their result as a new set of time series, or alerting rules, which trigger alerts when certain conditions are observed to be true. Rule Groups are part of Prometheus' configuration and are used to define a list of alerts and recording rules.

## Table Usage Guide

The `prometheus_rule_group` table provides insights into Rule Groups within Prometheus. As a DevOps engineer, explore group-specific details through this table, including the rules, their configuration, and their current status. Utilize it to monitor and manage your Prometheus rules, understand the current state of your alerting and recording rules, and quickly identify any potential issues.

## Examples

### List all rule groups
Explore all the rule groups in your Prometheus monitoring system to understand how your rules are organized and to identify potential areas for optimization or reconfiguration. This can be especially beneficial for large-scale systems where efficient rule management is crucial.

```sql+postgres
select
  *
from
  prometheus_rule_group;
```

```sql+sqlite
select
  *
from
  prometheus_rule_group;
```