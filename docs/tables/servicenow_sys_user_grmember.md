# Table: servicenow_sys_user_grmember

User Group Member.

## Examples

### Which users belong to a particular group?

```sql
select g.user ->> 'value' as user_id
from servicenow_sys_user_grmember g
where g.group ->> 'value' = 'ff0370019f22120047a2d126c42e702b';
```

### Which users were added to any group in the last 24 hours?

```sql
select distinct g.user ->> 'value' as user_id
from servicenow_sys_user_grmember g
where g.sys_created_on >= now() - interval '24 hours';
```

### How many users are in each group?

```sql
select g.name as group_name, count(distinct m.user->>'value') as user_count
from servicenow_sys_user_grmember m
join servicenow_sys_user_group g on m.group->>'value' = g.sys_id
group by g.name;
```

### Which users are not members of any group?

```sql
select u.name as user_name
from servicenow_sys_user u
left join servicenow_sys_user_grmember m on u.sys_id = m.user->>'value'
where m.sys_id is null;
```

### Which groups have no members?

```sql
select g.name as group_name
from servicenow_sys_user_group g
left join servicenow_sys_user_grmember m on g.sys_id = m.group->>'value'
where m.sys_id is null;
```

### Which users are members of a specific group?

```sql
select u.name as user_name, g.name as group_name
from servicenow_sys_user_grmember m
join servicenow_sys_user u on m.user->>'value' = u.sys_id
join servicenow_sys_user_group g on m.group->>'value' = g.sys_id
where g.name = 'it securities';
```
