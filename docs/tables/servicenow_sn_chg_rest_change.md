# Table: servicenow_sn_chg_rest_change

Change Management - Change request.

## Examples

### What are the most common category for change requests?

```sql
select
  category,
  count(*) as cnt 
from
  servicenow_sn_chg_rest_change 
group by
  category 
order by
  cnt desc limit 5;
```

### How many change requests are currently open for a specific category?

```sql
select
  count(*) 
from
  servicenow_sn_chg_rest_change 
where
  state != 3 
  and category = 'network';
```

### How many changes are in the "work in progress" state?

```sql
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 2;
```

### What is the average duration of changes made by each user?

```sql
select
  user,
  avg(duration) as avg_duration 
from
  servicenow_sn_chg_rest_change 
group by
  user;
```

### How many changes are in the "scheduled" state and have a "high" priority?

```sql
select
  count(*) as num_changes 
from
  servicenow_sn_chg_rest_change 
where
  state = 1 
  and priority = 2;
```
