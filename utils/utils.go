package utils

import (
	"bufio"
	"os"
)

// Open input.txt file and return its content by scanning it line by line
func Input(s string) []string {
	var input []string

	// Open file
	file, err := os.Open(s)
	if err != nil {
		// fmt.Println("Error opening file", err)
		panic(err)
	}
	defer file.Close()

	// Scan/Parse file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Add parsed line to our return variable
		input = append(input, line)
	}

	return input
}
