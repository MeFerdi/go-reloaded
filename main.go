package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := modifyText(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hi! your Text completely modified.")
}

func modifyText(inputFile, outputFile string) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := applyModifications(line)
		fmt.Fprintln(output, modifiedLine)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func applyModifications(line string) string {
	words := strings.Fields(line)
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 && isHexadecimal(words[i-1]) {
				decimalValue := convertToDecimal(words[i-1], 16)
				words[i-1] = decimalValue
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(bin)":
			if i > 0 && isBinary(words[i-1]) {
				decimalValue := convertToDecimal(words[i-1], 2)
				words[i-1] = decimalValue
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = strings.Title(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		default:
			if strings.HasPrefix(words[i], "(up)") || strings.HasPrefix(words[i], "(low)") || strings.HasPrefix(words[i], "(cap)") {
				modifyCase(&words, i)
			} else if i > 0 && isPunctuation(words[i]) && len(words) > i+1 && isPunctuation(words[i+1]) {
				words[i] = strings.TrimSpace(words[i])
			} else if words[i] == "'" && i > 0 && i < len(words)-1 {
				words[i-1], words[i+1] = "'"+words[i-1], words[i+1]+"'"
				words = removeEmptyElement(words, i)
			} else if words[i] == "a" && i < len(words)-1 {
				if isVowelorH(words[i+1]) {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}

func modifyCase(words *[]string, index int) {
	parts := strings.Split((*words)[index][1:len((*words)[index])-1], ",")
	if len(parts) != 2 {
		return
	}

	count := 1
	if len(parts) > 1 {
		count, _ = strconv.Atoi(parts[1])
	}
	if count < 1 || index-count < 0 {
		return
	}

	switch parts[0] {
	case "up":
		for i := index - count; i < index; i++ {
			(*words)[i] = strings.ToUpper((*words)[i])
		}
	case "low":
		for i := index - count; i < index; i++ {
			(*words)[i] = strings.ToLower((*words)[i])
		}
	case "cap":
		for i := index - count; i < index; i++ {
			(*words)[i] = strings.Title((*words)[i])
		}
	}
}

func isHexadecimal(s string) bool {
	_, err := strconv.ParseInt(s, 16, 64)
	return err == nil
}

func isBinary(s string) bool {
	_, err := strconv.ParseInt(s, 2, 64)
	return err == nil
}

func convertToDecimal(s string, base int) string {
	decimalValue, _ := strconv.ParseInt(s, base, 64)
	return strconv.FormatInt(decimalValue, 10)
}

func isPunctuation(s string) bool {
	punctuations := ".,!?:;"
	return strings.ContainsRune(punctuations, []rune(s)[0])
}

func removeEmptyElement(slice []string, index int) []string {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func isVowelorH(s string) bool {
	firstChar := unicode.ToLower(rune(s[0]))
	return firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'h'
}
