package main

import (
	"bytes"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Load(t *testing.T) {
	Convey("When loading a valid config file", t, func() {
		fs = mockFS{
			statErr: false,
			fileContents: `{
	"author": "John Smith",
	"email": "john.smith@example.com"
}`,
		}
		cfg, err := LoadConfigFile("valid_config.json")

		Convey("Should load author", func() {
			So(cfg.Author, ShouldEqual, "John Smith")
		})

		Convey("Should load email", func() {
			So(cfg.Email, ShouldEqual, "john.smith@example.com")
		})

		Convey("Should not return an error", func() {
			So(err, ShouldEqual, nil)
		})
	})

	Convey("When loading a config file with a missing author", t, func() {
		fs = mockFS{
			statErr: false,
			fileContents: `{
	"email": "john.smith@example.com"
}`,
		}
		_, err := LoadConfigFile("invalid_config_author.json")

		Convey("Should return an error", func() {
			So(err, ShouldEqual, ErrAuthorRequired)
		})
	})

	Convey("When loading a config file with a missing email", t, func() {
		fs = mockFS{
			statErr: false,
			fileContents: `{
	"author": "John Smith"
}`,
		}
		_, err := LoadConfigFile("invalid_config_email.json")

		Convey("Should return an error", func() {
			So(err, ShouldEqual, ErrEmailRequired)
		})
	})
}

func Test_Load_MissingFile(t *testing.T) {
	Convey("Should call InitConfig to create a beacon log", t, func() {
		out = &mockWriter{}
		var result []string

		runner = mockRunner{
			func(cmd string, args ...string) ([]byte, error) {
				buf := bytes.Buffer{}
				buf.WriteString(cmd)
				for _, v := range args {
					buf.WriteString(" ")
					buf.WriteString(v)
				}
				result = append(result, buf.String())

				return buf.Bytes(), nil
			},
		}

		fs = mockFS{
			true,
			"",
		}
		_, err := LoadConfigFile("fake_path.json")
		log.Println(err)
		So(len(result), ShouldEqual, 2)
	})
}

func Test_Load_EmptyPath(t *testing.T) {
	Convey("Should return ErrInvalidPath when no empty string used for config file", t, func() {
		_, err := LoadConfigFile("")

		So(err, ShouldEqual, ErrInvalidPath)
	})
}

func Test_LoadConfig_InvalidJSON(t *testing.T) {
	Convey("Should return an error when json is invalid", t, func() {
		_, err := LoadConfig([]byte("this isn't json!"))

		So(err, ShouldNotEqual, nil)
	})
}

func Test_InitConfig(t *testing.T) {

	Convey("When config file already exists", t, func() {
		fs = mockFS{}
		out = &mockWriter{}

		err := InitConfig("./test_config")

		Convey("Should return ErrConfigExists", func() {
			So(err, ShouldEqual, ErrConfigExists)
		})
	})

	Convey("When config file doesn't exist", t, func() {
		var result []string
		runner = mockRunner{
			func(cmd string, args ...string) ([]byte, error) {
				buf := bytes.Buffer{}
				buf.WriteString(cmd)
				for _, v := range args {
					buf.WriteString(" ")
					buf.WriteString(v)
				}
				result = append(result, buf.String())

				return buf.Bytes(), nil
			},
		}
		out = &mockWriter{}
		fs = mockFS{
			true,
			"",
		}
		err := InitConfig("./test_config")

		Convey("Should not return an error", func() {
			So(err, ShouldEqual, nil)
		})

		Convey("Should attempt to get git user name", func() {
			So(result[0], ShouldContainSubstring, "user.name")
		})

		Convey("Should attempt to get git user email", func() {
			So(result[1], ShouldContainSubstring, "user.email")
		})
	})
}

type mockRunner struct {
	toRun (func(string, ...string) ([]byte, error))
}

func (m mockRunner) Run(cmd string, args ...string) ([]byte, error) {
	return m.toRun(cmd, args...)
}
