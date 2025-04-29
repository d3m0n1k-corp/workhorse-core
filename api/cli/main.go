package main

import (
	"encoding/json"
	"fmt"
	"workhorse-core/app"
)

// TODO : Move this to a test file, which will use a json file as input and check the output against a known value.
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
		},
		"app2": {
			"name": "workhorse",
			"version": "1.0.0",
			"email": "a@a.com"
		},
		"app3": {
			"app": {
			"name": "workhorse",
			"version": "1.0.0",
			"email": "a@a.com"
			},
			"app2": {
				"name": "workhorse",
				"version": "1.0.0",
				"email": "a@a.com"
			}
		}
	}
`

	conf := `{}`

	name := "json_stringify"

	result, err := app.ExecuteConverter(name, input, conf)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result)
}
