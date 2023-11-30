---
title: "Steampipe Table: servicenow_sys_user_role - Query ServiceNow User Roles using SQL"
description: "Allows users to query User Roles in ServiceNow, specifically the details of roles assigned to users, providing insights into role-based access control and user permissions."
---

# Table: servicenow_sys_user_role - Query ServiceNow User Roles using SQL

ServiceNow User Roles is a feature within the ServiceNow platform that facilitates role-based access control (RBAC). It enables administrators to assign specific roles to users, thereby controlling their access to certain parts of the system. This feature is essential for managing user permissions and ensuring security within the ServiceNow environment.

## Table Usage Guide

The `servicenow_sys_user_role` table provides insights into user roles within ServiceNow. As an administrator, explore role-specific details through this table, including role names, descriptions, and associated users. Utilize it to uncover information about roles, such as those with excessive permissions, the distribution of roles among users, and the verification of user access rights.

## Examples

### Roles considered elevated privileges
Discover the segments that have elevated privileges in the ServiceNow user role system. This can be useful to identify potential risks or security concerns within your system.

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  elevated_privilege = true;
```

### Roles that can be granted independently
Identify roles within ServiceNow that are permitted to be granted independently. This allows for a more flexible and customizable management of user permissions within the system.

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  grantable = true;
```

### Roles that require a subscription
Determine the roles that necessitate a subscription within your ServiceNow environment. This can be useful for managing access and budgeting resources.

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  requires_subscription = 'yes';
```

### Roles that can be delegated
Identify instances where certain roles can be delegated within the ServiceNow system. This is useful in understanding the hierarchy and distribution of responsibilities within your organization.

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  can_delegate = true;
```

### Scoped administrator roles
Explore which roles in ServiceNow have been assigned scoped administrator privileges. This can help manage security and access control within your organization.

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  scoped_admin = true;
```

### Roles that include other roles
Discover the segments that have embedded roles within them in your ServiceNow user roles, to better manage and understand your system's access control hierarchy. This can be beneficial in identifying potential security risks or redundancies.

```sql
select
  name,
  includes_roles 
from
  servicenow_sys_user_role 
where
  includes_roles is not null;
```

### Roles requiring 'Assignable By' role
Discover the roles that require an 'Assignable By' role for allocation, enabling you to manage and delegate user permissions effectively.

```sql
select
  name,
  assignable_by 
from
  servicenow_sys_user_role 
where
  assignable_by is not null;
```

### Sys_id and description of a specific role
Determine the unique system identifier and description of a specific user role within a ServiceNow environment. This can be useful for understanding the permissions and capabilities associated with that role.

```sql
select
  sys_id description 
from
  servicenow_sys_user_role 
where
  name = 'pdb_user';
```