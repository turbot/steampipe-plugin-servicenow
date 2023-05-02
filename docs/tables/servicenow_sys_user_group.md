# Table: servicenow_sys_user_group

User Group.

## Examples

### Find all active groups

```sql
select * from servicenow.servicenow_sys_user_group
where active = true;
```

### Count the number of groups created by each user

```sql
select sys_created_by, count(*) as group_count
from servicenow.servicenow_sys_user_group
group by sys_created_by
order by group_count desc;
```

### Count the number of groups with each type

```sql
select type, count(*) as group_count
from servicenow.servicenow_sys_user_group
group by type;
```

### Find all groups with name starting with HR

```sql
select name, description
from servicenow.servicenow_sys_user_group
where name like 'HR%';
```

### Find child groups of a parent group

```sql
select * from servicenow.servicenow_sys_user_group
where parent ->> 'value' = 'ff0370019f22120047a2d126c42e702b';
```

### Find all groups created by a specific user

```sql
select * from servicenow.servicenow_sys_user_group
where sys_created_by = 'admin';
```

### Count the number of groups with each source

```sql
select source, count(*) as group_count from servicenow.servicenow_sys_user_group
group by source;
```

### Find all groups with a specific role

```sql
select * from servicenow.servicenow_sys_user_group
where roles like '%catalog%';
```

### Find all groups with a specific manager who is not excluded from email notifications

```sql
select * from servicenow.servicenow_sys_user_group
where manager ->> 'value' = 'f298d2d2c611227b0106c6be7f154bc8' and exclude_manager = false;
```
