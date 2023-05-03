# Table: servicenow_sys_group_has_role

Group Role.

## Examples

### What are the top 5 roles that are granted to groups?

```sql
select r.name as role_name, count(*) as num_granted_to_groups
from servicenow_sys_group_has_role ghr
left join servicenow_sys_user_role r on ghr.role->>'value' = r.sys_id
where role is not null
group by role_name
order by num_granted_to_groups desc
limit 5;
```

### What are the top 10 groups that have the most roles granted to them?

```sql
select ug.name as group_name, count(*) as num_granted_roles
from servicenow_sys_group_has_role ghr
left join servicenow_sys_user_group ug on ghr."group"->>'value' = ug.sys_id
where "group" is not null
group by group_name
order by num_granted_roles desc
limit 10;
```

### How many groups have roles that are inherited?

```sql
select count(distinct "group" ->> 'value') as num_groups_with_inherited_roles
from servicenow_sys_group_has_role
where inherits = true;
```

### What are the groups that have a specific role granted to them?

```sql
select ug.name as group_name
from servicenow_sys_group_has_role ghr
left join servicenow_sys_user_group ug on ghr."group"->>'value' = ug.sys_id
where role ->> 'value' = '282bf1fac6112285017366cb5f867469';
```

### How many groups have a specific role granted to them?

```sql
select count(distinct "group" ->> 'value') as num_groups_with_role
from servicenow_sys_group_has_role
where role ->> 'value' = 'ec1816c3871323004caf66d107cb0b1e'
and "group" is not null;
```

### How many roles are granted to each group?

```sql
select ug.name as group_name, count(*) as num_granted_roles
from servicenow_sys_group_has_role ghr
left join servicenow_sys_user_group ug on ghr."group"->>'value' = ug.sys_id
where "group" is not null
group by group_name
order by num_granted_roles desc;
```
