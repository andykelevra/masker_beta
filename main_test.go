package main

import "testing"

type Case struct {
	name           string
	testString     string
	expectedString string
}

func TestSpamMasker(t *testing.T) {

	testCases := []Case{
		{"test1",
			"Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			"Here's my spammy page: http://******************* see you.",
		},
		{"test2",
			"Here's my spammy page: http://hehefouls.net HAHAHA see you.",
			"Here's my spammy page: http://************* HAHAHA see you.",
		},
		{"test3",
			"Here's my spammy page: hTTp://youth-elixir.com",
			"Here's my spammy page: hTTp://youth-elixir.com",
		},
		{"test4",
			"Here's my spammy page: http://jjjjj",
			"Here's my spammy page: http://*****",
		},
		{"test5",
			"http://elixir.com Here's my spammy page",
			"http://********** Here's my spammy page",
		},
		{"test6",
			"Here's my spammy page: http://",
			"Here's my spammy page: http://",
		},
	}

	substring := "http://"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SpamMasker(tc.testString, substring)
			if result != tc.expectedString {
				t.Errorf("%s incorrect result, expect %s, got %s", tc.name, tc.expectedString, result)
			}
		})
	}
}
