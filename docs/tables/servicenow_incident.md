---
title: "Steampipe Table: servicenow_incident - Query ServiceNow Incidents using SQL"
description: "Allows users to query ServiceNow Incidents, providing detailed information about each incident recorded in the ServiceNow platform."
---

# Table: servicenow_incident - Query ServiceNow Incidents using SQL

ServiceNow Incident Management is a service within the ServiceNow platform that helps organizations respond to, resolve, and report on IT incidents. It aims to restore service operation to normal as quickly as possible and minimize the adverse impact on business operations. This ensures that the highest possible levels of service quality and availability are maintained.

## Table Usage Guide

The `servicenow_incident` table provides insights into each incident recorded within the ServiceNow Incident Management system. As an IT operations manager, you can explore incident-specific details through this table, including incident state, priority, and associated metadata. Utilize it to uncover information about incidents, such as their current status, assigned personnel, and resolution progress.

## Examples

### What is the distribution of incidents by severity level?
Gain insights into the distribution of incidents according to their severity level. This query helps in understanding the proportion of each severity level in relation to the total incidents, aiding in effective incident management.

```sql
select
  severity,
  count(*) as incident_count,
  count(*) * 100.0 / sum(count(*)) over() as percentage 
from
  servicenow_incident 
group by
  severity 
order by
  severity;
```

### What is the average time taken to resolve incidents by category?
Explore the efficiency of incident resolution processes by determining the average time taken to resolve incidents, grouped by their respective categories. This can provide valuable insights into areas where response times may need improvement.

```sql
select
  category,
  avg(calendar_stc) as avg_time_to_resolve 
from
  servicenow_incident 
where
  incident_state = 7 	-- only include resolved incidents
group by
  category;
```

### What was the most common reason for holding incidents?
Analyze the incidents to understand the most frequent reason for their hold status. This can help in identifying recurring issues and implementing preventive measures.

```sql
select
  hold_reason,
  count(*) as incident_count 
from
  servicenow_incident 
where
  hold_reason is not null 
group by
  hold_reason 
order by
  incident_count desc limit 1;
```

### What is the average time taken to resolve incidents by severity level?
Analyze the average resolution time for incidents, grouped by severity level. This can be useful to identify areas of improvement in incident management and to prioritize resources based on severity.

```sql
select
  severity,
  avg(calendar_stc) as avg_time_to_resolve 
from
  servicenow_incident 
where
  incident_state = 7 	-- only include resolved incidents
group by
  severity;
```

### What are the top 10 categories of incidents by count?
Uncover the details of the most frequently occurring incident categories to better prioritize and strategize your response efforts. This can help in efficiently allocating resources and improving response times to critical issues.

```sql
select
  category,
  count(*) as incident_count 
from
  servicenow_incident 
group by
  category 
order by
  incident_count desc limit 10;
```

### How many incidents were resolved in the last 30 days?
Analyze the volume of resolved incidents over the past month. This is useful for assessing the efficiency and effectiveness of your incident response team.

```sql
select
  count(*) as resolved_incidents 
from
  servicenow_incident 
where
  incident_state = '6' 
  and resolved_at >= now() - interval '30 days';
```

### How many incidents were reopened more than once?
Assess the frequency of incidents that have been reopened multiple times, providing insights into potential issues with problem resolution or customer satisfaction.

```sql
select
  count(*) as reopened_incidents 
from
  servicenow_incident 
where
  reopen_count > 1;
```