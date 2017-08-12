package messagelog

import "testing"

func TestString(t *testing.T) {
	type test struct {
		beaconLogData BeaconLog
		expected      string
	}

	testArray := []test{
		{
			beaconLogData: BeaconLog{
				Logs: []Log{
					Log{
						ID:      "1234abcd",
						Date:    "2017/07/19",
						Email:   "j@jh.com",
						Author:  "John Henry",
						Message: "I broke lots of things!",
					},
				},
			},
			expected: `
==========================================
Date: 2017/07/19
Author: John Henry (j@jh.com)
Message: I broke lots of things!
==========================================
`,
		},
	}

	for _, test := range testArray {
		got := test.beaconLogData.String()
		if got != test.expected {
			t.Errorf("Got doesn't equal expected result:\nGot:%v\nWant:%v\n", got, test.expected)
		}
	}

}
