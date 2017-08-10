/**
* Welcome to Beacon; an application for keeping teams in sync with breaking changes in their codebases.
* This application is built with Golang! The Maintainer (yours truly) has not written much Golang so
* Refactors, tips, and working things out together is greatly appreciated.
* Please refer to our github issues page for working on this project! https://github.com/the-heap/beacon/issues
*
*** ======== Useful links, Resources, etc: ============ ***
  * 🗯 Join our Slack group:              https://slackin-onxcmypksl.now.sh/
  * 🎒 Learn about the Heap:              http://theheap.us/page/about/
  * 🎩 New to Open source resoures:       https://theheap.us/page/resources/
  * 🐹 Golang tips:                       https://gobyexample.com/
  * 🎧 Cool Golang podcast:               http://gotime.fm/
*** =================================================== ***
*
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type beaconConfig struct {
	Author string `json:"author"`
	Email  string `json:"email"`
}

// ============================
// Vars and Constants
// ============================

var beaconLogData BeaconLog

// ============================
// FUNCS
// ============================

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// Print the beacon log to the terminal
func printLog(data BeaconLog) {
	for _, element := range data.Logs {
		fmt.Println("")
		fmt.Println("==========================================")
		fmt.Println("Date: ", element.Date)
		fmt.Println("Author: ", element.Author+" ("+element.Email+")") // I bet there's a nicer way to do this.
		fmt.Println("Message: ", element.Message)
		fmt.Println("==========================================")
	}
}

// loadConfig loads the .beaconrc file into the beaconConfig
func loadConfig(config *beaconConfig) error {
	configFile, err := ioutil.ReadFile("./.beaconrc")
	if err != nil {
		return err
	}

	return json.Unmarshal(configFile, &config)
}

// ============================
// MAIN!
// ============================
func main() {
	var config beaconConfig

	if err := loadConfig(&config); err != nil {
		log.Fatal(err)
	}

	// read JSON file from disk
	beaconLogFile, err := ioutil.ReadFile("./beacon_log.json")

	checkError(err)

	// unmarshal json and store it in the pointer to beaconLogData {?}
	// NOTE: figure out if you can use `checkError` here; don't yet understand golang's idiomatic errors handling.
	if err := json.Unmarshal(beaconLogFile, &beaconLogData); err != nil {
		panic(err)
	}

	// Print the beacon log!
	printLog(beaconLogData)
}
