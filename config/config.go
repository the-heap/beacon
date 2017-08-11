// Package config exists to load, parse and deliver a configuration struct
// to Beacon
package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jbydeley/beacon/errors"
)

const (
	// ErrInvalidPath is returned when the path provided is empty
	ErrInvalidPath = errors.Error("invalid path")

	// ErrAuthorRequired is returned when a config is loaded without an Author
	ErrAuthorRequired = errors.Error("author not found in config file")

	// ErrEmailRequired is returned when a config is loaded without an Email
	ErrEmailRequired = errors.Error("email not found in config file")
)

// Config holds all configuration for beacon to work
type Config struct {
	Author string `json:"author"`
	Email  string `json:"email"`
}

// LoadFile reads the file provided and returns a Config.
func LoadFile(path string) (*Config, error) {
	if path == "" {
		return nil, ErrInvalidPath
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap("error reading file", err)
	}

	return LoadConfig(body)
}

// LoadConfig will parse a config []byte and confirm that the config
// has the required fields
func LoadConfig(config []byte) (*Config, error) {
	cfg := &Config{}
	err := json.Unmarshal(config, &cfg)
	if err != nil {
		return nil, errors.Wrap("error parsing JSON", err)
	}

	if cfg.Author == "" {
		return nil, ErrAuthorRequired
	}

	if cfg.Email == "" {
		return nil, ErrEmailRequired
	}

	return cfg, nil
}
