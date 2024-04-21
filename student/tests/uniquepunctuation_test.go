package student

import (
	"testing"
)

func TestUniquePunctuation(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello world"},
		{"hello 'world'", "hello 'world'"},
		{"hello 'world' '", "hello 'world'"},
		{"hello ' world'", "hello 'world'"},
		{"hello world '", "hello world'"},
		{"hello 'world' world", "hello 'world' world"},
		{"hello 'world' 'world'", "hello 'world' 'world'"},
		{"hello 'world' ' world'", "hello 'world' 'world'"},
		{"hello ' world' world", "hello 'world' world"},
		{"hello ' world' 'world'", "hello 'world' 'world'"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := UniquePunctuation(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}
