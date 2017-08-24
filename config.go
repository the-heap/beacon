package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	// ErrInvalidPath is returned when the path provided is empty
	ErrInvalidPath = Error("invalid path")

	// ErrAuthorRequired is returned when a config is loaded without an Author
	ErrAuthorRequired = Error("author not found in config file")

	// ErrEmailRequired is returned when a config is loaded without an Email
	ErrEmailRequired = Error("email not found in config file")
)

// Config holds all configuration for beacon to work
type Config struct {
	Author string `json:"author"`
	Email  string `json:"email"`
}

// LoadConfig reads the file provided and returns a Config (if no errs)
// If there is no config file this function will create one.
func LoadConfig(path string) (*Config, error) {
	configData := &Config{}

	if path == "" {
		return nil, ErrInvalidPath
	}

	// Check the beacon rc file exists; if not, create it.
	_, err := os.Stat("./.beaconrc")
	if err != nil {
		InitConfig()
	}

	// Read the config file from disk
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, Wrap("Error reading the `.beaconrc` file.", err)
	}

	if err := json.Unmarshal(configFile, &configData); err != nil {
		return nil, Wrap("error parsing JSON", err)
	}

	if configData.Author == "" {
		return nil, ErrAuthorRequired
	}

	if configData.Email == "" {
		return nil, ErrEmailRequired
	}

	return configData, nil

}

// InitConfig is used to setup Beacon on first run
// - Checks for the beacon rc file.
// - If there is no creds configured, try and retrieve them from git and implement
func InitConfig() {
	_, err := os.Stat("./.beaconrc")
	// Beacon file Exists
	if err == nil {
		fmt.Println("You already have a beaconrc! You can update by hand if needed.")

		// beaconrc file does not exist
	} else if os.IsNotExist(err) {
		// Get Git credentials
		gitUserName, err := exec.Command("sh", "-c", "git config --get user.name").Output()
		if err != nil {
			fmt.Println("No Git user name found", err)
		}

		gitEmail, err := exec.Command("sh", "-c", "git config --get user.email").Output()
		if err != nil {
			fmt.Println("No Git email found", err)
		}

		// Welcome message
		fmt.Println("\nHi, I'm Beacon! 🚨. I'm here to help you and your team keep in touch about breaking changes.")
		fmt.Println("I set up a `beaconrc` file in your directory root with your Git information. You can change this whenever you need!")
		fmt.Println("Thanks for using Beacon and have fun breaking stuff! 🔨")

		// create and write RC file
		file, err := os.Create("./.beaconrc")
		if err != nil {
		}

		json.NewEncoder(file).Encode(Config{ToStringCutNewLine(gitUserName), ToStringCutNewLine(gitEmail)})

	} else {
		fmt.Println("The stat call to your beaconrc failed", err)
	}
}

// InitBeaconLog creates a new `beacon_log.json` file in the directory in which beacon was invoked from
func InitBeaconLog() {
	// Check if the file exists!
	_, err := os.Stat("./beacon_log.json")
	if err != nil {
		fmt.Println("No Beacon Log found! Creating one now.")
		fmt.Println("You are good to go! 🔥")

		// Create the beacon file.
		file, err := os.Create("./beacon_log.json")
		if err != nil {
			fmt.Println("Failed to create beacon log")
		}

		file.WriteString("[]")
		file.Close()
	}
	os.Exit(1)
}
