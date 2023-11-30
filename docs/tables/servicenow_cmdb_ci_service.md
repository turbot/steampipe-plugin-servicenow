---
title: "Steampipe Table: servicenow_cmdb_ci_service - Query ServiceNow Configuration Management Database (CMDB) Services using SQL"
description: "Allows users to query ServiceNow CMDB Services, providing insights into service-related information, including operational status, version, and associated business services."
---

# Table: servicenow_cmdb_ci_service - Query ServiceNow Configuration Management Database (CMDB) Services using SQL

ServiceNow Configuration Management Database (CMDB) is a service that acts as a data warehouse for IT installations. It holds data relating to a collection of IT assets (commonly referred to as configuration items (CI)), as well as to descriptive relationships between such assets. A key functionality of CMDB is to support the IT department's configuration management process.

## Table Usage Guide

The `servicenow_cmdb_ci_service` table provides insights into services within ServiceNow Configuration Management Database (CMDB). As an IT administrator or a DevOps engineer, explore service-specific details through this table, including operational status, version, and associated business services. Utilize it to uncover information about services, such as those that are currently active, the versions of the services, and the business services associated with them.

## Examples

### What are the top 10 most frequently used services?
Explore the most commonly used services in your organization to help prioritize resource allocation and streamline operations. This query can be particularly useful in identifying areas that may benefit from increased investment or optimization.

```sql
select
  name,
  count(name) as frequency 
from
  servicenow_cmdb_ci_service 
group by
  name 
order by
  frequency desc limit 10;
```

### List services by category
Explore which services fall under each category to better understand the organization and classification of your services. This can help in managing resources and planning future developments.

```sql
select
  category,
  name 
from
  servicenow_cmdb_ci_service 
order by
  category,
  name;
```

### List services by status
Explore which services are currently operational by assessing their status. This can help in identifying any services that might be experiencing issues, thereby allowing for timely troubleshooting and maintenance.

```sql
select
  name,
  operational_status 
from
  servicenow_cmdb_ci_service;
```

### List services created in the last 30 days
Explore which services have been recently added by identifying those created within the past month. This can help in tracking the growth and evolution of your service catalog over time.

```sql
select
  name,
  sys_created_on 
from
  servicenow_cmdb_ci_service 
where
  sys_created_on >= now() - interval '30 days';
```

### List services by the assigned user
Explore which services are assigned to which users, helping to manage workload distribution and responsibility tracking. This is beneficial for understanding the current state of task allocation within your team.

```sql
select
  s.name,
  u.name 
from
  servicenow_cmdb_ci_service s 
  left join
    servicenow_sys_user u 
    on u.sys_id = s.assigned_to ->> 'value'
```