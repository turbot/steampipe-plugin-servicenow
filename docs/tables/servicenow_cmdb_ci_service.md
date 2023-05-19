# Table: servicenow_cmdb_ci_service

Configuration Item Service.

## Examples

### What are the top 10 most frequently used services?

```sql
select
  name,
  count(name) as frequency 
from
  servicenow_cmdb_ci_service 
group by
  name 
order by
  frequency desc limit 10;
```

### List services by category

```sql
select
  category,
  name 
from
  servicenow_cmdb_ci_service 
order by
  category,
  name;
```

### List services by status

```sql
select
  name,
  operational_status 
from
  servicenow_cmdb_ci_service;
```

### List services created in the last 30 days

```sql
select
  name,
  sys_created_on 
from
  servicenow_cmdb_ci_service 
where
  sys_created_on >= now() - interval '30 days';
```

### List services by the assigned user

```sql
select
  s.name,
  u.name 
from
  servicenow_cmdb_ci_service s 
  left join
    servicenow_sys_user u 
    on u.sys_id = s.assigned_to ->> 'value'
```
