---
title: "Steampipe Table: servicenow_now_contact - Query ServiceNow Contacts using SQL"
description: "Allows users to query ServiceNow Contacts, providing details about the individuals or groups that are associated with incidents, problems, and changes in the ServiceNow system."
---

# Table: servicenow_now_contact - Query ServiceNow Contacts using SQL

ServiceNow Contacts are individuals or groups associated with incidents, problems, and changes within the ServiceNow system. These contacts may include employees, customers, vendors, or any other entities that interact with the organization's ServiceNow system. They play a crucial role in incident management, problem management, and change management processes within ServiceNow.

## Table Usage Guide

The `servicenow_now_contact` table provides insights into the contacts within the ServiceNow system. As a system administrator or IT service manager, you can explore contact-specific details through this table, including their roles in incidents, problems, and changes. Utilize it to uncover information about contacts, such as their associations with incidents, problems, and changes, and the details of these associations.

## Examples

### How many contacts are in the servicenow_now_contact table?
Determine the total number of contacts recorded in your ServiceNow platform to assess the size of your network. This can be useful for understanding your reach and planning communication or marketing strategies.

```sql+postgres
select
  count(*) as num_contacts 
from
  servicenow_now_contact;
```

```sql+sqlite
select
  count(*) as num_contacts 
from
  servicenow_now_contact;
```

### What are the unique sources of the contacts?
Discover the segments that represent unique origins within your contact list, which can help streamline communication strategies and customize outreach efforts.

```sql+postgres
select distinct
  source 
from
  servicenow_now_contact;
```

```sql+sqlite
select distinct
  source 
from
  servicenow_now_contact;
```

### How many contacts were created in the last 7 days?
Discover the number of new contacts added within the last week. This can be useful for monitoring contact list growth and assessing recent engagement efforts.

```sql+postgres
select
  count(*) as num_contacts_last_7_days 
from
  servicenow_now_contact 
where
  sys_created_on >= now() - interval '7 days';
```

```sql+sqlite
select
  count(*) as num_contacts_last_7_days 
from
  servicenow_now_contact 
where
  sys_created_on >= datetime('now','-7 days');
```

### What is the most common country of origin for the contacts?
Explore which country is most frequently listed as the origin for the contacts. This could be useful for understanding the geographical distribution of your contacts and tailoring your services accordingly.

```sql+postgres
select
  country,
  count(*) as num_contacts 
from
  servicenow_now_contact 
group by
  country 
order by
  num_contacts desc limit 1;
```

```sql+sqlite
select
  country,
  count(*) as num_contacts 
from
  servicenow_now_contact 
group by
  country 
order by
  num_contacts desc limit 1;
```

### How many contacts were created by each source in the last 30 days?
Discover the segments that each source has contributed to your contacts list in the past month. This information is useful for understanding which sources are most effective in generating new contacts.

```sql+postgres
select
  source,
  count(*) as num_contacts_last_30_days 
from
  servicenow_now_contact 
where
  sys_created_on >= now() - interval '30 days' 
group by
  source;
```

```sql+sqlite
select
  source,
  count(*) as num_contacts_last_30_days 
from
  servicenow_now_contact 
where
  sys_created_on >= datetime('now','-30 days') 
group by
  source;
```

### What is the average number of contacts created per day in the last 90 days?
Analyze the trends in your contact creation by determining the average number of new contacts added each day over the past three months. This can help you understand your team's productivity and identify any significant changes in your contact volume.

```sql+postgres
select
  avg(num_contacts_per_day) as avg_contacts_per_day_last_90_days 
from
  (
    select
      count(*) as num_contacts_per_day 
    from
      servicenow_now_contact 
    where
      sys_created_on >= now() - interval '90 days' 
    group by
      sys_created_on 
  )
  as daily_contacts;
```

```sql+sqlite
select
  avg(num_contacts_per_day) as avg_contacts_per_day_last_90_days 
from
  (
    select
      count(*) as num_contacts_per_day 
    from
      servicenow_now_contact 
    where
      sys_created_on >= datetime('now', '-90 day') 
    group by
      sys_created_on 
  )
  as daily_contacts;
```