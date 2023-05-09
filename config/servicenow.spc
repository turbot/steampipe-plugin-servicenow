connection "servicenow" {
  plugin = "servicenow"

  # `instance_url` (required) - Your ServiceNow instance URL.
  instance_url = "https://<your_servicenow_instance>.service-now.com"

  # `username` (required) - Your ServiceNow username.
  username = "john.hill"
  # `password` (required) - Your ServiceNow password.
  password = "j0t3-$j@H3"

  # Optionally, to use OAuth2 authentication mode, you'll need to have the `client_id` and `client_secret` of
  # a registered application in ServiceNow. You can register an application by going to
  # `All > System oAuth > Application Registry` and creating a new OAuth API endpoint for external clients.

  # `client_id` (optional) - ServiceNow login application client ID
  client_id = "9148ce34f5252110392c96f819dbd422"
  # `client_secret` (optional) - ServiceNow login application client secret
  client_secret = "Ly#)2auTd"

  # `objects` (required) - Additional ServiceNow tables you want to query.
  #objects = ["cmdb_model", "cmn_location"]
}
