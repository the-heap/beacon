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

// String implements the stringer interface for the BeaconLog struct
func (data BeaconLog) String() string {
	var stringOutput string
	for _, element := range data.Logs {
		stringOutput = stringOutput + fmt.Sprintf("\n==========================================\n")
		stringOutput = stringOutput + fmt.Sprintf("Date: %s\n", element.Date)
		stringOutput = stringOutput + fmt.Sprintf("Author: %s (%s)\n", element.Author, element.Email)
		stringOutput = stringOutput + fmt.Sprintf("Message: %s\n", element.Message)
		stringOutput = stringOutput + fmt.Sprintf("==========================================\n")
	}
	return stringOutput
}
