package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"student"
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
	if strings.Contains(line, "(cap") {
		line = student.Capitalize(line)
	}
	if strings.Contains(line, "(up)") {
		line = student.UpperCase(line)
	}
	if strings.Contains(line, "(low)") {
		line = student.LowerCase(line)
	}
	if strings.Contains(line, "(hex)") {
		line = student.HexToDecimal(line)
	}
	if strings.Contains(line, "(bin)") {
		line = student.BinToDecimal(line)
	}
	if strings.Contains(line, "a") {
		line = student.IsVowelorH(line)
	}

	return line
}
