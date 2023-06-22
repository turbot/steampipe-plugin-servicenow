---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/servicenow.svg"
brand_color: "#62D84E"
display_name: "ServiceNow"
name: "servicenow"
description: "Use SQL to query CMDB CI services, servers, incidents, objects and more from ServiceNow."
og_description: "Query ServiceNow with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/servicenow-social-graphic.png"
---

# ServiceNow + Steampipe

[ServiceNow](https://www.servicenow.com/) is a cloud-based platform that provides digital workflows to help businesses improve their operational efficiency and enhance customer and employee experiences.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  number,
  short_description,
  category,
  priority
from
  servicenow_incident
where
  state = 1;
```

```
+------------+---------------------------------------+----------+----------+
| number     | short_description                     | category | priority |
+------------+---------------------------------------+----------+----------+
| INC0000039 | Trouble getting to Oregon mail server | network  | 5        |
| INC0000046 | Can't access SFA software             | software | 3        |
| INC0009001 | Unable to post content on a Wiki page | inquiry  | 3        |
| INC0000057 | Performance problems with wifi        | inquiry  | 5        |
| INC0009005 | Email server is down.                 | software | 1        |
| INC0007002 | Need access to the common drive.      | inquiry  | 4        |
+------------+---------------------------------------+----------+----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/servicenow/tables)**

## Get started

### Install

Download and install the latest ServiceNow plugin:

```bash
steampipe plugin install servicenow
```

### Credentials

| Item        | Description                                                                                                                                                                          |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | ServiceNow API requires an instance URL, username, and password for all requests.                                                                                                    |
| Radius      | Each connection represents a single ServiceNow Installation.                                                                                                                         |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/servicenow.spc`)<br />2. Credentials specified in environment variables, e.g., `SERVICENOW_PASSWORD`. |

### Configuration

Installing the latest servicenow plugin will create a config file (~/.steampipe/config/servicenow.spc) with a single connection named `servicenow`:

```hcl
connection "servicenow" {
  plugin = "servicenow"

  # `instance_url` (required) - Your ServiceNow instance URL.
  # instance_url = "https://<your_servicenow_instance>.service-now.com"

  # `username` (required) - Your ServiceNow username.
  # username = "john.hill"
  # `password` (required) - Your ServiceNow password.
  # password = "j0t3-$j@H3"

  # Optionally, to use OAuth2 authentication mode, you'll need to have the `client_id` and `client_secret` of
  # a registered application in ServiceNow. You can register an application by going to
  # `All > System oAuth > Application Registry` and creating a new OAuth API endpoint for external clients.

  # `client_id` (optional) - ServiceNow login application client ID.
  # client_id = "9148ce34f5252110392c96f819dbd422"
  # `client_secret` (optional) - ServiceNow login application client secret.
  # client_secret = "Ly#2auTd"

  # `objects` (optional) - Additional ServiceNow tables you want to query.
  # objects = ["cmdb_model", "cmn_location"]
}
```

### Credentials from Environment Variables

The ServiceNow plugin will use environment variables to obtain credentials:

#### Basic Auth

```sh
export SERVICENOW_INSTANCE_URL=https://<your_servicenow_instance>.service-now.com
export SERVICENOW_USERNAME=john.hill
export SERVICENOW_PASSWORD=j0t3-$j@H3
```

#### OAuth2 - password mode

```sh
export SERVICENOW_INSTANCE_URL=https://<your_servicenow_instance>.service-now.com
export SERVICENOW_USERNAME=john.hill
export SERVICENOW_PASSWORD=j0t3-$j@H3
export SERVICENOW_CLIENT_ID=9148ce34f5252110392c96f819dbd422
export SERVICENOW_CLIENT_SECRET=Ly#2auTd
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-servicenow
- Community: [Slack Channel](https://steampipe.io/community/join)
