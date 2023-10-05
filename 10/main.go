package main

import (
	"fmt"
)

var input = 1113122113

// https://en.wikipedia.org/wiki/Look-and-say_sequence

func splitToDigits(n int) []int {
	var res []int

	for n != 0 {
		// Prepend last digit of our number to our slice
		// Otherwise we could also append normaly and reverse the slice afterwards
		res = append([]int{n % 10}, res...)
		// Removes last digit of our number
		n /= 10
	}

	return res
}

func say(n []int) []int {
	var res []int

	currDig := 0
	nCurrDig := 0

	for i, d := range n {
		// count each sequence of digits
		switch {
		// Special case for first digit of sequence
		case currDig == 0:
			currDig = d
			nCurrDig++
		case i == (len(n) - 1):
			if d == currDig {
				nCurrDig++
				res = append(res, nCurrDig)
				res = append(res, currDig)
			} else {
				res = append(res, nCurrDig)
				res = append(res, currDig)
				res = append(res, 1)
				res = append(res, d)
			}
			// Then increment nCurrDig if current digit is the same as previous
		case d == currDig:
			nCurrDig++
			// If we change digit or reach the end of our slice
		case d != currDig:
			// Append current digit to our res the number of time it appeared in original sequence
			res = append(res, nCurrDig)
			res = append(res, currDig)
			currDig = d
			nCurrDig = 1
		}
		// fmt.Println("Line:", i, "digit:", d, "currDig:", currDig, "nCurrDig:", nCurrDig)
	}

	return res
}

func part1() {
	digits := splitToDigits(input)

	for i := 0; i < 40; i++ {
		digits = say(digits)
	}

	fmt.Println("Look & Say applied 40 times :")
	fmt.Println(len(digits), "digits")
	fmt.Println("----------------------------")
}

func part2() {
	digits := splitToDigits(input)

	for i := 0; i < 50; i++ {
		digits = say(digits)
	}

	fmt.Println("Look & Say applied 50 times :")
	fmt.Println(len(digits), "digits")
	fmt.Println("----------------------------")
}

func main() {
	part1()
	part2()
}
