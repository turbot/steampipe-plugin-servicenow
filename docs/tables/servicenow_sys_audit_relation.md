---
title: "Steampipe Table: servicenow_sys_audit_relation - Query ServiceNow Audit Relations using SQL"
description: "Allows users to query Audit Relations in ServiceNow, specifically the creation, update, and deletion of relationships between ServiceNow tables."
---

# Table: servicenow_sys_audit_relation - Query ServiceNow Audit Relations using SQL

ServiceNow Audit Relations is a feature within the ServiceNow platform that tracks the creation, modification, and deletion of relationships between tables in the system. It provides a historical record of changes, enabling users to understand the evolution of relationships over time and identify any unauthorized or unexpected changes. The audit trail includes details such as the user who made the change, the date and time of the change, and the before and after values of the relationship.

## Table Usage Guide

The `servicenow_sys_audit_relation` table provides insights into the audit trails of relationships between tables in ServiceNow. As a system administrator or auditor, explore the details of these relationships through this table, including the user who made the change, the date and time of the change, and the before and after values of the relationship. Utilize it to maintain system integrity, ensure compliance with internal policies and external regulations, and detect any unauthorized or unexpected changes.

## Examples

### Retrieve all relations that were deleted
Discover the total number of changes made in the ServiceNow change management module, specifically focusing on high-priority items that are in a new or open state. This can be useful for change managers to track and manage change requests that require immediate attention.

```sql
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 1 
  and priority = 2;
```

### Retrieve all relations that were created
Explore the relations that were established within the ServiceNow system, specifically focusing on those that have undergone an audit but have not been deleted. This can be useful in understanding the history and integrity of your system's relationships.

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  audit is not null 
  and audit_delete is null;
```

### Retrieve all relations that were updated
Identify instances where changes have been made in the relations, which can be useful for auditing purposes or for tracking changes over time. This can be particularly beneficial in maintaining the integrity of the data and ensuring that any modifications are properly documented.

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  audit is not null 
  and audit_delete is not null;
```

### Retrieve all relations for a specific table
Analyze the relationships of a specific table to gain insights into how different elements within your database are interconnected. This is particularly useful for understanding dependencies and impacts in complex database structures.

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  tablename = 'incident';
```

### Retrieve all relations created by a specific user
Explore the relationships established by a particular user to gain insights into their activity and interactions within the system. This can be useful for auditing purposes or to understand user behavior.

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  sys_created_by = 'jsmith';
```

### Retrieve all relations created between a specific date range
Explore the relationships established within a specific timeframe to gain insights into system interactions and modifications during that period. This could be particularly useful for auditing purposes or understanding system changes over time.

```sql
select
  * 
from
  servicenow_sys_audit_relation 
where
  sys_created_on between '2022-01-01' and '2022-12-31';
```

### Retrieve the number of relations created by each user
Analyze the distribution of relations created by each user in the ServiceNow system. This helps in understanding user activity and identifying any unusual behavior or trends.

```sql
select
  sys_created_by,
  count(*) as relation_count 
from
  servicenow_sys_audit_relation 
group by
  sys_created_by;
```

### Retrieve the number of relations created per day
Explore the daily creation of relations to understand the frequency and volume of new connections being established within your ServiceNow system. This information can be useful for tracking system growth and identifying patterns or anomalies.

```sql
select
  date(sys_created_on),
  count(*) as relation_count 
from
  servicenow_sys_audit_relation 
group by
  date(sys_created_on);
```

### Retrieve the number of relations created per table
Explore the frequency of relationships established in each segment of your data. This can help you understand which areas of your data are most interconnected, aiding in the management and optimization of your data architecture.

```sql
select
  tablename,
  count(*) as relation_count 
from
  servicenow_sys_audit_relation 
group by
  tablename;
```