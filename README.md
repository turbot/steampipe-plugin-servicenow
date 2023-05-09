![image](https://hub.steampipe.io/images/plugins/turbot/servicenow-social-graphic.png)

# ServiceNow Plugin for Steampipe

Use SQL to query sys_user, cmdb_ci, incident and more from ServiceNow.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/servicenow)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/servicenow/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-servicenow/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install servicenow
```

Run a query:

```sql

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
