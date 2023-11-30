---
title: "Steampipe Table: servicenow_sys_group_has_role - Query ServiceNow Sys Group Roles using SQL"
description: "Allows users to query Sys Group Roles in ServiceNow, providing insights into the roles assigned to specific groups within the ServiceNow platform."
---

# Table: servicenow_sys_group_has_role - Query ServiceNow Sys Group Roles using SQL

ServiceNow is a cloud-based platform designed to help manage digital workflows for enterprise operations. It provides a suite of applications and tools for IT service management (ITSM), IT operations management (ITOM), and IT business management (ITBM). Within the ServiceNow platform, Sys Group Roles represent the roles assigned to specific groups, which determine the access permissions and capabilities of the group members within the platform.

## Table Usage Guide

The `servicenow_sys_group_has_role` table provides insights into the roles assigned to groups within the ServiceNow platform. As a system administrator or IT manager, you can utilize this table to understand the access permissions and capabilities of different groups, aiding in effective access management and security compliance. It can also be used to identify any discrepancies in role assignments, ensuring that each group has the appropriate roles for their function.

## Examples

### What are the top 5 roles that are granted to groups?
Discover the most commonly assigned roles within a group to understand the distribution of responsibilities and privileges. This can be useful for auditing purposes or for optimizing group management strategies.

```sql
select
  r.name as role_name,
  count(*) as num_granted_to_groups 
from
  servicenow_sys_group_has_role ghr 
  left join
    servicenow_sys_user_role r 
    on ghr.role ->> 'value' = r.sys_id 
where
  role is not null 
group by
  role_name 
order by
  num_granted_to_groups desc limit 5;
```

### What are the top 10 groups that have the most roles granted to them?
Discover the segments that have the most roles assigned to them in an organization. This is useful for understanding which groups have the most responsibilities or privileges, aiding in access management and security planning.

```sql
select
  ug.name as group_name,
  count(*) as num_granted_roles 
from
  servicenow_sys_group_has_role ghr 
  left join
    servicenow_sys_user_group ug 
    on ghr."group" ->> 'value' = ug.sys_id 
where
  "group" is not null 
group by
  group_name 
order by
  num_granted_roles desc limit 10;
```

### How many groups have roles that are inherited?
Determine the number of groups that have inherited roles in ServiceNow. This can be used to assess the extent of role inheritance within your organization's groups, which may be useful for managing permissions and access controls.

```sql
select
  count(distinct "group" ->> 'value') as num_groups_with_inherited_roles 
from
  servicenow_sys_group_has_role 
where
  inherits = true;
```

### What are the groups that have a specific role granted to them?
Determine the areas in which certain groups have been granted a specific role. This is useful in managing access control and ensuring appropriate permissions are in place.

```sql
select
  ug.name as group_name 
from
  servicenow_sys_group_has_role ghr 
  left join
    servicenow_sys_user_group ug 
    on ghr."group" ->> 'value' = ug.sys_id 
where
  role ->> 'value' = '282bf1fac6112285017366cb5f867469';
```

### How many groups have a specific role granted to them?
Determine the number of groups that have been assigned a particular role in your ServiceNow system. This can be useful for assessing access control and ensuring appropriate permissions are distributed.

```sql
select
  count(distinct "group" ->> 'value') as num_groups_with_role 
from
  servicenow_sys_group_has_role 
where
  role ->> 'value' = 'ec1816c3871323004caf66d107cb0b1e' 
  and "group" is not null;
```

### How many roles are granted to each group?
Determine the distribution of roles across different groups, identifying which groups have been assigned the most roles. This can help in understanding role allocation patterns and identifying any potential discrepancies or over-allocations.

```sql
select
  ug.name as group_name,
  count(*) as num_granted_roles 
from
  servicenow_sys_group_has_role ghr 
  left join
    servicenow_sys_user_group ug 
    on ghr."group" ->> 'value' = ug.sys_id 
where
  "group" is not null 
group by
  group_name 
order by
  num_granted_roles desc;
```