connection "servicenow" {
  plugin = "servicenow"

  # `instance_url` (required) - Your ServiceNow instance URL.
  instance_url = "https://dev129225.service-now.com"
  # `username` (required) - Your ServiceNow username.
  username = "admin"
  # `password` (required) - Your ServiceNow password.
  password = "j0?6:dA@H3"

  # Get your API credentials from ServiceNow by navigating to All > System oAuth > Application Registry
  # Then click New and create an OAuth API endpoint for external clients.
  # `client_id` (required) - Your application's client ID
  client_id = "9148ce34f5252110392c96f819dbd422"
  # `client_secret` (required) - Your application's client secret.
  client_secret = "L43-5*cauTd"

  # `objects` (required) - The ServiceNow tables you want to query.
  objects = ["sys_user", "sys_user_group", "incident"]
}
