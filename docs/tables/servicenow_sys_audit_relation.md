# Table: servicenow_sys_audit_relation

Table relationship audit record.

## Examples

### Retrieve all relations that were deleted:

```sql
SELECT *
FROM servicenow_sys_audit_relation
WHERE audit_delete IS NOT NULL;
```

### Retrieve all relations that were created:

```sql
SELECT *
FROM servicenow_sys_audit_relation
WHERE audit IS NOT NULL AND audit_delete IS NULL;
```

### Retrieve all relations that were updated:

```sql
SELECT *
FROM servicenow_sys_audit_relation
WHERE audit IS NOT NULL AND audit_delete IS NOT NULL;
```

### Retrieve all relations for a specific table:

```sql
SELECT *
FROM servicenow_sys_audit_relation
WHERE tablename = 'incident';
```

### Retrieve all relations created by a specific user:

```sql
SELECT *
FROM servicenow_sys_audit_relation
WHERE sys_created_by = 'jsmith';
```

### Retrieve all relations created between a specific date range:

```sql
SELECT *
FROM servicenow_sys_audit_relation
WHERE sys_created_on BETWEEN '2022-01-01' AND '2022-12-31';
```

### Retrieve the number of relations created by each user:

```sql
SELECT sys_created_by, COUNT(*) as relation_count
FROM servicenow_sys_audit_relation
GROUP BY sys_created_by;
```

### Retrieve the number of relations created per day:

```sql
SELECT DATE(sys_created_on), COUNT(*) as relation_count
FROM servicenow_sys_audit_relation
GROUP BY DATE(sys_created_on);
```

### Retrieve the number of relations created per table:

```sql
SELECT tablename, COUNT(*) as relation_count
FROM servicenow_sys_audit_relation
GROUP BY tablename;
```
