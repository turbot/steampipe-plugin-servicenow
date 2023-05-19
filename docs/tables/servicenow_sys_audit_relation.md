# Table: servicenow_sys_audit_relation

Table relationship audit record.

## Examples

### Retrieve all relations that were deleted

```sql
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 1 
  and priority = 2;
```

### Retrieve all relations that were created

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  audit is not null 
  and audit_delete is null;
```

### Retrieve all relations that were updated

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  audit is not null 
  and audit_delete is not null;
```

### Retrieve all relations for a specific table

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  tablename = 'incident';
```

### Retrieve all relations created by a specific user

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  sys_created_by = 'jsmith';
```

### Retrieve all relations created between a specific date range

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  sys_created_on between '2022-01-01' and '2022-12-31';
```

### Retrieve the number of relations created by each user

```sql
select
  sys_created_by,
  count(*) as relation_count 
from
  servicenow_sys_audit_relation 
group by
  sys_created_by;
```

### Retrieve the number of relations created per day

```sql
select
  date(sys_created_on),
  count(*) as relation_count 
from
  servicenow_sys_audit_relation 
group by
  date(sys_created_on);
```

### Retrieve the number of relations created per table

```sql
select
  tablename,
  count(*) as relation_count 
from
  servicenow_sys_audit_relation 
group by
  tablename;
```
