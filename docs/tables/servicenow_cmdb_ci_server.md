---
title: "Steampipe Table: servicenow_cmdb_ci_server - Query ServiceNow CMDB CI Servers using SQL"
description: "Allows users to query Configuration Item (CI) Servers in ServiceNow's Configuration Management Database (CMDB), providing insights into server-specific details such as name, status, class, and other related attributes."
---

# Table: servicenow_cmdb_ci_server - Query ServiceNow CMDB CI Servers using SQL

ServiceNow's Configuration Management Database (CMDB) is a service that stores information about all technical services. The CMDB includes CI Servers, which represent logical entities or instances that are part of an IT infrastructure. These servers are used to store and manage data, run applications, and provide services.

## Table Usage Guide

The `servicenow_cmdb_ci_server` table provides insights into Configuration Item (CI) Servers within ServiceNow's Configuration Management Database (CMDB). As an IT administrator or manager, explore server-specific details through this table, including server names, statuses, classes, and other related attributes. Utilize it to manage and monitor your IT infrastructure effectively, ensuring optimal performance and availability of services.

## Examples

### List servers with their IP addresses and OS versions
Explore which servers are associated with specific IP addresses and operating systems. This is useful for maintaining an organized inventory of your servers and understanding their configurations.

```sql+postgres
select
  name,
  ip_address,
  os 
from
  servicenow_cmdb_ci_server
order by
  name;
```

```sql+sqlite
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
Gain insights into the distribution of operating systems across your servers. This query helps in understanding the spread of different OS versions, enabling strategic decisions for software compatibility and updates.

```sql+postgres
select
  os,
  count(*)
from
  servicenow_cmdb_ci_server 
group by
  os;
```

```sql+sqlite
select
  os,
  count(*)
from
  servicenow_cmdb_ci_server 
group by
  os;
```

### List all servers that have been in running for more than 3 years
Explore servers that have been operating for an extended period, specifically those running for over three years. This can aid in identifying potential maintenance needs or assessing the longevity and reliability of your server infrastructure.

```sql+postgres
select
  name,
  sys_updated_on 
from
  servicenow_cmdb_ci_server 
where
  sys_created_on <= now() - interval '3 years' and
  install_status = 1;
```

```sql+sqlite
select
  name,
  sys_updated_on 
from
  servicenow_cmdb_ci_server 
where
  sys_created_on <= datetime('now', '-3 years') and
  install_status = 1;
```

### List the servers with a specific serial number
Discover the servers that have a specific serial number. This can be beneficial for understanding the distribution and usage of a particular server model within your infrastructure.

```sql+postgres
select
  name,
  serial_number,
  ip_address
from
  servicenow_cmdb_ci_server
where
  serial_number = '478933-e78-8823';
```

```sql+sqlite
select
  name,
  serial_number,
  ip_address
from
  servicenow_cmdb_ci_server
where
  serial_number = '478933-e78-8823';
```