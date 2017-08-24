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
	"fmt"
	"log"
	"os"
)

// `main` performs the following tasks:
// - Load up the configuration file for beacon; if it doesn't exist, make it.
// - Load the beacon log, if it doesn't exist; make it.
// - Provide a switch statement for allowing different functionality based on command line arguments.
func main() {
	cfg, err := LoadConfig("./.beaconrc")
	if err != nil {
		log.Fatal(err)
	}

	// Load beacon_log and prepend newLog to the file
	beaconLog := LoadLog("./beacon_log.json")

	// Parse command line arguments to determine what Beacon action to take
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "add":
			beaconLog = append(beaconLog, NewLog(os.Args[2], cfg))
			SaveNewLog("./beacon_log.json", beaconLog)
			os.Exit(0)

		case "all":
			ReadLog(beaconLog, -1)
			os.Exit(0)

		case "show":
			CheckArgs([]string{"int"}, os.Args[2:], "show")

			// TODO: show variable amounts of reads
			ReadLog(beaconLog, 2)
			fmt.Println("")

		case "init":
			InitConfig()

		default:
			os.Exit(1)
		}
	}

}
