---
title: "Steampipe Table: servicenow_sn_chg_rest_change_affected_cmdb_ci - Query ServiceNow Change Management Configuration Items using SQL"
description: "Allows users to query Change Management Configuration Items in ServiceNow, specifically the ones affected by changes, providing insights into change impact and related configuration items."
---

# Table: servicenow_sn_chg_rest_change_affected_cmdb_ci - Query ServiceNow Change Management Configuration Items using SQL

ServiceNow Change Management is a service within ServiceNow that helps organizations manage and control IT infrastructure changes. It provides a systematic approach to control the lifecycle of all changes, facilitating beneficial changes to be made with minimal disruption to IT services. Configuration Items (CI) in ServiceNow represent specific, tangible objects in an environment such as hardware, software, and services.

## Table Usage Guide

The `servicenow_sn_chg_rest_change_affected_cmdb_ci` table provides insights into Configuration Items (CI) affected by changes within ServiceNow Change Management. As an IT service manager, explore CI-specific details through this table, including CI's impacted by specific changes, their relationships, and associated metadata. Utilize it to uncover information about the impact of changes, such as those affecting critical services, the relationships between CIs, and verification of change plans.

## Examples

### How many CMDB CI items are affected by each change?
Determine the impact of each change on your Configuration Management Database (CMDB) by identifying the number of Configuration Items (CI) affected by each task. This can help prioritize tasks based on the scale of their impact.

```sql
select
  task_name,
  count(ci_item_sys_id) as num_ci_affected 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  task_name 
order by
  num_ci_affected desc;
```

### What is the total number of manual proposed changes in the table?
Explore the total count of proposed changes that have been manually inputted, providing a quick overview of interventions that may require further review or approval. This could be useful in assessing the volume of manual interventions and their potential impact on system stability.

```sql
select
  count(*) 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  manual_proposed_change = true;
```

### Which CMDB CI items were affected by a specific change?
Determine the configuration items (CI) affected by a specific change in your ServiceNow Change Management Database (CMDB). This can be useful for understanding the impact of changes, allowing for more informed decision making.

```sql
select
  ci_item_name,
  task_name,
  applied_date 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  task_name = 'CHG0000060' 
order by
  applied_date desc;
```

### What is the distribution of applied and not applied changes in the table?
Explore the distribution of applied and unapplied changes to understand the overall change management process. This can help identify potential bottlenecks and areas for improvement in the change application process.

```sql
select
  applied,
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  applied;
```

### Which change affected a specific CMDB CI item?
Analyze the settings to understand which modifications impacted a particular configuration item (CI) in your Configuration Management Database (CMDB). This is particularly useful for tracking changes and troubleshooting issues related to specific CIs.

```sql
select
  task_name,
  applied_date 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  ci_item_sys_id = '2216daf0d7820200c1ed0fbc5e6103ca' 
order by
  applied_date desc;
```

### What is the total number of changes that affected each CMDB CI item?
Determine the frequency of changes impacting each configuration item to assess the stability and potential risk areas in your IT environment. This can help prioritize areas for improvement and risk mitigation.

```sql
select
  ci_item_name,
  count(task_name) as num_changes 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  ci_item_name 
order by
  num_changes desc;
```

### Who created the most changes in the table?
Discover the users who have made the most changes in a system, providing a way to identify key contributors or potential sources of system instability. This information can be useful in managing system maintenance and troubleshooting.

```sql
select
  sys_created_by,
  count(*) as num_changes_created 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  sys_created_by 
order by
  num_changes_created desc;
```

### Which changes were applied manually?
Determine the areas in which changes were manually applied to understand potential risks and ensure proper change management protocols were followed. This is beneficial in maintaining system integrity and avoiding unexpected issues due to manual interventions.

```sql
select
  task_name,
  ci_item_name 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
where
  manual_proposed_change = true;
```

### What is the most common task in the table?
Explore which task appears most frequently within a certain service, providing insight into the most common operation or action that occurs within that context. This can help in identifying areas for process optimization or resource allocation.

```sql
select
  task,
  task_name,
  count(*) as num_tasks 
from
  servicenow_sn_chg_rest_change_affected_cmdb_ci 
group by
  task,
  task_name 
order by
  num_tasks desc limit 1;
```

### What is the average number of CMDB CI items affected by each change?
Determine the average number of Configuration Items (CI) impacted by each change in your ServiceNow change management process. This can help in assessing the potential impact and risk of changes, aiding in better change planning and management.

```sql
select
  avg(num_ci_affected) as avg_num_ci_affected 
from
  (
    select
      task_name,
      count(ci_item_sys_id) as num_ci_affected 
    from
      servicenow_sn_chg_rest_change_affected_cmdb_ci 
    group by
      task_name
  )
  as subquery;
```