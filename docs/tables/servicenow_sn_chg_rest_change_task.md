---
title: "Steampipe Table: servicenow_sn_chg_rest_change_task - Query ServiceNow Change Tasks using SQL"
description: "Allows users to query Change Tasks in ServiceNow, specifically the details related to changes required in an IT infrastructure."
---

# Table: servicenow_sn_chg_rest_change_task - Query ServiceNow Change Tasks using SQL

ServiceNow Change Tasks are a part of the Change Management process in ServiceNow, where each task represents a set of actions required to implement a change in the IT infrastructure. These tasks are created as part of a Change Request and are assigned to various teams for implementation. They help in breaking down the change into manageable tasks and tracking the progress of each task separately.

## Table Usage Guide

The `servicenow_sn_chg_rest_change_task` table provides insights into Change Tasks within ServiceNow Change Management. As an IT Operations Manager, explore task-specific details through this table, including task status, assigned team, and associated metadata. Utilize it to track the progress of each change task, identify bottlenecks, and ensure timely completion of changes in the IT infrastructure.

## Examples


### What is the state_name distribution of change tasks in the servicenow_sn_chg_rest_change_task table?
Analyze the distribution of change tasks across various states in the Servicenow platform. This is useful to identify which states have the most change tasks, thereby providing insights into areas that may require more attention or resources.

```sql+postgres
select
  state_name,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  state_name 
order by
  count desc;
```

```sql+sqlite
select
  state_name,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  state_name 
order by
  count desc;
```

### How many change tasks have been created by each user in the servicenow_sn_chg_rest_change_task table?
Analyze the distribution of task creation in your system to understand the workload and productivity of each user. This could be useful for assessing individual contributions or identifying potential bottlenecks in your workflow.

```sql+postgres
select
  sys_created_by,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  sys_created_by;
```

```sql+sqlite
select
  sys_created_by,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  sys_created_by;
```

### How many change tasks have been assigned to each assignment group in the servicenow_sn_chg_rest_change_task table?
Analyze the distribution of assigned tasks among different groups in your service management system. This can help in understanding workload distribution and identifying any potential bottlenecks or uneven task allocation.

```sql+postgres
select
  assignment_group_name,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  assignment_group_name;
```

```sql+sqlite
select
  assignment_group_name,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  assignment_group_name;
```

### How many change tasks have been assigned to each user and what is their average priority?
Explore the distribution of task assignments and understand the average priority level assigned to each user. This can help in assessing workload and task importance in a team.

```sql+postgres
select
  assigned_to_name,
  count(*) as num_tasks_assigned,
  avg(priority) as avg_priority 
from
  servicenow_sn_chg_rest_change_task 
group by
  assigned_to_name;
```

```sql+sqlite
select
  assigned_to_name,
  count(*) as num_tasks_assigned,
  avg(priority) as avg_priority 
from
  servicenow_sn_chg_rest_change_task 
group by
  assigned_to_name;
```

### How many change tasks have been completed by each user in the last 30 days?
Discover the productivity of each user by counting the number of tasks they have completed in the last 30 days. This is beneficial in assessing individual performance and identifying high-performing team members.

```sql+postgres
select
  assigned_to_name,
  count(*) as num_tasks_completed 
from
  servicenow_sn_chg_rest_change_task 
where
  state = 3 
  and sys_updated_on >= now() - interval '30 days' 
group by
  assigned_to_name;
```

```sql+sqlite
select
  assigned_to_name,
  count(*) as num_tasks_completed 
from
  servicenow_sn_chg_rest_change_task 
where
  state = 3 
  and sys_updated_on >= datetime('now', '-30 days') 
group by
  assigned_to_name;
```