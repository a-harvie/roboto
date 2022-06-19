package main

import "fmt"

func Main(args map[string]interface{}) map[string]interface{} {
	fmt.Printf("received request with args: %#v\n", args)
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"
	return msg
}
