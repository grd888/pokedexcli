package cli

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Create a CLI instance for testing
	cli := NewCLI()
	
	cases := []struct {
		input    string
		expected []string
		name     string
	}{
		{input: " hello world  ", expected: []string{"hello", "world"}, name: "Leading/trailing spaces"},
		{input: " Hello World  ", expected: []string{"hello", "world"}, name: "Mixed case"},
		{input: "", expected: []string{}, name: "Empty string"},
		{input: "   ", expected: []string{}, name: "Only spaces"},
		{input: "  leading ", expected: []string{"leading"}, name: "Leading spaces only"},
		{input: "trailing   ", expected: []string{"trailing"}, name: "Trailing spaces only"},
		{input: "  multiple   spaces  between ", expected: []string{"multiple", "spaces", "between"}, name: "Multiple spaces"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cli.cleanInput(c.input)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("cleanInput(%q) == %q, want %q", c.input, actual, c.expected)
			}
		})
	}
}
