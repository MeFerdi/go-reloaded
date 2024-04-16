package student

import (
	"strings"
)

func FormatPunctuation(s string) string {
	MyPunctuations := []string{",", ".", "!", "?", ":", ";"}

	MyString := strings.Split(s, " ")
	for i := 0; i < len(MyString); i++ {
		TheString := MyString[i]
		for _, ThePunctuation := range MyPunctuations {
			if len(TheString) > 0 && len(TheString) > 1 {
				if string(TheString[0]) == ThePunctuation && string(TheString[len(TheString)-1]) != ThePunctuation && i != len(MyString)-1 {
					MyString[i-1] = MyString[i-1] + ThePunctuation
					MyString[i] = TheString[1:]
				} else if string(TheString[0]) == ThePunctuation && i == len(MyString)-1 {
					MyString[i-1] = MyString[i-1] + TheString
					MyString = MyString[:len(MyString)-1]
				} else if string(TheString[0]) == ThePunctuation && string(TheString[len(TheString)-1]) == ThePunctuation {
					MyString[i-1] = MyString[i-1] + TheString
					MyString = append(MyString[:i], MyString[i+1:]...)
					i-- // Decrement i to account for the removed element
				}
			}
		}
	}
	return strings.Join(MyString, " ")
}
