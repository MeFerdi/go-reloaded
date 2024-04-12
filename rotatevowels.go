package main

import (
	"fmt"
	"strings"
	"unicode"
)

func formatA(s string) string {
	var sb strings.Builder
	prevWasSpace := true
	nextIsVowel := false
	for i, c := range s {
		if c == ' ' {
			prevWasSpace = true
		} else if c == 'a' || c == 'A' {
			if i+1 < len(s) {
				nextChar := s[i+1]
				if unicode.IsVowel(rune(nextChar)) || nextChar == 'h' || nextChar == 'H' {
					nextIsVowel = true
				} else {
					nextIsVowel = false
				}
			}
			if nextIsVowel {
				sb.WriteString("an")
			} else {
				sb.WriteString("a")
			}
			prevWasSpace = false
		} else {
			sb.WriteRune(c)
			prevWasSpace = false
		}
		if !prevWasSpace {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(formatA("There it was. A amazing rock!"))
}
