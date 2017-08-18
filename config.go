package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	// ErrInvalidPath is returned when the path provided is empty
	ErrInvalidPath = Error("invalid path")

	// ErrAuthorRequired is returned when a config is loaded without an Author
	ErrAuthorRequired = Error("author not found in config file")

	// ErrEmailRequired is returned when a config is loaded without an Email
	ErrEmailRequired = Error("email not found in config file")

	// ErrBeaconLogExists is returned when trying to initialize a beacon log when one exists
	ErrConfigExists = Error(".beaconrc already exists")
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

	_, err := fs.Stat(path)
	if err != nil {
		InitConfig(path)
	}

	body, err := fs.ReadFile(path)
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
// - Checks for the beacon rc file.
// - If there is no creds configured, try and retrieve them from git and implement
func InitConfig(path string) error {
	_, err := fs.Stat(path)
	// Beacon file Exists
	if err == nil {
		return ErrConfigExists

		// beaconrc file does not exist
	} else if os.IsNotExist(err) {
		// Get Git credentials
		gitUserName, err := runner.Run("sh", "-c", "git config --get user.name")
		if err != nil {
			log.Println("No Git user name found", err)
		}

		gitEmail, err := runner.Run("sh", "-c", "git config --get user.email")
		if err != nil {
			log.Println("No Git email found", err)
		}

		// Welcome message
		fmt.Fprintln(out, "\nHi, I'm Beacon! 🚨 — I'm here to help you and your team keep in touch about breaking changes.")
		fmt.Fprintln(out, "I set up a `beaconrc` file in your directory root with your Git information. You can change this whenever you need!")
		fmt.Fprintln(out, "Thanks for using Beacon and have fun breaking stuff! 🔨")

		// create and write RC file
		file, err := fs.Create(path)
		if err != nil {
			return err
		}

		return json.NewEncoder(file).Encode(Config{ToStringCutNewLine(gitUserName), ToStringCutNewLine(gitEmail)})
	}

	return err
}
