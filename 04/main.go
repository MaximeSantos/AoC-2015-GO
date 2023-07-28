package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

var input string = "bgvyzdsv"

// Hash a string and returns the string of the hash
func toMD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

// Brute force looking for lowest positive number (no leading zero) that results in a md5 hash that starts with 5 zeros
func mine(input string) (int, bool) {
	// Hardcoded limit
	for i := 1; i < 10000000; i++ {
		// Join the input with our current number
		s := fmt.Sprintf("%s%d", input, i)
		// Hash the current string
		h := toMD5(s)

		if h[:6] == "000000" {
			return i, true
		}
	}
	return 0, false
}

func main() {
	// Prints time main takes to run
	start := time.Now()
	defer func(start time.Time) {
		d := time.Since(start)
		fmt.Printf("main took %dms to execute\n", d.Milliseconds())
	}(start)

	number, b := mine(input)
	if b {
		fmt.Printf("The number you are looking for is %d\n", number)
		return
	}
	fmt.Println("Failed to find the correct number")
}

// This kind of puzzle might be a good candidate for concurrency

// --- Day 4: The Ideal Stocking Stuffer ---

// Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

// To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

// For example:

//     If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
//     If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

// ! 254575

// --- Part Two ---

// Now find one that starts with six zeroes.

// ! 1038736
