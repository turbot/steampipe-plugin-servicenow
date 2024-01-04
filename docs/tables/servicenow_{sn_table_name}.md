---
title: "Steampipe Table: servicenow_{sn_table_name} - Query ServiceNow tables using SQL"
description: "Allows users to query tables in ServiceNow, providing insights into service management data and potential service anomalies."
---

# Table: servicenow_{sn_table_name} - Query ServiceNow tables using SQL

ServiceNow is a cloud-based platform that provides solutions for IT Service Management (ITSM), IT Operations Management (ITOM), and IT Business Management (ITBM). It helps organizations manage digital workflows, thereby increasing productivity by creating, reading, and updating data for various service management processes. ServiceNow offers a range of IT service management options for on-premise and cloud-based applications.

## Table Usage Guide

The `servicenow_{sn_table_name}` table provides insights into ServiceNow tables. As a ServiceNow administrator or IT service manager, explore custom table-specific details, including status, assignment groups, and associated metadata. Utilize it to uncover information about ServiceNow tables, such as those with high priority, the assignment groups of tables, and the verification of service level agreements.

## Examples

### Inspect the table structure

Assuming your connection configuration is:

```hcl
connection "servicenow" {
  plugin = "servicenow"
  instance_url = "https://my-instance.service-now.com"
  username = "my-user"
  password = "my-password"
  objects = ["task", "u_my_custom_sn_table", "cmdb_model", "cmn_location"]
}
```

List all tables with:

```sh
.inspect servicenow
+---------------------------------+--------------------------------+
| table                           | description                    |
+---------------------------------+--------------------------------+
| servicenow_cmdb_model           | Product Model.                 |
| servicenow_cmn_location         | Location.                      |
| servicenow_task                 | Task.                          |
| servicenow_u_my_custom_sn_table | My custom table on ServiceNow. |
| (...)                           | (...)                          |
+---------------------------------+--------------------------------+
```

### How many tasks were closed in the last 30 days?
Determine the number of tasks that were finalized within the past month. This is useful for tracking productivity and understanding the recent workload.

```sql+postgres
select
  count(*) as closed_tasks 
from
  servicenow_task 
where
  state = '4' 
  and closed_at >= now() - interval '30 days';
```

```sql+sqlite
select
  count(*) as closed_tasks 
from
  servicenow_task 
where
  state = '4' 
  and closed_at >= datetime('now', '-30 days');
```

### How many tasks have been created by each user?
Discover the productivity levels of each user by analyzing the number of tasks they have created. This can be useful for workload management and identifying high-performing individuals.  

```sql+postgres
select
  sys_created_by,
  count(*) as num_tasks_created 
from
  servicenow_task 
group by
  sys_created_by 
order by
  sys_created_by;
```

```sql+sqlite
select
  sys_created_by,
  count(*) as num_tasks_created 
from
  servicenow_task 
group by
  sys_created_by 
order by
  sys_created_by;
```

### How many tasks have been opened in the last 24 hours?
Assess the volume of tasks initiated within the past day to understand recent workload and resource allocation needs. This can help in managing team capacity and planning for future task assignments.

```sql+postgres
select
  count(*) as num_tasks 
from
  servicenow_task 
where
  opened_at >= now() - interval '24 hours';
```

```sql+sqlite
select
  count(*) as num_tasks 
from
  servicenow_task 
where
  opened_at >= datetime('now', '-24 hours');
```