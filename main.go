package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("[Usage] go run main.go <file-name> <search-word>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("[Error] file cannot open:", err)
	}
	defer file.Close()

	word := os.Args[2]

	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			coloredWord := fmt.Sprintf("\033[31m%s\033[0m", word)
			lineWithColoredWord := strings.Replace(line, word, coloredWord, -1)
			fmt.Printf("line %d: %s\n", lineNum, lineWithColoredWord)
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[Error] error occurred while reading file:", err)
	}
}
