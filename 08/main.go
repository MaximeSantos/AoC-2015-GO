package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Open input.txt file and return its content by scanning it line by line
func input() []string {
	var input []string

	// Open file
	file, err := os.Open("input.txt")
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

// Given an array of strings, returns the total number of characters
func part1_countAll(str []string) int {
	total := 0

	for _, s := range str {
		total += len(s)
	}

	return total
}

// Given an array of strings, returns the total number of escaped characters
func part1_countEscaped(str []string) int {
	total := 0

	for _, s := range str {
		unquoted, _ := strconv.Unquote(s)
		total += len(unquoted)
	}

	return total
}

func part2_countEncoded(str []string) int {
	total := 0

	for _, s := range str {
		encoded := strconv.Quote(s)
		total += len(encoded)
	}

	return total
}

// Prints the number of characters of code for string literals minus the number of characters in memory for the values of the strings in total for the entire file
func part1() {
	input := input()

	c := part1_countAll(input)
	e := part1_countEscaped(input)

	total := c - e

	fmt.Println("Part 1:")
	fmt.Println("Total number of characters :", c)
	fmt.Println("Total number of escaped characters :", e)
	fmt.Println("Total number of characters - escaped characters :", total)
}

func part2() {
	input := input()

	c := part1_countAll(input)
	e := part2_countEncoded(input)

	total := e - c
	fmt.Println("Part 2:")
	fmt.Println("Total number of encoded characters :", e)
	fmt.Println("Total number of characters :", c)
	fmt.Println("Total number of encoded characters - characters :", total)
}

func main() {
	part1()
	part2()
}

/*
Escaped characters :
	\\ = \
	\" = "
	\x + two hexadecimal charaters = corresponding character according to ASCII code
*/
