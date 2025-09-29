package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " now    testing CleanInput",
			expected: []string{"now", "testing", "CleanInput"},
		},
		{
			input:    "Check  the length   of the actual slice",
			expected: []string{"Check", "the", "length", "of", "the", "actual", "slice"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected: %v\n Actual: %v", c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected: %v\n Actual: %v", c.expected, actual)
			}

		}

		fmt.Println("===================")
		fmt.Printf("Input: %v\nExpected: %v\nActual: %v\n", c.input, c.expected, actual)
		fmt.Println("===================")
	}

}
