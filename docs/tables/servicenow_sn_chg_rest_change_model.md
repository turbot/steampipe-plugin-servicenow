---
title: "Steampipe Table: servicenow_sn_chg_rest_change_model - Query ServiceNow Change Models using SQL"
description: "Allows users to query Change Models in ServiceNow, specifically the details of change models including their unique identifiers, names, and descriptions."
---

# Table: servicenow_sn_chg_rest_change_model - Query ServiceNow Change Models using SQL

ServiceNow Change Models provide a pre-defined and standardized method to handle changes in the IT infrastructure. They are used to manage and track changes in a systematic and repeatable way. Change Models contain detailed steps and procedures that should be followed while implementing a change.

## Table Usage Guide

The `servicenow_sn_chg_rest_change_model` table provides insights into Change Models within ServiceNow. As a Change Manager, you can explore model-specific details through this table, including model identifiers, names, and descriptions. Use this table to understand the various change models available, their purposes, and steps involved in their implementation.

## Examples

### Get the name and description of all change models, ordered by name
Explore the various change models in your system, sorted by their names. This can be useful in understanding the different types of changes that can occur and planning accordingly.

```sql+postgres
select
  name,
  description 
from
  servicenow_sn_chg_rest_change_model 
order by
  name;
```

```sql+sqlite
select
  name,
  description 
from
  servicenow_sn_chg_rest_change_model 
order by
  name;
```

### Get all the active change models available in the UI
Explore which change models are currently active and available for use within the user interface. This allows you to understand which models are accessible for implementing changes, aiding in efficient change management.

```sql+postgres
select
  * 
from
  servicenow_sn_chg_rest_change_model 
where
  active = true 
  and available_in_ui = true;
```

```sql+sqlite
select
  * 
from
  servicenow_sn_chg_rest_change_model 
where
  active = 1 
  and available_in_ui = 1;
```

### Get the default change model
Explore which change models are set as default in your ServiceNow configuration. This can be useful in understanding and managing your change process effectively.

```sql+postgres
select
  * 
from
  servicenow_sn_chg_rest_change_model 
where
  default_change_model = true;
```

```sql+sqlite
select
  * 
from
  servicenow_sn_chg_rest_change_model 
where
  default_change_model = 1;
```

### Get the number of times each change model has been modified
Analyze the frequency of modifications to each change model to understand the most frequently updated ones. This can help identify areas of frequent change, potentially highlighting areas for process improvement.

```sql+postgres
select
  name,
  sys_mod_count 
from
  servicenow_sn_chg_rest_change_model 
order by
  sys_mod_count desc;
```

```sql+sqlite
select
  name,
  sys_mod_count 
from
  servicenow_sn_chg_rest_change_model 
order by
  sys_mod_count desc;
```