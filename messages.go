package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

// Log is a single entry in the Log File that represents a breaking change;
//
// _Example_: a user submits a breaking change to the database,
// in which they would fill out all the info in this struct.
type Log struct {
	Date    int64
	Email   string
	Author  string
	Message string
}

// New instantiates a log
func New(msg string, cfg *Config) Log {
	return Log{
		Author:  cfg.Author,
		Email:   cfg.Email,
		Message: msg,
		Date:    time.Now().Unix(),
	}
}

// LoadBeaconLog loads the beacon log from the specified file
func LoadBeaconLog(logFileName string) []Log {
	var logs []Log

	// Check if file exists
	_, err := os.Stat("./beacon_log.json")
	if err != nil {
		InitBeaconLog()
	}

	// read JSON file from disk
	logFile, err := ioutil.ReadFile(logFileName)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(logFile, &logs); err != nil {
		log.Fatal(err)
	}

	// Make sure that we are FILO sorted
	sort.Sort(ByDate(logs))

	return logs
}

// SaveNewLog will persist the log to the file
func SaveNewLog(logFile string, data []Log) error {
	file, err := os.Create(logFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	return json.NewEncoder(file).Encode(data)
}

// ShowLog prints the number of beacon entries requested
func ShowLog(logs []Log, count int) {
	if count > len(logs) || count < 0 {
		count = len(logs)
	}

	for i := count; i > 0; i-- {
		fmt.Println(logs[i-1])
	}
}

// String implements the stringer interface for the BeaconLog struct
func (l Log) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "\n==========================================\n")
	fmt.Fprintf(buf, "Date: %s\n", time.Unix(l.Date, 0).Format(time.Stamp))
	fmt.Fprintf(buf, "Author: %s (%s)\n", l.Author, l.Email)
	fmt.Fprintf(buf, "Message: %s\n", l.Message)
	fmt.Fprintf(buf, "==========================================\n")
	return buf.String()
}
