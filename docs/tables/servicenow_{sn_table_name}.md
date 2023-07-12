# Table: servicenow_{sn_table_name}

Query data from the table `servicenow_{sn_table_name}`, e.g, `servicenow_task`, `servicenow_u_my_custom_sn_table`.
A table is automatically created to represent each ServiceNow table in the `objects` argument of the [plugin configuration file](../../config/servicenow.spc).

For instance, given the following connection configuration:

```hcl
connection "servicenow" {
  plugin = "servicenow"
  instance_url = "https://my-instance.service-now.com"
  username = "my-user"
  password = "my-password"
  objects = ["task", "u_my_custom_sn_table", "cmdb_model", "cmn_location"]
```

This plugin will automatically create the tables `servicenow_task`, `servicenow_u_my_custom_sn_table`, `servicenow_cmdb_model`, and `servicenow_cmn_location`.

Please note that plugin initialization time will increase depending on the number of objects included in the `objects` argument.


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
```

List all tables with:

```sql
.inspect servicenow
+---------------------------------+--------------------------------+
| table                           | description                    |
+---------------------------------+--------------------------------+
| servicenow_cmdb_model           | Product Model.                 |
| servicenow_cmn_location         | Location.                      |
| servicenow_task                 | Task.                          |
| servicenow_u_my_custom_sn_table | My custom table on ServiceNow. |
| ...                             | ...                            |
+---------------------------------+--------------------------------+
```

### How many tasks were closed in the last 30 days?

```sql
select
  count(*) as closed_tasks 
from
  servicenow_task 
where
  state = '4' 
  and closed_at >= now() - interval '30 days';
```

### How many tasks have been created by each user?
  
```sql
SELECT sys_created_by, COUNT(*) AS num_tasks_created
FROM servicenow_task
GROUP BY sys_created_by
order BY sys_created_by;
```

### How many tasks have been opened in the last 24 hours?

```sql
SELECT COUNT(*) AS num_tasks
FROM servicenow_task
WHERE opened_at >= now() - interval '24 hours';
```
