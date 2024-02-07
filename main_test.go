package main

import "testing"

func TestSpamMasker(t *testing.T) {

	t.Run("first", func(t *testing.T) {
		s := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
		expected := "Here's my spammy page: http://******************* see you."
		substring := "http://"
		result := SpamMasker(s, substring)

		if result != expected {
			t.Errorf("incorrect result, expect %s, got %s", expected, result)
		}
	})

	t.Run("second", func(t *testing.T) {
		s := "Here's my spammy page: http://hehefouls.net HAHAHA see you."
		expected := "Here's my spammy page: http://************* HAHAHA see you."
		substring := "http://"
		result := SpamMasker(s, substring)

		if result != expected {
			t.Errorf("incorrect result, expect %s, got %s", expected, result)
		}
	})

	t.Run("third", func(t *testing.T) {
		s := "Here's my spammy page: hTTp://youth-elixir.com"
		expected := "Here's my spammy page: hTTp://youth-elixir.com"
		substring := "http://"
		result := SpamMasker(s, substring)

		if result != expected {
			t.Errorf("incorrect result, expect %s, got %s", expected, result)
		}
	})
}
