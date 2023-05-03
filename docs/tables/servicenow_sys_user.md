# Table: servicenow_sys_user

User.

## Examples

### List users who haven't logged in for more than 90 days:

```sql
SELECT user_name, last_login_time
FROM servicenow.servicenow_sys_user
WHERE last_login_time < NOW() - INTERVAL '90 DAYS';
```

### Users with Multifactor Authentication Disabled

```sql
SELECT COUNT(*) AS users_with_mfa_enabled
FROM servicenow.servicenow_sys_user
WHERE enable_multifactor_authn = false;
```

### Distribution of Users by Department

```sql
SELECT department, COUNT(*) AS num_users
FROM servicenow.servicenow_sys_user
GROUP BY department;
```

### Users with Failed Login Attempts

```sql
SELECT COUNT(*) AS users_with_failed_attempts, AVG(failed_attempts) AS avg_failed_attempts
FROM servicenow.servicenow_sys_user
WHERE failed_attempts > 0;
```

### Users with Profile Photo

```sql
SELECT COUNT(*) AS users_with_photo
FROM servicenow.servicenow_sys_user
WHERE photo IS NOT NULL;
```

### Distribution of Users by Country

```sql
SELECT country, COUNT(*) AS num_users
FROM servicenow.servicenow_sys_user
GROUP BY country;
```

### Users with No Manager Assigned

```sql
SELECT COUNT(*) AS users_with_manager
FROM servicenow.servicenow_sys_user
WHERE manager IS NULL;
```

### Average Time Since Last Login

```sql
SELECT AVG(EXTRACT(EPOCH FROM (NOW() - last_login_time))) AS avg_time_since_last_login_in_milliseconds
FROM servicenow.servicenow_sys_user
WHERE active = true;
```

### Distribution of Users by Preferred Language

```sql
SELECT preferred_language, COUNT(*) AS num_users
FROM servicenow.servicenow_sys_user
GROUP BY preferred_language;
```
