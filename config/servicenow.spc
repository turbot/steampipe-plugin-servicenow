connection "servicenow" {
  plugin = "servicenow"

  instance_url = "https://dev129225.service-now.com"
  client_id = "9148ce34f5252110392c96f819dbd422"
  client_secret = ""
  username = "admin"
  password = ""

  objects = ["sys_user", "sys_user_group", "incident", "customer_contact"]
}
