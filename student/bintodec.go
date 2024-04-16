package student

import (
	"fmt"
	"strconv"
	"strings"
)

func BinToDecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" {
			binValue, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = fmt.Sprintf("%d", binValue)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
