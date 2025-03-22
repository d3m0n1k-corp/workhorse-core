package main

import (
	"encoding/json"
	"fmt"
	"workhorse-core/app"
)

func main() {
	res := app.ListConverters()
	res_json, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(res_json))
	input := `
	{
		"app": {
			"name": "workhorse",
			"version": "1.0.0",
			"email": "a@a.com"
		}
	}
`

	conf := `{}`

	name := "json_to_yaml"

	result, err := app.ExecuteConverter(name, input, conf)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result)
}
