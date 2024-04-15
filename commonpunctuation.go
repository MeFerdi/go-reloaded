package student

import (
	"strings"
)

func FormatPunctuation(s string) string {
	MyPunctuations := []string{",", ".", "!", "?", ":", ";"}

	MyString := strings.Split(s, " ")
	for i := 1; i < len(MyString); i++ {
		TheString := MyString[i]
		for _, ThePunctuation := range MyPunctuations {
			if string(TheString[0]) == ThePunctuation && string(TheString[len(TheString)-1]) != ThePunctuation {
				MyString[i-1] = MyString[i-1] + ThePunctuation
				MyString[i] = TheString[1:]
			}
			if string(TheString[0]) == ThePunctuation && string(TheString[len(TheString)-1]) == ThePunctuation {
				MyString[i-1] = MyString[i-1] + TheString
				MyString = append(MyString[:i], MyString[i+1:]...)
			}
		}

	}
	return strings.Join(MyString, " ")
}
