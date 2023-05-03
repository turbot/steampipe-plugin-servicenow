# Table: servicenow_sn_chg_rest_change

Change Management - Change request.

## Examples

### What are the most common category for change requests?
```sql
SELECT category, COUNT(*) AS cnt 
FROM servicenow_sn_chg_rest_change 
GROUP BY category 
ORDER BY cnt DESC 
LIMIT 5;
```

### How many change requests are currently open for a specific category?
```sql
SELECT COUNT(*) 
FROM servicenow_sn_chg_rest_change 
WHERE state != 3 
AND category = 'network';
```

### How many changes are in the "work in progress" state?

```sql
SELECT COUNT(*) AS num_changes
FROM servicenow_sn_chg_rest_change
WHERE state = 2;
```

### What is the average duration of changes made by each user?

```sql
SELECT user, AVG(duration) AS avg_duration
FROM servicenow_sn_chg_rest_change
GROUP BY user;
```

### How many changes are in the "scheduled" state and have a "high" priority?

```sql
SELECT COUNT(*) AS num_changes
FROM servicenow_sn_chg_rest_change
WHERE state = 1 AND priority = 2;
```
