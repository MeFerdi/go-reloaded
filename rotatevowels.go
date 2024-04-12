package student

import (
	"strings"
)

func IsVowelorH(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "a" || words[i] == "A" {
			// Replace "a" with "an" and vice versa
			if i+1 < len(words) && (words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'i' || words[i+1][0] == 'o' || words[i+1][0] == 'u' || words[i+1][0] == 'h') {
				words[i] = "an"
			} else {
				words[i] = "a"
			}
		} else if words[i] == "an" {
			// Replace "an" with "a"
			words[i] = "a"
		} else if words[i] == "h" || words[i] == "H" {
			// Handle the letter "h" as needed
			// For example, you can choose to keep it as is or replace it with something else
			// Here, we are keeping it as is
		}
	}
	return strings.Join(words, " ")
}
