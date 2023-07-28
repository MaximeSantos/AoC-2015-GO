package main

import (
	"fmt"
	"strings"
)

// func vowel(s string) bool {
// 	var count int
// 	for i := 0; i < len(s); i++ {
// 		switch string(s[i]) {
// 		case "a", "e", "i", "o", "u":
// 			count++
// 		}
// 		if count == 3 {
// 			return true
// 		}
// 	}
// 	return false
// }

// func double(s string) bool {
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == s[i-1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func naughty(s string) bool {
// 	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
// 		return false
// 	}
// 	return true
// }

// Checks that the given string contains at least a pair of letters that repeats at least once
func pair(s string) bool {
	// starting at first letter, create a pair of letters
	for i := 0; i < len(s)-2; i++ {
		pair := s[i:(i + 2)]
		// now, starting at first next pair of letter, check if it equal to our selected one
		for j := i + 2; j < len(s)-1; j++ {
			if string(s[j])+string(s[j+1]) == pair {
				return true
			}
		}
	}
	return false
}

// Checks that the given string contains at least a letter that repeats after another letter
func repeat(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			return true
		}
	}
	return false
}

func nice(s string) bool {
	// Part1
	// must contain at least three vowels (aeiou)
	// must contain at least one double letter
	// must NOT contain ab, cd, pq, or xy
	// return naughty(s) && double(s) && vowel(s)

	// Part2
	// must constain a pair of letter that appears at least twice
	// must contain a letter that repeats with one extra letter in between
	return pair(s) && repeat(s)
}

func main() {
	var counter int
	slice := strings.Fields(input)

	for _, s := range slice {
		if nice(s) {
			counter++
		}
	}
	fmt.Printf("Number of nice strings : %d\n", counter)
}

// --- Day 5: Doesn't He Have Intern-Elves For This? ---

// Santa needs help figuring out which strings in his text file are naughty or nice.

// A nice string is one with all of the following properties:

//     - It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
//     - It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
//     - It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

// For example:

//     - ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
//     - aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
//     - jchzalrnumimnmhp is naughty because it has no double letter.
//     - haegwjzuvuyypxyu is naughty because it contains the string xy.
//     - dvszwmarrgswjxmb is naughty because it contains only one vowel.

// How many strings are nice?

// ! 258

// --- Part Two ---

// Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

// Now, a nice string is one with all of the following properties:

//     It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
//     It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

// For example:

//     qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
//     xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
//     uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
//     ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.

// How many strings are nice under these new rules?

// ! 53
