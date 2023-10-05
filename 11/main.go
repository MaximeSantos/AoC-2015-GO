package main

import (
	"errors"
	"fmt"
	"strings"
)

var input = "vzbxkghb"

/*
Increment the string
- start with rightmost letter until it wraps from z to a
- go to next letter on the left and keep going

On each step, check for three conditions :
- On straight of at least three following letters
- No letter I, O or L
- At least two different double letters

That is Santa's next password
*/

// Dirty increment for 8 characters long strings
// TODO Cleanup by doing a proper loop
func increment(s string) string {
	var res string

	b := []byte(s)

	if b[len(b)-1] < 122 {
		b[len(b)-1] = b[len(b)-1] + 1
	} else if b[len(b)-2] < 122 {
		b[len(b)-1] = 97
		b[len(b)-2] = b[len(b)-2] + 1
	} else if b[len(b)-3] < 122 {
		b[len(b)-1] = 97
		b[len(b)-2] = 97
		b[len(b)-3] = b[len(b)-3] + 1
	} else if b[len(b)-4] < 122 {
		b[len(b)-1] = 97
		b[len(b)-2] = 97
		b[len(b)-3] = 97
		b[len(b)-4] = b[len(b)-4] + 1
	} else if b[len(b)-5] < 122 {
		b[len(b)-1] = 97
		b[len(b)-2] = 97
		b[len(b)-3] = 97
		b[len(b)-4] = 97
		b[len(b)-5] = b[len(b)-5] + 1
	} else if b[len(b)-6] < 122 {
		b[len(b)-1] = 97
		b[len(b)-2] = 97
		b[len(b)-3] = 97
		b[len(b)-4] = 97
		b[len(b)-5] = 97
		b[len(b)-6] = b[len(b)-6] + 1
	} else if b[len(b)-7] < 122 {
		b[len(b)-1] = 97
		b[len(b)-2] = 97
		b[len(b)-3] = 97
		b[len(b)-4] = 97
		b[len(b)-5] = 97
		b[len(b)-6] = 97
		b[len(b)-7] = b[len(b)-7] + 1
	} else {
		return "aaaaaaaa"
	}

	res = string(b)

	return res
}

// Check that a string of characters contains at least one suite of 3 characters that follow each other alphabetically
// Returns true if everything is in order
func checkForStraight(s string) bool {
	b := []byte(s)

	// Loop stops before the last two characters
	// Starting character cannot be "y"
	for i := 0; i < len(b)-2; i++ {
		if b[i] < 121 && b[i+1] == b[i]+1 && b[i+2] == b[i]+2 {
			return true
		}
	}

	return false
}

// Check that the given string does not contain the letters "i" "o" or "l"
// Returns true if everything is in order
func checkForForbidden(s string) bool {
	if strings.Contains(s, "i") || strings.Contains(s, "o") || strings.Contains(s, "l") {
		return false
	}

	return true
}

// Check that the given string contains at least two different double characters
// Returns true if everything is in order
func checkForDouble(s string) bool {
	tmp := ""

	b := []byte(s)

	for i := 0; i < len(b)-1; i++ {
		// Check that we have two of the same characters for our current pair
		if b[i] == b[i+1] {
			pair := string(b[i]) + string(b[i+1])

			if tmp != "" && tmp != pair {
				return true
			}

			tmp = pair
		}
	}

	return false
}

func lookForPassword(s string) (string, error) {
	curr := s

	// Limited loop to begin with
	for i := 0; i < 1000000000; i++ {
		curr = increment(curr)
		if checkForForbidden(curr) && checkForDouble(curr) && checkForStraight(curr) {
			// fmt.Println("Found password in", i-1, "increments")
			return curr, nil
		}
	}

	return "", errors.New("failed to find correct password in the loop limit")

}

func part1() {
	fmt.Println("Part 1 :")

	firstPassword, err := lookForPassword(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("First password is :", firstPassword)
}

func part2() {
	fmt.Println("-----------------")
	fmt.Println("Part 2 :")

	firstPassword, err := lookForPassword(input)
	if err != nil {
		panic(err)
	}

	secondPassword, err := lookForPassword(firstPassword)
	if err != nil {
		panic(err)
	}

	fmt.Println("Second password is :", secondPassword)

}

func main() {
	part1()
	part2()
}
