package student

import (
	"testing"
)

func TestHexToDecimal(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"1E (hex)", "26"},
		{"23ae1 (hex)", "146049"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := HexToDecimal(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}
