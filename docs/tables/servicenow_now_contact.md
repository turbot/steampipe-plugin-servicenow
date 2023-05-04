# Table: servicenow_now_contact

Customer Service Management (CSM) contact records.

## Examples

### How many contacts are in the `servicenow_now_contact` table?

```sql
SELECT COUNT(*) AS num_contacts
FROM servicenow_now_contact;
```

### What are the unique sources of the contacts?

```sql
SELECT DISTINCT source
FROM servicenow_now_contact;
```

### How many contacts were created in the last 7 days?

```sql
SELECT COUNT(*) AS num_contacts_last_7_days
FROM servicenow_now_contact
WHERE sys_created_on >= now() - interval '7 days'
```

### What is the most common country of origin for the contacts?

```sql
SELECT country, COUNT(*) AS num_contacts
FROM servicenow_now_contact
GROUP BY country
ORDER BY num_contacts DESC
LIMIT 1;
```

### How many contacts were created by each source in the last 30 days?

```sql
SELECT source, COUNT(*) AS num_contacts_last_30_days
FROM servicenow_now_contact
WHERE sys_created_on >= now() - interval '30 days'
GROUP BY source;
```

### What is the average number of contacts created per day in the last 90 days?

```sql
SELECT AVG(num_contacts_per_day) AS avg_contacts_per_day_last_90_days
FROM (
  SELECT COUNT(*) AS num_contacts_per_day
  FROM servicenow_now_contact
  WHERE sys_created_on >= now() - interval '90 days'
  GROUP BY sys_created_on
) AS daily_contacts;
```
