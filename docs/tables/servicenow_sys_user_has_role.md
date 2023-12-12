---
title: "Steampipe Table: servicenow_sys_user_has_role - Query ServiceNow User Roles using SQL"
description: "Allows users to query User Roles in ServiceNow, specifically the relationship between users and their assigned roles, providing insights into user access levels and permissions."
---

# Table: servicenow_sys_user_has_role - Query ServiceNow User Roles using SQL

ServiceNow User Roles is a feature within ServiceNow that allows administrators to manage access levels and permissions for individual users. It provides a structured way to assign and monitor roles for various users, including administrators, developers, and IT support staff. ServiceNow User Roles helps in maintaining security and efficient workflow by ensuring that users have appropriate access to perform their tasks.

## Table Usage Guide

The `servicenow_sys_user_has_role` table provides insights into User Roles within ServiceNow. As a system administrator or IT manager, explore user-specific role details through this table, including assigned roles and related metadata. Utilize it to uncover information about user access levels, such as those with administrative permissions, the relationships between users and their roles, and the verification of user permissions.

## Examples

### Which users have been granted a specific role through inheritance?
This query allows you to identify which users have been assigned a specific administrative role through inheritance. It's useful in managing user permissions and understanding the distribution of administrative roles in your system.

```sql+postgres
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

```sql+sqlite
select
  json_extract(uhr.user, '$.value') as user_sys_id,
  u.name 
from
  servicenow_sys_user_has_role uhr 
  join
    servicenow_sys_user_role r 
    on json_extract(uhr.role, '$.value') = r.sys_id 
  join
    servicenow_sys_user u 
    on json_extract(uhr.user, '$.value') = u.sys_id 
where
  r.name = 'user_admin' 
  and uhr.inherited = 1;
```

### What is the total number of roles granted?
Analyze the settings to understand the total count of unique roles granted to users in your ServiceNow instance. This can help in assessing the distribution of responsibilities and privileges within your team.

```sql+postgres
select
  count(distinct role ->> 'value') as total_roles_granted 
from
  servicenow_sys_user_has_role;
```

```sql+sqlite
select
  count(distinct json_extract(role.value, '$.value')) as total_roles_granted 
from
  servicenow_sys_user_has_role,
  json_each(role);
```

### Which roles have been granted to a specific user?
Determine the areas in which a particular user has been assigned roles. This helps in understanding the user's permissions and access levels within the system.

```sql+postgres
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

```sql+sqlite
select
  r.name as role_name 
from
  servicenow_sys_user_role r,
  servicenow_sys_user_has_role uhr 
where
  json_extract(uhr.role, '$.value') = r.sys_id 
  and json_extract(uhr.user, '$.value') = 'd8f57f140b20220050192f15d6673a98';
```

### How many users have been granted each role?
Determine the distribution of user roles within your system to understand the level of access granted to different users. This can aid in security audits by identifying potential over-privileged users.

```sql+postgres
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

```sql+sqlite
select
  r.name as role_name,
  count(distinct json_extract(uhr.user, '$.value')) as user_count 
from
  servicenow_sys_user_role r 
  join
    servicenow_sys_user_has_role uhr 
    on json_extract(uhr.role, '$.value') = r.sys_id 
group by
  r.name;
```

### Which users have been granted a role with elevated privileges?
Identify instances where users have been given roles with increased access rights. This can be useful for auditing purposes, ensuring only authorized personnel have such privileges.

```sql+postgres
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

```sql+sqlite
select distinct
  json_extract(uhr.user, '$.value') as user_sys_id 
from
  servicenow_sys_user_has_role uhr 
  join
    servicenow_sys_user_role r 
    on json_extract(uhr.role, '$.value') = r.sys_id 
where
  r.elevated_privilege = 1;
```