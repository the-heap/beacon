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
	"fmt"
	"log"
	"os"

	"github.com/the-heap/beacon/config"
	"github.com/the-heap/beacon/messagelog"
)

// ============================
// FUNCS
// ============================

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
	cfg, err := config.LoadFile("./.beaconrc")
	if err != nil {
		log.Fatal(err)
	}

	// parse command line arguments to determine action
	// TODO: Log date and Log unique ID
	// TODO: Switch from prompt to `beacon add "this is my breaking change message"`
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "add":
			// construct the new Log entry
			newLog := messagelog.Log{}
			newLog.Message = prompt("Enter message: ")
			newLog.Email = cfg.Email
			newLog.Author = cfg.Author

			// Load beacon_log and prepend newLog to the file
			beaconLogData := messagelog.LoadLog("./beacon_log.json")
			beaconLogData.Logs = append([]messagelog.Log{newLog}, beaconLogData.Logs...)
			beaconLogData.Save("./beacon_log.json")
			os.Exit(0)

		case "all":
			// Load and print all of the Beacon logs.
			beaconLogData := messagelog.LoadLog("./beacon_log.json")
			fmt.Println(beaconLogData)
			os.Exit(0)

		default:
			os.Exit(1)
		}
	}

}
