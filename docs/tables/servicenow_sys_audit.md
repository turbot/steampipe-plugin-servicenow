# Table: servicenow_sys_audit

Table change record.

## Examples

### What are the most frequent changes made in the table?

```sql
select
  tablename,
  fieldname,
  count(*) as count 
from
  servicenow_sys_audit 
group by
  tablename,
  fieldname 
order by
  count desc limit 10;
```

### What are the most recent changes?

```sql
select
  * 
from
  servicenow_sys_audit 
order by
  sys_created_on desc limit 10;
```

### Which user made the most changes?

```sql
select
  sys_created_by,
  count(*) as count 
from
  servicenow_sys_audit 
group by
  sys_created_by 
order by
  count desc limit 10;
```

### How many records were modified on a specific date?

```sql
select
  count(*) as count 
from
  servicenow_sys_audit 
where
  sys_created_on::date = '2023-05-04'::date;
```

### What are the changes made by a specific user?

```sql
select
  * 
from
  servicenow_sys_audit 
where
  sys_created_by = 'JohnDoe' 
order by
  sys_created_on desc limit 10;
```
