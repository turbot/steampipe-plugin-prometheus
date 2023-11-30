---
title: "Steampipe Table: prometheus_target - Query Prometheus Targets using SQL"
description: "Allows users to query Prometheus Targets, specifically the targets' state, health, and scrape pool. This provides insights into the target's operational status and performance."
---

# Table: prometheus_target - Query Prometheus Targets using SQL

Prometheus Target is a resource within Prometheus that represents an individual node or endpoint that Prometheus instances are scraping. It provides a detailed view of the target's state, health, and scrape pool, which can be used to monitor the operational status and performance of the target. By querying Prometheus Targets, users can gain insights into the metrics being scraped from each target and the health of the scraping process.

## Table Usage Guide

The `prometheus_target` table provides insights into individual nodes or endpoints within Prometheus. As a system administrator or a DevOps engineer, this table can be used to explore target-specific details, including its state, health, and scrape pool. Utilize it to uncover information about the performance of each target, the success or failure of the scraping process, and the metrics being collected from each target.

## Examples

### List all targets
Discover all the monitoring targets in your Prometheus setup, helping you gain a comprehensive view of what metrics are being tracked across your systems. This can be particularly useful for auditing or troubleshooting purposes.

```sql
select
  *
from
  prometheus_target
```

### Targets that are not up
Explore which targets in your Prometheus monitoring system are not currently operational. This can help you quickly identify and address any potential issues, improving system performance and reliability.

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
Identify instances where targets haven't been scanned in the last 24 hours. This is useful for maintaining up-to-date data and ensuring the health of your system.

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