---
title: "Steampipe Table: servicenow_sn_chg_rest_change_impacted_cmdb_ci_service - Query ServiceNow Change Impacted CMDB CI Services using SQL"
description: "Allows users to query Change Impacted CMDB CI Services in ServiceNow, providing detailed information about impacted configuration item services associated with a change."
---

# Table: servicenow_sn_chg_rest_change_impacted_cmdb_ci_service - Query ServiceNow Change Impacted CMDB CI Services using SQL

ServiceNow Change Management is a service that allows organizations to manage IT changes systematically so that they can negotiate change effectively. It provides a structured approach to control the lifecycle of all changes, facilitating beneficial changes to be made with minimum disruption to IT services. This includes managing changes to the Configuration Management Database (CMDB) and tracking impacted services.

## Table Usage Guide

The `servicenow_sn_chg_rest_change_impacted_cmdb_ci_service` table provides insights into the impacted configuration item services associated with a change in ServiceNow's Change Management module. As a Change Manager, you can use this table to track and analyze the impact of changes on different services, helping to minimize disruptions and ensure a smooth change implementation process. This table is especially useful for identifying dependencies and potential risks associated with changes to the CMDB.

## Examples

### How many CMDB CI Services are impacted by each change?
Determine the total number of Configuration Item (CI) Services impacted by each change. This helps in understanding the extent of the impact each change has on services, enabling better change management and mitigation planning.

```sql
select
  task_name,
  count(cmdb_ci_service_sys_id) as num_ci_impacted 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
group by
  task_name 
order by
  num_ci_impacted desc;
```

### What is the total number of manually added changes in the table?
Assess the extent of manual interventions in the change management process. This query helps in identifying the total number of changes that were manually added, providing insights into the degree of human involvement in the change management process.

```sql
select
  count(*) 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
where
  manually_added = true;
```

### Which CMDB CI Services were impacted by a specific change?
Explore which Configuration Management Database (CMDB) services were impacted by a specific change. This is useful in assessing the scope and potential impact of changes, helping to manage risk and ensure continuity of services.

```sql
select
  cmdb_ci_service_name,
  task_name 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
where
  task_name = 'CHG0000060';
```

### Which change impacted a specific CMDB CI Service?
Determine the specific tasks that impacted a particular Configuration Management Database (CMDB) service. This is useful in identifying changes that may have caused issues or disruptions in the service.

```sql
select
  task_name 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
where
  cmdb_ci_service_sys_id = '2216daf0d7820200c1ed0fbc5e6103ca';
```

### What is the total number of changes that impacted each CMDB CI Service?
Discover the segments that have been most impacted by changes, by assessing the total count of changes per service. This can help prioritize areas for review and potential improvement.

```sql
select
  cmdb_ci_service_name,
  count(task_name) as num_changes 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
group by
  cmdb_ci_service_name 
order by
  num_changes desc;
```

### Who created the most changes in the table?
Analyze the frequency of modifications to understand who has made the most changes. This can be useful for assessing individual workload or identifying frequent contributors in a collaborative environment.

```sql
select
  sys_created_by,
  count(*) as num_changes_created 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
group by
  sys_created_by 
order by
  num_changes_created desc;
```

### Which changes were added manually?
Explore which changes were made manually to your services. This can be useful to identify any potential unauthorized changes or inconsistencies in your system.

```sql
select
  task_name,
  cmdb_ci_service_name 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
where
  manually_added = true;
```

### What is the most common task in the table?
Discover the most frequently occurring task within your service operations by using this query. This could be beneficial in identifying patterns, optimizing workflows, or allocating resources more effectively.

```sql
select
  task,
  task_name,
  count(*) as num_tasks 
from
  servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
group by
  task,
  task_name 
order by
  num_tasks desc limit 1;
```

### What is the average number of CMDB CI Services impacted by each change?
Analyze the impact of each change to understand the average number of Configuration Item (CI) Services affected. This information can be used to assess the potential disruption caused by changes and plan accordingly.

```sql
select
  avg(num_ci_impacted) as avg_num_ci_impacted 
from
  (
    select
      task_name,
      count(cmdb_ci_service_sys_id) as num_ci_impacted 
    from
      servicenow_sn_chg_rest_change_impacted_cmdb_ci_service 
    group by
      task_name
  )
  as subquery;
```