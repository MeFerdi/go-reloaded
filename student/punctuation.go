package student

import (
	"strings"
)

func formatPunctuation(s string) string {
	var sb strings.Builder
	prevWasSpace := true
	inQuotes := false
	startOfQuotes := 0
	for i, c := range s {
		if c == '\'' {
			if i > 0 && s[i-1] == '\'' {
				if !inQuotes {
					sb.WriteString("'")
					startOfQuotes = sb.Len()
					inQuotes = true
				}
			} else {
				if inQuotes {
					sb.WriteString("'")
					inQuotes = false
					endOfQuotes := sb.Len()
					words := strings.Fields(sb.String()[startOfQuotes:endOfQuotes])
					sb.Reset()
					for j, word := range words {
						if j > 0 {
							sb.WriteString(" ")
						}
						sb.WriteString(word)
					}
					sb.WriteString("'")
					prevWasSpace = false
				} else {
					if prevWasSpace {
						sb.WriteString(" ")
					}
					sb.WriteString("'")
					prevWasSpace = false
				}
			}
		} else if c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';' {
			if !prevWasSpace {
				sb.WriteString(" ")
			}
			sb.WriteRune(c)
			prevWasSpace = true
		} else if c == ' ' || c == '\t' {
			prevWasSpace = true
		} else {
			if prevWasSpace {
				sb.WriteString(" ")
				prevWasSpace = false
			}
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
