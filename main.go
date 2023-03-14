package main

import (
	"fmt"

	"github.com/turbot/steampipe-plugin-servicenow/rest"
)

func main() {
	client, err := rest.New(rest.Config{
		InstanceURL:  "https://dev129225.service-now.com",
		GrantType:    "password",
		ClientID:     "9148ce34f5252110392c96f819dbd422",
		ClientSecret: "LasdsTd",
		Username:     "admin",
		Password:     "j0tdsadasH3",
	})

	fmt.Println(err)
	fmt.Println(client)

	cmdbCi, _ := client.GetCmdbCIs(10)
	fmt.Println(len(cmdbCi.Result))
	for _, ci := range cmdbCi.Result {
		if ci.ShortDescription != "" {
			fmt.Println(ci.ShortDescription)
		}
	}
}
