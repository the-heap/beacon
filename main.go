package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var logFilePath = path.Join(os.Getenv("HOME"), "/Dropbox/The Heap/Beacon/beacon_log.json")

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile(path.Join(logFilePath))
	checkError(err)

	fmt.Print(string(dat))
}
