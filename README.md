![image](https://hub.steampipe.io/images/plugins/turbot/servicenow-social-graphic.png)

# ServiceNow Plugin for Steampipe

Use SQL to query sys_user, cmdb_ci, incident and more from ServiceNow.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/servicenow)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/servicenow/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-servicenow/issues)

## Quick start

### Install

Download and install the latest Steampipe plugin:

```bash
steampipe plugin install servicenow
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/servicenow#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/servicenow#configuration).

Configure your account details in `~/.steampipe/config/servicenow.spc`:

```hcl
connection "servicenow" {
  plugin = "servicenow"

  # Your ServiceNow instance URL
  instance_url = "https://<your_servicenow_instance>.service-now.com"

  # Authentication information
  username = "john.hill"
  password = "j0t3-$j@H3"
}
```

Or through environment variables:

```sh
export SERVICENOW_INSTANCE_URL=https://<your_servicenow_instance>.service-now.com
export SERVICENOW_USERNAME=john.hill
export SERVICENOW_PASSWORD=j0t3-$j@H3
```

Run steampipe:

```shell
steampipe query
```

List new incidents on your ServiceNow instance:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-servicenow.git
cd steampipe-plugin-servicenow
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/servicenow.spc
```

Try it!

```
steampipe query
> .inspect servicenow
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-servicenow/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [ServiceNow Plugin](https://github.com/turbot/steampipe-plugin-servicenow/labels/help%20wanted)
