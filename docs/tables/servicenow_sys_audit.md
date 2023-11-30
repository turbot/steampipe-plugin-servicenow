---
title: "Steampipe Table: servicenow_sys_audit - Query ServiceNow System Audit using SQL"
description: "Allows users to query System Audit records in ServiceNow, providing a comprehensive view of changes made to any record."
---

# Table: servicenow_sys_audit - Query ServiceNow System Audit using SQL

ServiceNow System Audit is a feature that tracks changes made to records within the ServiceNow platform. It provides a comprehensive view of modifications, including what was changed, who made the change, and when the change was made. This feature is essential for maintaining data integrity and accountability within ServiceNow.

## Table Usage Guide

The `servicenow_sys_audit` table provides insights into the audit records within ServiceNow System Audit. As a system administrator or compliance officer, explore audit-specific details through this table, including the type of changes, the individual responsible for the change, and the timestamp of the change. Utilize it to maintain accountability, ensure data integrity, and conduct forensic analysis when necessary.

## Examples

### What are the most frequent changes made in the table?
Analyze the most common modifications made within a system to understand recurring patterns and trends. This can help in identifying areas that require frequent updates, indicating potential areas for system optimization or process improvement.

```sql
select
  tablename,
  fieldname,
  count(*) as count 
from
  servicenow_sys_audit 
group by
  tablename,
  fieldname 
order by
  count desc limit 10;
```

### What are the most recent changes?
Uncover the details of the most recent modifications in your ServiceNow environment. This allows you to maintain an up-to-date understanding of changes, aiding in system management and troubleshooting.

```sql
select
  * 
from
  servicenow_sys_audit 
order by
  sys_created_on desc limit 10;
```

### Which user made the most changes?
Identify the user who has made the highest number of changes. This is useful for auditing purposes or to reward the most active contributors.

```sql
select
  sys_created_by,
  count(*) as count 
from
  servicenow_sys_audit 
group by
  sys_created_by 
order by
  count desc limit 10;
```

### How many records were modified on a specific date?
Explore the number of modifications made on a particular day. This can be useful in understanding the volume of changes in your system for specific dates, which can help with tracking activity and identifying potential anomalies.

```sql
select
  count(*) as count 
from
  servicenow_sys_audit 
where
  sys_created_on::date = '2023-05-04'::date;
```

### What are the changes made by a specific user?
Discover the modifications made by a particular individual. This is useful to track user activity and maintain accountability within your system.

```sql
select
  * 
from
  servicenow_sys_audit 
where
  sys_created_by = 'JohnDoe' 
order by
  sys_created_on desc limit 10;
```