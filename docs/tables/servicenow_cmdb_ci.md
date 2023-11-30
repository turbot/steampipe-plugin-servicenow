---
title: "Steampipe Table: servicenow_cmdb_ci - Query ServiceNow Configuration Items using SQL"
description: "Allows users to query Configuration Items (CI) in ServiceNow, specifically the details of each CI, providing insights into the configuration management database (CMDB)."
---

# Table: servicenow_cmdb_ci - Query ServiceNow Configuration Items using SQL

ServiceNow Configuration Items (CI) are service assets tracked in the Configuration Management Database (CMDB). They include tangible assets like hardware and software, as well as intangible assets like policies, agreements, and services. CIs provide information about the assets' configurations and the relationships between them.

## Table Usage Guide

The `servicenow_cmdb_ci` table provides insights into the Configuration Items (CI) within ServiceNow's Configuration Management Database (CMDB). As a DevOps engineer, you can explore CI-specific details through this table, including CI classes, states, and associated metadata. Utilize it to uncover information about CIs, such as those related to specific services, the relationships between CIs, and the verification of CI states.

## Examples

### List CMDB Configuration Item name, asset tag, serial number
Explore the inventory of Configuration Items (CIs) to keep track of their names, asset tags, and serial numbers. This query is particularly useful for auditing and asset management purposes.

```sql
select
  ci.name, ci.asset_tag, ci.serial_number
from
  servicenow_cmdb_ci ci
limit 100;
```

### Retrieve a list of all CIs that are currently down or experiencing issues
Explore which Configuration Items (CIs) are currently not functioning optimally or facing issues. This can aid in quickly identifying problematic areas and taking necessary corrective actions.

```sql
select
  sys_id,
  name,
  sys_class_name,
  short_description
from
  servicenow_cmdb_ci
where
  operational_status != '1';
```

### Identify all CIs that have been recently modified or updated
Explore which Configuration Items (CIs) have undergone changes in the past week. This is useful in understanding recent system changes and potentially identifying any unexpected or unauthorized modifications.

```sql
select
  sys_id,
  name,
  sys_class_name,
  short_description
from
  servicenow_cmdb_ci
where
  sys_updated_on > now() - interval '7 days';
```

### Get a list of all CIs that have been retired or decommissioned
Explore which Configuration Items (CIs) have been decommissioned or retired. This can be useful in assessing the lifecycle of your resources and planning for replacement or upgrades.

```sql
select
  sys_id,
  name,
  sys_class_name,
  short_description
from
  servicenow_cmdb_ci
where
  install_status = '7';
```

### Find all CIs that are currently undergoing maintenance or repair
Explore which Configuration Items (CIs) are currently under maintenance or repair. This is useful in managing system downtime and planning resource allocation.

```sql
select
  sys_id,
  name,
  sys_class_name,
  short_description
from
  servicenow_cmdb_ci
where
  sys_id in (
    select cmdb_ci
    from servicenow_task
    where
      state in (3,4) and
      short_description like '%maintenance%'
    );
```

### Retrieve all CIs with a specific serial number
Explore which Configuration Items (CIs) share a common serial number. This can be useful in tracking and managing assets within your IT environment.

```sql
select
  sys_id,
  name,
  sys_class_name,
  short_description
from
  servicenow_cmdb_ci
where
  serial_number = '123456789';
```