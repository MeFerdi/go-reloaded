package student

import (
	"testing"

	"student/student"
)

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello world"},
		{"hello world(cap), this is Ferdinand", "Hello this is Ferdinand"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := student.Capitalize(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}

// func Capitalize(s string) string {
// 	words := strings.Split(s, " ")

// 	for i := 0; i < len(words); i++ {
// 		if strings.Contains(words[i], "(cap") {
// 			if strings.Contains(words[i], "(cap,") {
// 				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
// 				if i-num >= 0 && i-num < len(words) {
// 					for j := i - num; j < i; j++ {
// 						if j >= 0 && j < len(words) {
// 							words[j] = strings.ToUpper(string([]rune(words[j])[0])) + strings.ToLower(words[j][1:])
// 						}
// 					}
// 					words = append(words[:i], words[i+2:]...)
// 				}
// 			} else {
// 				if i >= 0 && i < len(words) {
// 					words[i-1] = strings.ToUpper(string([]rune(words[i-1])[0])) + strings.ToLower(words[i-1][1:])
// 					words = append(words[:i], words[i+1:]...)
// 				}
// 			}
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
