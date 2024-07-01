package main

import (
	"fmt"
)

// Hello receive an event in format
// { "name": "Mike"}
// and returns a greeting in format
// { "body": "Hello, Mike"}
func Hello(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	greetings := "world"
	name, ok := args["name"].(string)
	if ok {
		greetings = name
	}
	res["body"] = "Hello, " + greetings
	fmt.Printf("Hello, %s\n", greetings)
	return res
}