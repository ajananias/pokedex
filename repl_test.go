package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "       ",
			expected: []string{},
		},
		{
			input:    "right side Spaced TEST   ",
			expected: []string{"right", "side", "spaced", "test"},
		},
		{
			input:    "    multiple    spACEs   inside   text   ",
			expected: []string{"multiple", "spaces", "inside", "text"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Failed: cleanInput(%q) = %v; want %v", c.input, actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Failed: cleanInput(%q)[%d] = %q; want %q", c.input, i, word, expectedWord)
			}
		}
	}
}
