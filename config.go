package main

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"io/ioutil"
	"os"
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
// - if there is no beacon_log.json, create it.
func InitConfig() {
	_, err := os.Stat("./beaconrdac")
	if err == nil {
		fmt.Println("You already have a beaconrc")
		// TODO: Instruct user they can update it...
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist!")
		// TODO: Prompt user to make one, add credentials etc.
	} else {
		fmt.Println("The stat call to your beaconrc failed", err)
	}
}
