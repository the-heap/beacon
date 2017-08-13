package main

import "testing"

func TestString(t *testing.T) {
	type test struct {
		Log      Log
		expected string
	}

	testArray := []struct {
		Log      Log
		expected string
	}{
		{
			Log{
				Date:    1502569641,
				Email:   "j@jh.com",
				Author:  "John Henry",
				Message: "I broke lots of things!",
			},
			`
==========================================
Date: Aug 12 16:27:21
Author: John Henry (j@jh.com)
Message: I broke lots of things!
==========================================
`,
		},
	}

	for _, test := range testArray {
		got := test.Log.String()
		if got != test.expected {
			t.Errorf("Got doesn't equal expected result:\nGot:%v\nWant:%v\n", got, test.expected)
		}
	}

}
