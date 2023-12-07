---
title: "Steampipe Table: servicenow_sn_chg_rest_change - Query ServiceNow Change Management Records using SQL"
description: "Allows users to query Change Management Records in ServiceNow, specifically the details of change requests, providing insights into the lifecycle of IT changes."
---

# Table: servicenow_sn_chg_rest_change - Query ServiceNow Change Management Records using SQL

ServiceNow Change Management is a service within the ServiceNow platform that enables organizations to manage all changes in IT infrastructure. It provides a structured approach to control and manage IT changes, minimizing the impact of change-related incidents and improving day-to-day operations. ServiceNow Change Management helps you maintain control over the entire change process, from creation, review, and approval, to scheduling, implementation, and post-implementation review.

## Table Usage Guide

The `servicenow_sn_chg_rest_change` table provides insights into change requests within ServiceNow Change Management. As an IT manager or administrator, explore change request details through this table, including the current state, priority, category, and associated metadata. Utilize it to monitor the lifecycle of IT changes, such as those pending approval, in progress, or completed, and ensure the controlled management of IT changes.

## Examples

### What are the most common category for change requests?
Discover the segments that have the highest frequency of change requests, helping prioritize areas for process improvement or resource allocation.

```sql+postgres
select
  category,
  count(*) as cnt 
from
  servicenow_sn_chg_rest_change 
group by
  category 
order by
  cnt desc limit 5;
```

```sql+sqlite
select
  category,
  count(*) as cnt 
from
  servicenow_sn_chg_rest_change 
group by
  category 
order by
  cnt desc limit 5;
```

### How many change requests are currently open for a specific category?
Discover the number of ongoing change requests in a specific category, such as 'network'. This is useful for tracking and managing workload related to different types of change requests.

```sql+postgres
select
  count(*) 
from
  servicenow_sn_chg_rest_change 
where
  state != 3 
  and category = 'network';
```

```sql+sqlite
select
  count(*) 
from
  servicenow_sn_chg_rest_change 
where
  state != 3 
  and category = 'network';
```

### How many changes are in the "work in progress" state?
Analyze the number of changes that are currently in the 'work in progress' state. This is useful for tracking progress and managing workload in a service management context.

```sql+postgres
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 2;
```

```sql+sqlite
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 2;
```

### What is the average duration of changes made by each user?
Analyze the average time taken by each user to make changes, which could be useful for assessing individual productivity or identifying inefficiencies in the change process.

```sql+postgres
select
  user,
  avg(duration) as avg_duration 
from
  servicenow_sn_chg_rest_change 
group by
  user;
```

```sql+sqlite
select
  user,
  avg(duration) as avg_duration 
from
  servicenow_sn_chg_rest_change 
group by
  user;
```

### How many changes are in the "scheduled" state and have a "high" priority?
Determine the volume of high-priority tasks that are currently scheduled. This can assist in gauging the workload and prioritizing resource allocation.

```sql+postgres
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 1 
  and priority = 2;
```

```sql+sqlite
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 1 
  and priority = 2;
```