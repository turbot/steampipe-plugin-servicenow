# Table: servicenow_sn_chg_rest_change_affected_cmdb_ci

Change Management - CMDB CI affected by change.

## Examples

### How many CMDB CI items are affected by each change?

```sql
select
  task_name,
  count(ci_item_sys_id) as num_ci_affected 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  task_name 
order by
  num_ci_affected desc;
```

### What is the total number of manual proposed changes in the table?

```sql
select
  count(*) 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  manual_proposed_change = true;
```

### Which CMDB CI items were affected by a specific change?

```sql
select
  ci_item_name,
  task_name,
  applied_date 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  task_name = 'CHG0000060' 
order by
  applied_date desc;
```

### What is the distribution of applied and not applied changes in the table?

```sql
select
  applied,
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  applied;
```

### Which change affected a specific CMDB CI item?

```sql
select
  task_name,
  applied_date 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  ci_item_sys_id = '2216daf0d7820200c1ed0fbc5e6103ca' 
order by
  applied_date desc;
```

### What is the total number of changes that affected each CMDB CI item?

```sql
select
  ci_item_name,
  count(task_name) as num_changes 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  ci_item_name 
order by
  num_changes desc;
```

### Who created the most changes in the table?

```sql
select
  sys_created_by,
  count(*) as num_changes_created 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  sys_created_by 
order by
  num_changes_created desc;
```

### Which changes were applied manually?

```sql
select
  task_name,
  ci_item_name 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  manual_proposed_change = true;
```

### What is the most common task in the table?

```sql
select
  task,
  task_name,
  count(*) as num_tasks 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  task,
  task_name 
order by
  num_tasks desc limit 1;
```

### What is the average number of CMDB CI items affected by each change?

```sql
select
  avg(num_ci_affected) as avg_num_ci_affected 
from
  (
    select
      task_name,
      count(ci_item_sys_id) as num_ci_affected 
    from
      servicenow_sn_chg_rest_change_affected_cmdb_ci 
    group by
      task_name
  )
  as subquery;
```
