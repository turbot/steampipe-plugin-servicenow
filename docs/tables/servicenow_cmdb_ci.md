# Table: servicenow_cmdb_ci

Configuration Item.

## Examples

### List CMDB Configuration Item name, asset tag, serial number

```sql
select
  ci.name, ci.asset_tag, ci.serial_number
from
  servicenow_cmdb_ci ci
limit 100;
```

### Retrieve a list of all CIs that are currently down or experiencing issues

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
