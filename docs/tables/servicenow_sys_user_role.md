# Table: servicenow_sys_user_role

Defines available roles in the ServiceNow.

## Examples

### Roles considered elevated privileges

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  elevated_privilege = true;
```

### Roles that can be granted independently

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  grantable = true;
```

### Roles that require a subscription

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  requires_subscription = 'yes';
```

### Roles that can be delegated

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  can_delegate = true;
```

### Scoped administrator roles

```sql
select
  name 
from
  servicenow_sys_user_role 
where
  scoped_admin = true;
```

### Roles that include other roles

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

```sql
select
  sys_id description 
from
  servicenow_sys_user_role 
where
  name = 'pdb_user';
```
