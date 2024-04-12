package student

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToDecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexValue, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = fmt.Sprintf("%d", hexValue)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
