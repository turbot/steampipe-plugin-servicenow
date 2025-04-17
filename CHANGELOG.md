## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#33](https://github.com/turbot/steampipe-plugin-servicenow/pull/33))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#33](https://github.com/turbot/steampipe-plugin-servicenow/pull/33))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#29](https://github.com/turbot/steampipe-plugin-servicenow/pull/29))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#29](https://github.com/turbot/steampipe-plugin-servicenow/pull/29))

## v0.3.1 [2024-01-08]

_Bug fixes_

- Fixed the plugin to return only static tables instead of an error when the `objects` config argument is not set or credentials are not set correctly. ([#26](https://github.com/turbot/steampipe-plugin-servicenow/pull/26))

## v0.3.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#22](https://github.com/turbot/steampipe-plugin-servicenow/pull/22))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#22](https://github.com/turbot/steampipe-plugin-servicenow/pull/22))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-servicenow/blob/main/docs/LICENSE). ([#22](https://github.com/turbot/steampipe-plugin-servicenow/pull/22))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#21](https://github.com/turbot/steampipe-plugin-servicenow/pull/21))

## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#14](https://github.com/turbot/steampipe-plugin-servicenow/pull/14))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#11](https://github.com/turbot/steampipe-plugin-servicenow/pull/11))
- Recompiled plugin with Go version `1.21`. ([#11](https://github.com/turbot/steampipe-plugin-servicenow/pull/11))

## v0.1.1 [2023-07-24]

_Bug fixes_

- Fixed the formatting of the output of example queries in dynamic table doc. ([#6](https://github.com/turbot/steampipe-plugin-servicenow/pull/6))

## v0.1.0 [2023-07-14]

_Enhancements_

- Added documentation for servicenow dynamic tables. ([#4](https://github.com/turbot/steampipe-plugin-servicenow/pull/4))

## v0.0.1 [2023-06-22]

_What's new?_

- New tables added
  - [servicenow_cmdb_ci](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_cmdb_ci)
  - [servicenow_cmdb_ci_server](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_cmdb_ci_server)
  - [servicenow_cmdb_ci_service](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_cmdb_ci_service)
  - [servicenow_incident](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_incident)
  - [servicenow_now_consumer](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_now_consumer)
  - [servicenow_now_contact](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_now_contact)
  - [servicenow_object](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_object)
  - [servicenow_sn_chg_rest_change](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sn_chg_rest_change)
  - [servicenow_sn_chg_rest_change_affected_cmdb_ci](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sn_chg_rest_change_affected_cmdb_ci)
  - [servicenow_sn_chg_rest_change_impacted_cmdb_ci_service](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sn_chg_rest_change_impacted_cmdb_ci_service)
  - [servicenow_sn_chg_rest_change_model](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sn_chg_rest_change_model)
  - [servicenow_sn_chg_rest_change_task](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sn_chg_rest_change_task)
  - [servicenow_sys_audit](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_audit)
  - [servicenow_sys_audit_relation](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_audit_relation)
  - [servicenow_sys_group_has_role](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_group_has_role)
  - [servicenow_sys_user](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_user)
  - [servicenow_sys_user_grmember](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_user_grmember)
  - [servicenow_sys_user_group](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_user_group)
  - [servicenow_sys_user_has_role](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_user_has_role)
  - [servicenow_sys_user_role](https://hub.steampipe.io/plugins/turbot/servicenow/tables/servicenow_sys_user_role)
