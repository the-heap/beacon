// Package messagelog provides the data type for holding the beacon log data
// and methods for loading and storing to/from the log file
package messagelog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

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

// LoadLog loads the beacon log from the specified file
func LoadLog(logFile string) BeaconLog {
	var beaconLogData BeaconLog

	// read JSON file from disk
	beaconLogFile, err := ioutil.ReadFile(logFile)

	if err != nil {
		log.Fatal(err)
	}

	// unmarshal json and store it in the pointer to beaconLogData {?}
	// NOTE: figure out if you can use `checkError` here; don't yet understand golang's idiomatic errors handling.
	if err := json.Unmarshal(beaconLogFile, &beaconLogData); err != nil {
		log.Fatal(err)
	}
	return beaconLogData
}

// PrintLog prints the beacon log to the terminal
func (data *BeaconLog) PrintLog() {
	for _, element := range data.Logs {
		fmt.Println("")
		fmt.Println("==========================================")
		fmt.Println("Date: ", element.Date)
		fmt.Println("Author: ", element.Author+" ("+element.Email+")") // I bet there's a nicer way to do this.
		fmt.Println("Message: ", element.Message)
		fmt.Println("==========================================")
	}
}
