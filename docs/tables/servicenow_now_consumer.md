---
title: "Steampipe Table: servicenow_now_consumer - Query ServiceNow Consumers using SQL"
description: "Allows users to query ServiceNow Consumers, providing insights into the consumer's details, including name, client ID, client secret, and redirect URL."
---

# Table: servicenow_now_consumer - Query ServiceNow Consumers using SQL

ServiceNow is a cloud-based platform designed to improve business efficiency and productivity. It supports various IT service management tasks and is typically used by businesses to handle incident reports or IT service requests. A ServiceNow Consumer is an entity that can make requests to and consume responses from an API.

## Table Usage Guide

The `servicenow_now_consumer` table provides insights into ServiceNow Consumers within the ServiceNow platform. As an IT manager or developer, explore consumer-specific details through this table, including the client ID, client secret, and redirect URL. Utilize it to manage and monitor your API consumers, ensuring they are correctly configured and operating as expected.

## Examples

### What are the first and last names and email of all active consumers?
Explore the active consumers by identifying their first and last names along with their email addresses. This is useful for maintaining up-to-date records or for reaching out to active consumers for feedback or promotional campaigns.

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
Determine the distribution of consumers across various states to understand regional demographics and market penetration. This query can help guide decisions on resource allocation, targeted marketing, and expansion strategies.

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
Discover the segments that consist of consumers without a listed mobile phone number. This can help in identifying potential gaps in your customer communication channels.

```sql
select
  count(*) 
from
  servicenow_now_consumer 
where
  mobile_phone is null;
```

### What is the most common prefix for consumer names?
Explore which prefix is most frequently used among consumer names, helping to identify common naming conventions and patterns. This could be particularly useful for data organization and customer segmentation strategies.

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
Explore how many consumers have registered with an email address that includes the word 'gmail'. This can be useful for understanding the popularity of different email service providers among your consumers.

```sql
select
  count(*) 
from
  servicenow_now_consumer 
where
  email like '%gmail%';
```

### Which consumers have a title that contains the word "Manager"?
Discover the segments that include individuals with managerial roles. This query can be used to identify those consumers who hold a title containing the term "Manager", which can be useful in tailoring communication or resources for this specific group.

```sql
select
  * 
from
  servicenow_now_consumer 
where
  title like '%Manager%';
```

### How many consumers have a photo attached to their record?
Discover the segments that have consumers with photos attached to their records. This is useful for understanding the extent of user engagement and personalization within your platform.

```sql
select
  count(*) 
from
  servicenow_now_consumer 
where
  photo is not null;
```