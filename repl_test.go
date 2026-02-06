package main

import (
	"testing"
)

type cleanInputTestCase struct {
		input    string
		expected []string
}

func TestCleanInput(t *testing.T) {
	cases := map[string]cleanInputTestCase{
		"many spaces": {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		"leading and trailing spaces": {
			input:    "  hello pokemon world  ",
			expected: []string{"hello", "pokemon", "world"},
		},
		"no spaces": {
			input:    "/hello/world/",
			expected: []string{"/hello/world/"},
		},
		// add more cases here
	}

	for name, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Fatalf("%s Wrong len. CleanInput(%q) = %q; want %q", name, c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Fatalf("%s CleanInput(%q) = %q; want %q", name, c.input, actual, c.expected)
			}
		}
	}
}
