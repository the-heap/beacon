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
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/the-heap/beacon/messagelog"
)

// ============================
// Types
// ============================

type beaconConfig struct {
	Author string `json:"author"`
	Email  string `json:"email"`
}

// ============================
// FUNCS
// ============================

func checkError(e error) {
	if e != nil {
		panic(e)
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

func prompt(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	message, _ := reader.ReadString('\n')

	return message
}

// ============================
// MAIN!
// ============================
func main() {
	var config beaconConfig

	if err := loadConfig(&config); err != nil {
		log.Fatal(err)
	}

	// parse command line arguments to determine action
	// TODO: Log date and Log unique ID
	// TODO: Switch from prompt to `beacon add "this is my breaking change message"`
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "add":
			log := messagelog.Log{}
			log.Message = prompt("Enter message: ")
			log.Author = config.Author
			fmt.Printf("%+v\n", log)
			os.Exit(0)
		case "all":
			beaconLogData := messagelog.LoadLog("./beacon_log.json")
			beaconLogData.PrintLog()
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}

}
