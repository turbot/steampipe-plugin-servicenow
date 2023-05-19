# Table: servicenow_now_contact

Customer Service Management (CSM) contact records.

## Examples

### How many contacts are in the `servicenow_now_contact` table?

```sql
select
  count(*) as num_contacts 
from
  servicenow_now_contact;
```

### What are the unique sources of the contacts?

```sql
select distinct
  source 
from
  servicenow_now_contact;
```

### How many contacts were created in the last 7 days?

```sql
select
  count(*) as num_contacts_last_7_days 
from
  servicenow_now_contact 
where
  sys_created_on >= now() - interval '7 days'
```

### What is the most common country of origin for the contacts?

```sql
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

```sql
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

### What is the average number of contacts created per day in the last 90 days?

```sql
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
