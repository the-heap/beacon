package main

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	TestAuthorName  = "John Smith"
	TestAuthorEmail = "john.smith@example.com"
)

func Test_New(t *testing.T) {
	Convey("When creating a log with New", t, func() {
		rc = mockClock{}

		cfg := &Config{
			Author: TestAuthorName,
			Email:  TestAuthorEmail,
		}

		Convey("Should return a new log", func() {
			log := New("New Log", cfg)

			So(log.Author, ShouldEqual, TestAuthorName)
		})
	})
}

func Test_SaveNewLog(t *testing.T) {
	Convey("When saving a log file", t, func() {
		err := SaveNewLog("beacon_test.json", []Log{})
		So(err, ShouldEqual, nil)
	})
}

func Test_ShowLog(t *testing.T) {
	rc = mockClock{}

	cfg := &Config{
		Author: TestAuthorName,
		Email:  TestAuthorEmail,
	}
	Convey("When ShowLog is called", t, func() {
		logs := []Log{
			New("First", cfg),
			New("Second", cfg),
		}

		Convey("With -1 for the count", func() {
			w := &mockWriter{}

			Convey("Should return all logs", func() {
				ShowLog(w, logs, -1)
				So(w.String(), ShouldEqual, `
==========================================
Date: Aug 16 21:02:03
Author: John Smith (john.smith@example.com)
Message: Second
==========================================


==========================================
Date: Aug 16 21:02:03
Author: John Smith (john.smith@example.com)
Message: First
==========================================

`)
			})
		})

		Convey("With a number greater then the total logs", func() {
			Convey("Should return all logs", func() {
				w := &mockWriter{}
				ShowLog(w, logs, 5)
				So(w.String(), ShouldEqual, `
==========================================
Date: Aug 16 21:02:03
Author: John Smith (john.smith@example.com)
Message: Second
==========================================


==========================================
Date: Aug 16 21:02:03
Author: John Smith (john.smith@example.com)
Message: First
==========================================

`)
			})
		})

		Convey("With a number greater then 0 and less then the total logs", func() {
			Convey("Should return that number of logs", func() {
				w := &mockWriter{}
				ShowLog(w, logs, 1)
				So(w.String(), ShouldEqual, `
==========================================
Date: Aug 16 21:02:03
Author: John Smith (john.smith@example.com)
Message: First
==========================================

`)
			})
		})
	})
}
func TestString(t *testing.T) {
	Convey("String should format the Log", t, func() {
		log := Log{
			Date:    1502569641,
			Email:   TestAuthorEmail,
			Author:  TestAuthorName,
			Message: "I broke lots of things!",
		}

		expected := `
==========================================
Date: Aug 12 16:27:21
Author: John Smith (john.smith@example.com)
Message: I broke lots of things!
==========================================
`

		actual := log.String()

		So(actual, ShouldEqual, expected)
	})
}

func Test_LoadBeaconLog(t *testing.T) {
	Convey("When loading a beacon log", t, func() {
		Convey("When the log doesn't exist", func() {
			fs = mockFS{
				statErr:      true,
				fileContents: "[]",
			}
			logs := LoadBeaconLog("fake_file.json")
			So(len(logs), ShouldEqual, 0)
		})
	})
}
func Test_InitBeaconLog(t *testing.T) {

	Convey("When initializing a beacon log", t, func() {

		Convey("Should return ErrBeaconLogExists if beacon_log.json exists", func() {
			fs = mockFS{}

			err := InitBeaconLog("beacon_test.json")
			So(err, ShouldEqual, ErrBeaconLogExists)
		})

		Convey("Should create beacon_log.json if it doesn't exist", func() {
			fs = mockFS{
				true,
				"",
			}
			err := InitBeaconLog("beacon_test.json")

			Convey("Should have no errors", func() {
				So(err, ShouldEqual, nil)
			})
		})
	})
}

type mockWriter struct {
	data bytes.Buffer
}

func (m mockWriter) String() string {
	return m.data.String()
}
func (m *mockWriter) Write(data []byte) (int, error) {
	m.data.Write(data)
	return 0, nil
}

type mockFS struct {
	statErr      bool
	fileContents string
}

func (mockFS) Open(name string) (file, error) { return os.Open(name) }
func (m mockFS) Stat(name string) (os.FileInfo, error) {
	if m.statErr {
		return nil, os.ErrNotExist
	}
	return nil, nil
}

func (mockFS) Create(name string) (file, error) {
	return mockFile{}, nil
}

func (m mockFS) ReadFile(name string) ([]byte, error) {
	return []byte(m.fileContents), nil
}

type mockFile struct {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
}

func (mockFile) Stat() (os.FileInfo, error) {
	return os.Stat("fake file")
}

func (mockFile) WriteString(data string) (int, error) {
	return 0, nil
}

func (mockFile) Close() error {
	return nil
}

func (mockFile) Write(data []byte) (int, error) {
	return len(data), nil
}

type mockClock struct{}

func (mockClock) Now() time.Time {
	return time.Date(2017, time.August, 17, 1, 2, 3, 4, time.UTC)
}
