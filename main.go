package main

import (
	"fmt"

	"github.com/turbot/steampipe-plugin-servicenow/rest"
)

func main() {
	resp := &rest.OAuthTokenResponse{}
	code, err := rest.New(rest.Config{
		InstanceURL:  "https://dev129225.service-now.com",
		GrantType:    "password",
		ClientID:     "9148ce34f5252110392c96f819dbd422",
		ClientSecret: "LasdsTd",
		Username:     "admin",
		Password:     "j0tdsadasH3",
	})
	}, resp)

	fmt.Println(code)
	fmt.Println(err)
	fmt.Println(resp.AccessToken)
}
