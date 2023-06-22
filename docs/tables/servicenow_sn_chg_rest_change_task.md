# Table: servicenow_sn_chg_rest_change_task

Change Management - Change task.

## Examples


### What is the state_name distribution of change tasks in the servicenow_sn_chg_rest_change_task table?

```sql
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

```sql
select
  sys_created_by,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  sys_created_by;
```

### How many change tasks have been assigned to each assignment group in the servicenow_sn_chg_rest_change_task table?

```sql
select
  assignment_group_name,
  count(*) as count 
from
  servicenow_sn_chg_rest_change_task 
group by
  assignment_group_name;
```

### How many change tasks have been assigned to each user and what is their average priority?

```sql
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

```sql
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
