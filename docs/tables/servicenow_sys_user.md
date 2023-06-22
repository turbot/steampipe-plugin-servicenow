# Table: servicenow_sys_user

Manages user information in the ServiceNow.

## Examples

### List users who haven't logged in for more than 90 days

```sql
select
  user_name,
  last_login_time 
from
  servicenow.servicenow_sys_user 
where
  last_login_time < now() - interval '90 DAYS';
```

### Users with Multifactor Authentication Disabled

```sql
select
  count(*) as users_with_mfa_enabled 
from
  servicenow.servicenow_sys_user 
where
  enable_multifactor_authn = false;
```

### Distribution of Users by Department

```sql
select
  department,
  count(*) as num_users 
from
  servicenow.servicenow_sys_user 
group by
  department;
```

### Users with Failed Login Attempts

```sql
select
  count(*) as users_with_failed_attempts,
  avg(failed_attempts) as avg_failed_attempts 
from
  servicenow.servicenow_sys_user 
where
  failed_attempts > 0;
```

### Users with Profile Photo

```sql
select
  count(*) as users_with_photo 
from
  servicenow.servicenow_sys_user 
where
  photo is not null;
```

### Distribution of Users by Country

```sql
select
  country,
  count(*) as num_users 
from
  servicenow.servicenow_sys_user 
group by
  country;
```

### Users with No Manager Assigned

```sql
select
  count(*) as users_with_manager 
from
  servicenow.servicenow_sys_user 
where
  manager is null;
```

### Average Time Since Last Login

```sql
select
  avg(extract(epoch 
from
  (
    now() - last_login_time
  )
)) as avg_time_since_last_login_in_milliseconds 
from
  servicenow.servicenow_sys_user 
where
  active = true;
```

### Distribution of Users by Preferred Language

```sql
select
  preferred_language,
  count(*) as num_users 
from
  servicenow.servicenow_sys_user 
group by
  preferred_language;
```
