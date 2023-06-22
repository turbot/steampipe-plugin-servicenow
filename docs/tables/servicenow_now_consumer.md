# Table: servicenow_now_consumer

Customer Service Management (CSM) consumer records.

## Examples

### What are the first and last names and email of all active consumers?

```sql
select
  first_name,
  last_name,
  email 
from
  servicenow_now_consumer 
where
  active = true;
```

### How many consumers are there in each state?

```sql
select
  state,
  count(*) as num_consumers 
from
  servicenow_now_consumer 
group by
  state 
order by
  state desc;
```

### What is the total number of consumers who doesn't have a mobile phone number listed?

```sql
select
  count(*) 
from
  servicenow_now_consumer 
where
  mobile_phone is null;
```

### What is the most common prefix for consumer names?

```sql
select
  prefix,
  count(*) as num_consumers 
from
  servicenow_now_consumer 
group by
  prefix 
order by
  num_consumers desc limit 1;
```

### How many consumers have an email address that contains the word "gmail"?

```sql
select
  count(*) 
from
  servicenow_now_consumer 
where
  email like '%gmail%';
```

### Which consumers have a title that contains the word "Manager"?

```sql
select
  * 
from
  servicenow_now_consumer 
where
  title like '%Manager%';
```

### How many consumers have a photo attached to their record?

```sql
select
  count(*) 
from
  servicenow_now_consumer 
where
  photo is not null;
```
