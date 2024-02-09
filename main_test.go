package main

import "testing"

func TestSpamMasker(t *testing.T) {

	testCases := []string{"Here's my spammy page: http://hehefouls.netHAHAHA see you.",
		"Here's my spammy page: http://hehefouls.net HAHAHA see you.",
		"Here's my spammy page: hTTp://youth-elixir.com",
		"Here's my spammy page: http://jjjjj",
		"http://elixir.com Here's my spammy page",
	}
	expectedCases := []string{"Here's my spammy page: http://******************* see you.",
		"Here's my spammy page: http://************* HAHAHA see you.",
		"Here's my spammy page: hTTp://youth-elixir.com",
		"Here's my spammy page: http://*****",
		"http://********** Here's my spammy page",
	}

	substring := "http://"

	for i, testcase := range testCases {

		t.Run(testcase, func(t *testing.T) {
			result := SpamMasker(testcase, substring)
			if result != expectedCases[i] {
				t.Errorf("incorrect result, expect %s, got %s", expectedCases[i], result)
			}
		})
	}
}
