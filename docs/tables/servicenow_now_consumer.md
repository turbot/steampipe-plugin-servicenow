# Table: servicenow_now_consumer

Customer Service Management (CSM) consumer records.

## Examples

### What are the first and last names and email of all active consumers?
```sql
SELECT first_name, last_name, email
FROM servicenow_now_consumer
WHERE active = true;
```

### How many consumers are there in each state?
```sql
SELECT state, count(*) as num_consumers
FROM servicenow_now_consumer
GROUP BY state
order by state desc;
```

### What is the total number of consumers who doesn't have a mobile phone number listed?
```sql
SELECT COUNT(*)
FROM servicenow_now_consumer
WHERE mobile_phone IS NULL;
```

### What is the most common prefix for consumer names?
```sql
SELECT prefix, COUNT(*) as num_consumers
FROM servicenow_now_consumer
GROUP BY prefix
ORDER BY num_consumers DESC
LIMIT 1;
```

### How many consumers have an email address that contains the word "gmail"?
```sql
SELECT COUNT(*)
FROM servicenow_now_consumer
WHERE email LIKE '%gmail%';
```

### Which consumers have a title that contains the word "Manager"?
```sql
SELECT *
FROM servicenow_now_consumer
WHERE title LIKE '%Manager%';
```

1### How many consumers have a photo attached to their record?
```sql
SELECT COUNT(*)
FROM servicenow_now_consumer
WHERE photo IS NOT NULL;
```
