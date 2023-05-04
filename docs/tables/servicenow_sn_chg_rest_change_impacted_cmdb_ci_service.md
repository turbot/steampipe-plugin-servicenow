# Table: servicenow_sn_chg_rest_change_impacted_cmdb_ci_service

Change Management - CMDB CI Service impacted by change.

## Examples

### How many CMDB CI Services are impacted by each change?

```sql
select task_name, count(cmdb_ci_service_sys_id) as num_ci_impacted
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
group by task_name
order by num_ci_impacted desc;
```

### What is the total number of manually added changes in the table?

```sql
select count(*)
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
where manually_added = true;
```

### Which CMDB CI Services were impacted by a specific change?

```sql
select cmdb_ci_service_name, task_name
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
where task_name = 'CHG0000060';
```

### Which change impacted a specific CMDB CI Service?

```sql
select task_name
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
where cmdb_ci_service_sys_id = '2216daf0d7820200c1ed0fbc5e6103ca';
```

### What is the total number of changes that impacted each CMDB CI Service?

```sql
select cmdb_ci_service_name, count(task_name) as num_changes
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
group by cmdb_ci_service_name
order by num_changes desc;
```

### Who created the most changes in the table?

```sql
select sys_created_by, count(*) as num_changes_created
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
group by sys_created_by
order by num_changes_created desc;
```

### Which changes were added manually?

```sql
select task_name, cmdb_ci_service_name
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
where manually_added = true;
```

### What is the most common task in the table?

```sql
select task, task_name, count(*) as num_tasks
from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
group by task, task_name
order by num_tasks desc
limit 1;
```

### What is the average number of CMDB CI Services impacted by each change?

```sql
select avg(num_ci_impacted) as avg_num_ci_impacted
from (select task_name, count(cmdb_ci_service_sys_id) as num_ci_impacted
      from servicenow_sn_chg_rest_change_impacted_cmdb_ci_service
      group by task_name) as subquery;
```
