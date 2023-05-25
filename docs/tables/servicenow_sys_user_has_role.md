# Table: servicenow_sys_user_has_role

Tracks assigned roles for users.

## Examples

### Which users have been granted a specific role through inheritance?

```sql
select
  uhr.user ->> 'value' as user_sys_id,
  u.name 
from
  servicenow_sys_user_has_role uhr 
  join
    servicenow_sys_user_role r 
    on uhr.role ->> 'value' = r.sys_id 
  join
    servicenow_sys_user u 
    on uhr.user ->> 'value' = u.sys_id 
where
  r.name = 'user_admin' 
  and uhr.inherited = true;
```

### What is the total number of roles granted?

```sql
select
  count(distinct role ->> 'value') as total_roles_granted 
from
  servicenow_sys_user_has_role;
```

### Which roles have been granted to a specific user?

```sql
select
  r.name as role_name 
from
  servicenow_sys_user_role r 
  join
    servicenow_sys_user_has_role uhr 
    on uhr.role ->> 'value' = r.sys_id 
where
  uhr.user ->> 'value' = 'd8f57f140b20220050192f15d6673a98';
```

### How many users have been granted each role?

```sql
select
  r.name as role_name,
  count(distinct uhr.user ->> 'value') as user_count 
from
  servicenow_sys_user_role r 
  join
    servicenow_sys_user_has_role uhr 
    on uhr.role ->> 'value' = r.sys_id 
group by
  r.name;
```

### Which users have been granted a role with elevated privileges?

```sql
select distinct
  uhr.user ->> 'value' as user_sys_id 
from
  servicenow_sys_user_has_role uhr 
  join
    servicenow_sys_user_role r 
    on uhr.role ->> 'value' = r.sys_id 
where
  r.elevated_privilege = true;
```
