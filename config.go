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

// LoadConfigFile reads the file provided and returns a Config.
func LoadConfigFile(path string) (*Config, error) {
	if path == "" {
		return nil, ErrInvalidPath
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, Wrap("error reading file", err)
	}

	return LoadConfig(body)
}

// LoadConfig will parse a config []byte and confirm that the config
// has the required fields
func LoadConfig(config []byte) (*Config, error) {
	cfg := &Config{}
	err := json.Unmarshal(config, &cfg)
	if err != nil {
		return nil, Wrap("error parsing JSON", err)
	}

	if cfg.Author == "" {
		return nil, ErrAuthorRequired
	}

	if cfg.Email == "" {
		return nil, ErrEmailRequired
	}

	return cfg, nil
}

// InitConfig is used to setup Beacon on first run
// - It checks to see if there is a beaconrc file with the proper contents
// - ie. a user and email is configured
// - If there is no creds configured, try and retrieve them from git and Prompt ok useage
//   - if not ok, prompt for custom config mesages. (discuss - maybe we should just only accept git creds?)
// - if there is no beacon_log.json, prompt if they should create.
func InitConfig() {
	_, err := os.Stat("./.beacodlkanrc")
	// Beacon file Exists
	if err == nil {
		fmt.Println("You already have a beaconrc! You can update by hand if needed.")

		// beaconrc file does not exist
	} else if os.IsNotExist(err) {
		getGitUserName := exec.Command("sh", "-c", "git config --get user.name")
		gitUserName, err := getGitUserName.Output()
		if err != nil {
			fmt.Println("No Git user name found", err)
		}

		getGitEmail := exec.Command("sh", "-c", "git config --get user.email")
		gitEmail, err := getGitEmail.Output()
		if err != nil {
			fmt.Println("No Git email found", err)
		}

		// Get Git credentials
		fmt.Println("\nHi, I'm Beacon! 🚨 — I'm here to help you and your team keep in touch about breaking changes.")
		fmt.Println("I set up a `beaconrc` file in your directory root with your Git information. You can change this whenever you need!")
		fmt.Println("Thanks for using Beacon and have fun breaking stuff! 🔨")

		fmt.Print("User name: "+string(gitUserName), "Email: "+string(gitEmail), "\n")

	} else {
		fmt.Println("The stat call to your beaconrc failed", err)
	}
}
