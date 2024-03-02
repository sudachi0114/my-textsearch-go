package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("[Error] file cannot open:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[Error] error occurred while reading file:", err)
	}
}
