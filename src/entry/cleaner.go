package main

import (
	"encoding/json"
	"fileshandler"
	"fmt"
	"io/ioutil"
)

func main() {
	var rd []fileshandler.RawStruct
	dat, _ := ioutil.ReadFile("/workspaces/cleaner/tests/test_folders1.json")
	_ = json.Unmarshal(dat, &rd)
	// f := fr.Discover()
	fmt.Println("Bye")

}
