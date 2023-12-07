---
title: "Steampipe Table: servicenow_sys_user - Query ServiceNow Users using SQL"
description: "Allows users to query Users in ServiceNow, specifically user details and properties, providing insights into user management and operations."
---

# Table: servicenow_sys_user - Query ServiceNow Users using SQL

ServiceNow is a cloud-based platform designed to help businesses manage digital workflows for enterprise operations. It provides a centralized system to manage, coordinate, and develop effective workflows for business processes. ServiceNow includes a wide range of products and services, including IT service management, IT operations management, and IT business management.

## Table Usage Guide

The `servicenow_sys_user` table provides insights into user profiles within ServiceNow. As an IT manager or system administrator, explore user-specific details through this table, including roles, groups, and associated metadata. Utilize it to manage and monitor user activities, such as user roles and group memberships, and to verify user properties.

## Examples

### List users who haven't logged in for more than 90 days
Discover the segments that include users who haven't been active for more than 90 days. This can be useful in identifying inactive accounts for potential clean-up or outreach efforts.

```sql+postgres
select
  user_name,
  last_login_time 
from
  servicenow.servicenow_sys_user 
where
  last_login_time < now() - interval '90 DAYS';
```

```sql+sqlite
select
  user_name,
  last_login_time 
from
  servicenow_servicenow_sys_user 
where
  last_login_time < datetime('now', '-90 day');
```

### Users with Multifactor Authentication Disabled
Explore which users have not enabled multifactor authentication. This can be useful in identifying potential security risks within your system.

```sql+postgres
select
  count(*) as users_with_mfa_enabled 
from
  servicenow.servicenow_sys_user 
where
  enable_multifactor_authn = false;
```

```sql+sqlite
select
  count(*) as users_with_mfa_enabled 
from
  servicenow_servicenow_sys_user 
where
  enable_multifactor_authn = 0;
```

### Distribution of Users by Department
Discover the distribution of users across various departments. This query aids in understanding the user allocation in each department, providing valuable insights for resource planning and management.

```sql+postgres
select
  department,
  count(*) as num_users 
from
  servicenow.servicenow_sys_user 
group by
  department;
```

```sql+sqlite
select
  department,
  count(*) as num_users 
from
  servicenow_sys_user 
group by
  department;
```

### Users with Failed Login Attempts
Discover the segments that have experienced unsuccessful login attempts to gain insights into the average number of failed attempts per user. This information can aid in identifying potential security issues or areas for user experience improvement.

```sql+postgres
select
  count(*) as users_with_failed_attempts,
  avg(failed_attempts) as avg_failed_attempts 
from
  servicenow.servicenow_sys_user 
where
  failed_attempts > 0;
```

```sql+sqlite
select
  count(*) as users_with_failed_attempts,
  avg(failed_attempts) as avg_failed_attempts 
from
  servicenow_sys_user 
where
  failed_attempts > 0;
```

### Users with Profile Photo
Discover the number of users who have uploaded a profile photo on ServiceNow. This can be useful in identifying user engagement levels or in designing user-centric features.

```sql+postgres
select
  count(*) as users_with_photo 
from
  servicenow.servicenow_sys_user 
where
  photo is not null;
```

```sql+sqlite
select
  count(*) as users_with_photo 
from
  servicenow_sys_user 
where
  photo is not null;
```

### Distribution of Users by Country
Explore which countries have the most users to understand the global distribution of your user base. This can help in tailoring your services to cater to these regions more effectively.

```sql+postgres
select
  country,
  count(*) as num_users 
from
  servicenow.servicenow_sys_user 
group by
  country;
```

```sql+sqlite
select
  country,
  count(*) as num_users 
from
  servicenow_sys_user 
group by
  country;
```

### Users with No Manager Assigned
Discover the segments that have users without a manager assigned. This can be useful in identifying potential gaps in your team structure or areas that may require additional management oversight.

```sql+postgres
select
  count(*) as users_with_manager 
from
  servicenow.servicenow_sys_user 
where
  manager is null;
```

```sql+sqlite
select
  count(*) as users_with_manager 
from
  servicenow_sys_user 
where
  manager is null;
```

### Average Time Since Last Login
Understand the average duration since the last login for all active users. This can be useful for gauging user activity and identifying potential periods of inactivity.

```sql+postgres
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

```sql+sqlite
select 
  avg(strftime('%s', 'now') - strftime('%s', last_login_time)) as avg_time_since_last_login_in_milliseconds 
from 
  servicenow.servicenow_sys_user 
where 
  active = 1;
```

### Distribution of Users by Preferred Language
Explore the distribution of users based on their preferred language. This helps in understanding user preferences and tailoring language-specific services for better user experience.

```sql+postgres
select
  preferred_language,
  count(*) as num_users 
from
  servicenow.servicenow_sys_user 
group by
  preferred_language;
```

```sql+sqlite
select
  preferred_language,
  count(*) as num_users 
from
  servicenow_sys_user 
group by
  preferred_language;
```