package student

import (
	"strings"
)

func UniquePunctuation(s string) string {
	words := strings.Fields(s)
	count := 0

	for i := 0; i < len(words); i++ {
		if words[i] == "'" && count == 0 && i+1 < len(words) {
			words[i+1] = "'" + words[i+1]
			words = append(words[:i], words[i+1:]...)
			count++
		} else if words[i] == "'" && count != 0 && i-1 >= 0 {
			words[i-1] = words[i-1] + "'"
			words = append(words[:i], words[i+1:]...)
			count--
			i--
		} else if words[i] == "'" && (i == 0 || i == len(words)-1) {
			if i == 0 && i+1 < len(words) {
				words[i+1] = "'" + words[i+1]
			} else if i == len(words)-1 {
				words[i-1] = words[i-1] + "'"
			} else {
				count++
			}
		}
	}

	return strings.Join(words, " ")
}

// 	words := strings.Fields(s)
// 	count := 0

// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "'" && count == 0 {

// 			words[i+1] = "'" + words[i+1]

// 			words = append(words[:i], words[i+1:]...)

// 			count++
// 		} else if words[i] == "'" && count != 0 {

// 			words[i-1] = words[i-1] + "'"

// 			words = append(words[:i], words[i+1:]...)

// 			count--

// 			i--
// 		}
// 	}

// 	return strings.Join(words, " ")
// }
