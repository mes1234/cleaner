package main

import (
	"encoding/json"
	"fileshandler"
	"fmt"
	"io/ioutil"
)

func main() {
	var f interface{}
	dat, _ := ioutil.ReadFile("/workspaces/cleaner/tests/test_folders1.json")
	_ = json.Unmarshal(dat, &f)
	fmt.Println("Hello, World!")
	m := f.(map[string]interface{})
	fileshandler.Discover(m)

}
