# Table: servicenow_cmdb_ci_server

Configuration Item Server.

## Examples

### List servers with their IP addresses and OS versions

```sql
select
  name,
  ip_address,
  os 
from
  servicenow_cmdb_ci_server
order by
  name;
```

### Number of servers running each OS version

```sql
select
  os,
  count(*)
from
  servicenow_cmdb_ci_server 
group by
  os;
```

### List all servers that have been in running for more than 3 years

```sql
select
  name,
  sys_updated_on 
from
  servicenow_cmdb_ci_server 
where
  sys_created_on <= now() - interval '3 years' and
  install_status = 1;
```

### List the servers with a specific serial number

```sql
select
  name,
  serial_number,
  ip_address
from
  servicenow_cmdb_ci_server
where
  serial_number = '478933-e78-8823';
```
