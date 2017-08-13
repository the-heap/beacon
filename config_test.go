package main

import (
	"path"
	"testing"
)

func Test_Load(t *testing.T) {
	tt := []struct {
		Path string
		Err  error
	}{
		{
			"valid_config.json",
			nil,
		}, {
			"valid_config_extras.json",
			nil,
		}, {
			"invalid_config_author.json",
			ErrAuthorRequired,
		}, {
			"invalid_config_email.json",
			ErrEmailRequired,
		},
	}

	for _, tc := range tt {
		path := path.Join("testdata", tc.Path)
		_, err := LoadConfigFile(path)
		if err != tc.Err {
			t.Fatalf("Load(%v) returned an error (%v)", path, err)
		}
	}
}

func Test_Load_MissingFile(t *testing.T) {
	path := path.Join("testdata", "no_exist.json")
	_, err := LoadConfigFile(path)
	if err == nil {
		t.Fatalf("Load(%v) should have returned an error but did not", path)
	}
}

func Test_Load_EmptyPath(t *testing.T) {
	_, err := LoadConfigFile("")
	if err != ErrInvalidPath {
		t.Fatalf("Load(%v) should have return returned (%v)", "", ErrInvalidPath)
	}
}

func Test_LoadConfig_InvalidJSON(t *testing.T) {
	_, err := LoadConfig([]byte("this isn't json!"))
	if err == nil {
		t.Fatal("LoadConfig() return no error when passed invalid JSON")
	}
}
