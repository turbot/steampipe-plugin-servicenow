---
title: "Steampipe Table: servicenow_sys_user_grmember - Query ServiceNow User Group Members using SQL"
description: "Allows users to query User Group Members in ServiceNow, specifically providing information about users and the groups they belong to in a ServiceNow instance."
---

# Table: servicenow_sys_user_grmember - Query ServiceNow User Group Members using SQL

ServiceNow is a cloud-based software platform that provides IT service management (ITSM) and helps automate IT Business Management (ITBM). It specializes in IT services management (ITSM), IT operations management (ITOM), and IT business management (ITBM). One of the key aspects of ServiceNow is its ability to manage User Group Members, providing a structured way to organize and manage users in a ServiceNow instance.

## Table Usage Guide

The `servicenow_sys_user_grmember` table provides insights into User Group Members within ServiceNow. As an IT administrator, you can explore details through this table, including the user's ID, the group's ID they belong to, and other related metadata. Utilize it to manage and organize users in your ServiceNow instance, such as identifying which groups a user belongs to, and verifying user-group relationships.

## Examples

### Which users belong to a particular group?
Determine the members of a specific group in your ServiceNow system. This can be useful for managing permissions, delegating tasks, or understanding the composition of your team.

```sql
select
  g.user ->> 'value' as user_id 
from
  servicenow_sys_user_grmember g 
where
  g.group ->> 'value' = 'ff0370019f22120047a2d126c42e702b';
```

### Which users were added to any group in the last 24 hours?
Explore the recent changes in user group memberships to identify any new additions within the last day. This can help monitor and manage group access rights, ensuring only authorized users are included.

```sql
select distinct
  g.user ->> 'value' as user_id 
from
  servicenow_sys_user_grmember g 
where
  g.sys_created_on >= now() - interval '24 hours';
```

### How many users are in each group?
Determine the distribution of users across various groups to understand user management and group allocation. This can be useful in identifying areas of high user concentration and potential need for group restructuring.

```sql
select
  g.name as group_name,
  count(distinct m.user ->> 'value') as user_count 
from
  servicenow_sys_user_grmember m 
  join
    servicenow_sys_user_group g 
    on m.group ->> 'value' = g.sys_id 
group by
  g.name;
```

### Which users are not members of any group?
Discover the segments that consist of users who are not associated with any group. This can be beneficial in identifying potential risks or gaps in user management and ensuring that all users are properly categorized for access control and auditing purposes.

```sql
select
  u.name as user_name 
from
  servicenow_sys_user u 
  left join
    servicenow_sys_user_grmember m 
    on u.sys_id = m.user ->> 'value' 
where
  m.sys_id is null;
```

### Which groups have no members?
Discover the segments that have no members, which can help in identifying unused or redundant groups in your system for better resource management.

```sql
select
  g.name as group_name 
from
  servicenow_sys_user_group g 
  left join
    servicenow_sys_user_grmember m 
    on g.sys_id = m.group ->> 'value' 
where
  m.sys_id is null;
```

### Which users are members of a specific group?
Determine the members of a specific group to better understand the distribution of roles and responsibilities within your organization. This is particularly useful in managing access controls and ensuring proper security protocols are followed.

```sql
select
  u.name as user_name,
  g.name as group_name 
from
  servicenow_sys_user_grmember m 
  join
    servicenow_sys_user u 
    on m.user ->> 'value' = u.sys_id 
  join
    servicenow_sys_user_group g 
    on m.group ->> 'value' = g.sys_id 
where
  g.name = 'it securities';
```