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

// func HexToDecimal(s string) string {
// 	words := strings.Fields(s)
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(hex)" && i > 0 {
// 			hexValue, _ := strconv.ParseInt(words[i-1], 16, 64)
// 			words[i-1] = fmt.Sprintf("%d", hexValue)
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
