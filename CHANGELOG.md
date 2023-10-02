## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters.
- Recompiled plugin with Go version `1.21`.

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
