---
title: "Steampipe Table: servicenow_sys_user_group - Query ServiceNow User Groups using SQL"
description: "Allows users to query User Groups in ServiceNow, providing insights into group details, such as the group's name, description, and manager."
---

# Table: servicenow_sys_user_group - Query ServiceNow User Groups using SQL

ServiceNow User Groups is a resource within ServiceNow that allows for the management and organization of users into specific groups. It provides a centralized way to manage and assign roles, responsibilities, and access permissions to specific groups of users. User Groups in ServiceNow helps in streamlining the process of user management, ensuring efficient distribution of tasks, and maintaining security protocols.

## Table Usage Guide

The `servicenow_sys_user_group` table provides insights into User Groups within ServiceNow. As an administrator or IT manager, you can explore group-specific details through this table, including group names, descriptions, and managers. Utilize it to manage user roles and responsibilities, monitor group activities, and enforce access control policies effectively.

## Examples

### Find all active groups
Discover the segments that are currently active in your ServiceNow user groups. This can be useful for managing user access and permissions in real-time.

```sql+postgres
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  active = true;
```

```sql+sqlite
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  active = 1;
```

### Count the number of groups created by each user
Analyze the distribution of group creation among users to understand who has been most actively involved in group formation. This could be useful for identifying key contributors or potential bottlenecks in your team structure.

```sql+postgres
select
  sys_created_by,
  count(*) as group_count 
from
  servicenow.servicenow_sys_user_group 
group by
  sys_created_by 
order by
  group_count desc;
```

```sql+sqlite
select
  sys_created_by,
  count(*) as group_count 
from
  servicenow_sys_user_group 
group by
  sys_created_by 
order by
  group_count desc;
```

### Count the number of groups with each type
Determine the distribution of various types within a system's user groups. This can be useful for understanding the structure and organization of your user groups.

```sql+postgres
select
  type,
  count(*) as group_count 
from
  servicenow.servicenow_sys_user_group 
group by
  type;
```

```sql+sqlite
select
  type,
  count(*) as group_count 
from
  servicenow_servicenow_sys_user_group 
group by
  type;
```

### Find all groups with name starting with HR
Identify all the user groups within a system that have names beginning with 'HR'. This could be useful for HR departments to quickly locate and manage all relevant groups in their organization.

```sql+postgres
select
  name,
  description 
from
  servicenow.servicenow_sys_user_group 
where
  name like 'HR%';
```

```sql+sqlite
select
  name,
  description 
from
  servicenow_sys_user_group 
where
  name like 'HR%';
```

### Find child groups of a parent group
Explore which child groups fall under a specific parent group. This is useful for understanding the organizational structure and hierarchy within a system.

```sql+postgres
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  parent ->> 'value' = 'ff0370019f22120047a2d126c42e702b';
```

```sql+sqlite
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  json_extract(parent, '$.value') = 'ff0370019f22120047a2d126c42e702b';
```

### Find all groups created by a specific user
Discover the groups that have been created by a specific user. This can be useful for auditing purposes or for understanding the user's role and responsibilities within the system.

```sql+postgres
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  sys_created_by = 'admin';
```

```sql+sqlite
select
  * 
from
  servicenow_servicenow_sys_user_group 
where
  sys_created_by = 'admin';
```

### Count the number of groups with each source
Discover the distribution of user groups across various sources by counting the number of groups associated with each source. This can help in assessing the diversity of group origins and identifying sources with a high or low number of groups.

```sql+postgres
select
  source,
  count(*) as group_count 
from
  servicenow.servicenow_sys_user_group 
group by
  source;
```

```sql+sqlite
select
  source,
  count(*) as group_count 
from
  servicenow_servicenow_sys_user_group 
group by
  source;
```

### Find all groups with a specific role
Identify all user groups that have been assigned a specific role in the ServiceNow system. This could be useful in managing user permissions and ensuring appropriate access rights.

```sql+postgres
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  roles like '%catalog%';
```

```sql+sqlite
select
  * 
from
  servicenow_servicenow_sys_user_group 
where
  roles like '%catalog%';
```

### Find all groups with a specific manager who is not excluded from email notifications
Identify all groups managed by a specific individual who isn't excluded from receiving email notifications. This is useful for ensuring that important communications are reaching the correct managerial personnel.

```sql+postgres
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  manager ->> 'value' = 'f298d2d2c611227b0106c6be7f154bc8' 
  and exclude_manager = false;
```

```sql+sqlite
select
  * 
from
  servicenow.servicenow_sys_user_group 
where
  json_extract(manager, '$.value') = 'f298d2d2c611227b0106c6be7f154bc8' 
  and exclude_manager = 0;
```