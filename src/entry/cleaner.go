package main

import (
	"encoding/json"
	"fileshandler"
	"fmt"
	"io/ioutil"
)

func main() {
	var fr fileshandler.FileRaw
	dat, _ := ioutil.ReadFile("/workspaces/cleaner/tests/test_folders1.json")
	_ = json.Unmarshal(dat, &fr)
	f := *fr.Discover()
	fd := fileshandler.Directories(f)
	dirs, files := fd.Split()
	fmt.Println(dirs)
	fmt.Println(files)

}
