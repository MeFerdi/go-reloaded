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

// func UniquePunctuation(s string) string {
// 	words := strings.Fields(s)
// 	count := 0

// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "'" && count == 0 && i+1 < len(words) {
// 			words[i+1] = "'" + words[i+1]
// 			words = append(words[:i], words[i+1:]...)
// 			count++
// 		} else if words[i] == "'" && count != 0 && i-1 >= 0 {
// 			words[i-1] = words[i-1] + "'"
// 			words = append(words[:i], words[i+1:]...)
// 			count--
// 			i--
// 		} else if words[i] == "'" && (i == 0 || i == len(words)-1) {
// 			if i == 0 && i+1 < len(words) {
// 				words[i+1] = "'" + words[i+1]
// 			} else if i == len(words)-1 {
// 				words[i-1] = words[i-1] + "'"
// 			} else {
// 				count++
// 			}
// 		}
// 	}

// 	return strings.Join(words, " ")
// }
