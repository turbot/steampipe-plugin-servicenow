package main

import (
	"github.com/turbot/steampipe-plugin-servicenow/rest"
)

func main() {
	rest.New(rest.Config{
		URL:          "https://dev129225.service-now.com",
		GrantType:    "password",
		ClientID:     "9148ce34f5252110392c96f819dbd422",
		ClientSecret: "LasdsTd",
		Username:     "admin",
		Password:     "j0tdsadasH3",
	})
}
