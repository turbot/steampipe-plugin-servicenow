# Table: servicenow_incident

Incident.

## Examples

### What is the distribution of incidents by severity level?

```sql
select severity, count(*) as incident_count, count(*) * 100.0 / sum(count(*)) over() as percentage
from servicenow_incident
group by severity
order by severity;
```

### What is the average time taken to resolve incidents by category?

```sql
select category, avg(calendar_stc) as avg_time_to_resolve
from servicenow_incident
where incident_state = 7 -- only include resolved incidents
group by category;
```

### What was the most common reason for holding incidents?

```sql
select hold_reason, count(*) as incident_count
from servicenow_incident
where hold_reason is not null
group by hold_reason
order by incident_count desc
limit 1;
```

### What is the average time taken to resolve incidents by severity level?

```sql
select severity, avg(calendar_stc) as avg_time_to_resolve
from servicenow_incident
where incident_state = 7 -- only include resolved incidents
group by severity;
```

### What are the top 10 categories of incidents by count?

```sql
select category, count(*) as incident_count
from servicenow_incident
group by category
order by incident_count desc
limit 10;
```

### How many incidents were resolved in the last 30 days?

```sql
select count(*) as resolved_incidents
from servicenow_incident
where incident_state = '6' and resolved_at >= now() - interval '30 days';
```

### How many incidents were reopened more than once?

```sql
select count(*) as reopened_incidents
from servicenow_incident
where reopen_count > 1;
```
