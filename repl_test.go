package main

import (
	"fmt"
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "",
			expected: []string{},
		},
		{
			input: "   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  world  ",
			expected: []string{"world"},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)

		fmt.Printf("\nTest %d of %d:\n", i + 1, len(cases))
		fmt.Printf("input: %v\n", c.input)
		fmt.Printf("expected: %v\n", stringifiedSlice(c.expected))
		fmt.Printf("actual: %v\n", stringifiedSlice(actual))

		fmt.Println("")

		if len(actual) != len(c.expected) {
			t.Errorf("FAILED\nActual size: %v\nExpected size: %v", len(actual), len(c.expected))
			t.Errorf("Actual content: %v", stringifiedSlice(actual))
			return
		}

		failedCount := 0
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("FAILED:\n(Actual) %v != (Expected) %v ", word, expectedWord)
				failedCount++
			}
		}

		if failedCount > 0 {
			t.Fatalf("FAILED: %d mismatched", failedCount)
		} 

	}

}

func stringifiedSlice[T any](arr []T) string {
	text := "["
	for i, item := range arr {
		if i == len(arr) - 1 {
			text += fmt.Sprintf("\"%v\"", item)
		} else {
			text += fmt.Sprintf("\"%v\",", item)
		}
	}
	text += "]"
	return text
}