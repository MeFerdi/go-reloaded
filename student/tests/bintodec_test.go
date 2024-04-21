package student

import (
	"testing"
)

func TestBinToDecimal(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"01 (bin)", "1"},
		{"1101 (bin)", "13"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := BinToDecimal(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}
