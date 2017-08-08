/**
* Welcome to Beacon; an application for keeping teams in sync with breaking changes in their codebases.
* This application is built with Golang! The Maintainer (yours truly) has not written much Golang so
* Refactors, tips, and working things out together is greatly appreciated.
* Please refer to our github issues page for working on this project! https://github.com/the-heap/beacon/issues
*
*** ======== Useful links, Resources, etc: ============ ***
	* Join our Slack group: https://slackin-onxcmypksl.now.sh/ <-- might take bit of time load before you can register. ⏲
	* Learn about the Heap: http://theheap.us/page/about/ 🎒
	* Resources for contributing to open source with The Heap: https://theheap.us/page/resources/ 🎩
	* Golang tips: https://gobyexample.com/ 🐹
	* Podcast about Golang if you're interested: http://gotime.fm/ 🎧
*** =================================================== ***
*
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// ============================
// Types
// ============================

// BeaconLog represents the entire blob of json from the beacon file
type BeaconLog struct {
	Logs []Log
}

// Log is a single entry in the BeaconLog;
//
// _Example_: a user submits a breaking change to the database,
// in which they would fill out all the info in this struct.
type Log struct {
	ID      string
	Date    string
	Email   string
	Author  string
	Message string
}

// ============================
// Vars and Constants
// ============================

var logFilePath = path.Join(os.Getenv("HOME"), "/Dropbox/The Heap/Beacon/beacon_log.json")
var beaconLogData BeaconLog

// ============================
// Helpers
// ============================

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// ============================
// MAIN!
// ============================
func main() {

	// read JSON file from disk
	beaconLogFile, err := ioutil.ReadFile(path.Join(logFilePath))
	checkError(err)

	// unmarshal json and store it in the pointer to beaconLogData {?}
	// NOTE: figure out if you can use `checkError` here; don't yet understand golang's idiomatic errors handling.
	if err := json.Unmarshal(beaconLogFile, &beaconLogData); err != nil {
		panic(err)
	}

	// an example of accessing some unmarshalled json data.
	fmt.Println(beaconLogData.Logs[0].Message)
}
