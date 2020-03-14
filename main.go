package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var f interface{}
	dat, _ := ioutil.ReadFile("/workspaces/cleaner/tests/test_folders1.json")
	_ = json.Unmarshal(dat, &f)
	fmt.Println("Hello, World!")
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case bool:
			fmt.Println(k, "is bool", vv)
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
