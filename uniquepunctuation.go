package student

import (
	"strings"
)

func UniquePunctuation(s string) string {
	words := strings.Fields(s)
	count := 0

	for i := 0; i < len(words); i++ {
		if words[i] == "'" && count == 0 {

			words[i+1] = "'" + words[i+1]

			words = append(words[:i], words[i+1:]...)

			count++
		} else if words[i] == "'" && count != 0 {

			words[i-1] = words[i-1] + "'"

			words = append(words[:i], words[i+1:]...)

			count--

			i--
		}
	}

	return strings.Join(words, " ")
}
