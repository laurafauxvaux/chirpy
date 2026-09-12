package main

import "testing"

func TestCleanMessage(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello world",
			expected: "hello world",
		},
		{
			input:    "This is kerfuffle opinion",
			expected: "This is **** opinion",
		},
		{
			input:    "kerfuffle sharbert fornax",
			expected: "**** **** ****",
		},
	}

	for _, c := range cases {
		actual := cleanMessage(c.input)

		if actual != c.expected {
			t.Errorf("got %q, want%q", actual, c.expected)
		}
	}
}
