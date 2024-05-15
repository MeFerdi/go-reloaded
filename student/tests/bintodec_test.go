package student

import (
	"testing"

	"student/student"
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
			actual := student.BinToDecimal(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}

// func BinToDecimal(s string) string {
// 	words := strings.Fields(s)
// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(bin)" {
// 			binValue, _ := strconv.ParseInt(words[i-1], 2, 64)
// 			words[i-1] = fmt.Sprintf("%d", binValue)
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
