package main

import (
	"fmt"
	"strings"
)

func FormatPunctuation(s string) string {
	MyPunctuations := []string{",", ".", "!", "?", ":", ";"}

	MyString := strings.Split(s, " ")
	for i := 0; i < len(MyString); i++ {
		TheString := MyString[i]
		for _, ThePunctuation := range MyPunctuations {
			if (string(TheString[0]) == ThePunctuation) && (string(TheString[len(TheString)-1]) != ThePunctuation) && TheString != MyString[len(MyString)-1] {
				MyString[i-1] = MyString[i-1] + ThePunctuation
				MyString[i] = TheString[1:]
			}
			if (string(TheString[0]) == ThePunctuation) && TheString == MyString[len(MyString)-1] {
				MyString[i-1] = MyString[i-1] + TheString
				MyString = MyString[:len(MyString)-1]
			}
			if (string(TheString[0]) == ThePunctuation) && (string(TheString[len(TheString)-1]) == ThePunctuation) {
				MyString[i-1] = MyString[i-1] + TheString
				MyString = append(MyString[:i], MyString[i+1:]...)
			}
		}

	}
	return strings.Join(MyString, " ")
}

func main() {
	fmt.Println(FormatPunctuation("Punctuation tests are ... kinda boring ,don't you think !?"))
}
