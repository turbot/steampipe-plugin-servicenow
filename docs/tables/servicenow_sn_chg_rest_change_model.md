# Table: servicenow_sn_chg_rest_change_model

Change Management - Change Model.

## Examples

### Get the name and description of all change models, ordered by name

```sql
select
  name,
  description 
from
  servicenow_sn_chg_rest_change_model 
order by
  name;
```

### Get all the active change models available in the UI

```sql
select
  * 
from
  servicenow_sn_chg_rest_change_model 
where
  active = true 
  and available_in_ui = true;
```

### Get the default change model

```sql
select
  * 
from
  servicenow_sn_chg_rest_change_model 
where
  default_change_model = true;
```

### Get the number of times each change model has been modified

```sql
select
  name,
  sys_mod_count 
from
  servicenow_sn_chg_rest_change_model 
order by
  sys_mod_count desc;
```
