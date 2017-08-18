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
	"io"
	"log"
	"os"
)

var (
	fs     fileSystem
	runner cmdRunner
	rc     clock
	out    io.Writer
)

func main() {
	fs = osFS{}
	runner = realRunner{}
	rc = realClock{}
	out = os.Stdout

	cfg, err := LoadConfigFile("./.beaconrc")
	if err != nil {
		log.Fatal(err)
	}

	// Load beacon_log and prepend newLog to the file
	logs := LoadBeaconLog("./beacon_log.json")

	// parse command line arguments to determine action
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "add":
			logs = append(logs, New(os.Args[2], cfg))
			SaveNewLog("./beacon_log.json", logs)
			os.Exit(0)

		case "all":
			ShowLog(os.Stdout, logs, -1)
			os.Exit(0)

		case "show":
			CheckArgs([]string{"int"}, os.Args[2:], "show")

			ShowLog(os.Stdout, logs, 2)
			fmt.Println("")

		case "init":
			//InitConfig()

		default:
			os.Exit(1)
		}
	}

}
