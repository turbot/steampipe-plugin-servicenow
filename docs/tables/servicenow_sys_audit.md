# Table: servicenow_sys_audit

Table change record.

## Examples

### What are the most frequent changes made in the table?

```sql
SELECT tablename, fieldname, COUNT(*) as count
FROM servicenow_sys_audit
GROUP BY tablename, fieldname
ORDER BY count DESC
LIMIT 10;
```

### What are the most recent changes?

```sql
SELECT *
FROM servicenow_sys_audit
ORDER BY sys_created_on DESC
LIMIT 10;
```

### Which user made the most changes?

```sql
SELECT sys_created_by, COUNT(*) as count
FROM servicenow_sys_audit
GROUP BY sys_created_by
ORDER BY count DESC
LIMIT 10;
```

### How many records were modified on a specific date?

```sql
SELECT COUNT(*) as count
FROM servicenow_sys_audit
WHERE sys_created_on::date = '2023-05-04'::date;
```

### What are the changes made by a specific user?

```sql
SELECT *
FROM servicenow_sys_audit
WHERE sys_created_by = 'JohnDoe'
ORDER BY sys_created_on DESC
LIMIT 10;
```
