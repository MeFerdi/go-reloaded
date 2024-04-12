package main

import (
	"fmt"
	"strings"
)

func formatPunctuation(s string) string {
	var sb strings.Builder
	prevWasSpace := true
	inGroup := false
	for i, c := range s {
		if c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';' {
			if !prevWasSpace {
				sb.WriteRune(' ')
			}
			if i+2 < len(s) && (s[i+1] == c && s[i+2] == c) {
				inGroup = true
				sb.WriteRune(c)
				sb.WriteRune(c)
				sb.WriteRune(c)
				i += 2
			} else if i+1 < len(s) && (s[i+1] == '.' || s[i+1] == '!' || s[i+1] == '?') {
				inGroup = true
				sb.WriteRune(c)
				sb.WriteRune(s[i+1])
				i++
			} else {
				sb.WriteRune(c)
				sb.WriteRune(' ')
			}
			prevWasSpace = false
		} else {
			if inGroup {
				sb.WriteRune(c)
			} else {
				if c == ' ' {
					prevWasSpace = true
				} else {
					sb.WriteRune(c)
					prevWasSpace = false
				}
			}
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(formatPunctuation("I was sitting over there ,and then BAMM !!"))
	fmt.Println(formatPunctuation("I was thinking ... You were right"))
}
