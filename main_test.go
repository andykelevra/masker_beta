package main

import "testing"

func TestSpamMasker(t *testing.T) {
	s := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	substring := "http://"
	expected := "Here's my spammy page: http://******************* see you."

	result := SpamMasker(s, substring)

	if result != expected {
		t.Errorf("incorrect result, expect %s, got %s", expected, result)
	}
}
