package main

import (
	"fmt"
	"strconv"
	"strings"

	"AoC/15/utils"
)

// Example string : "children: 3"
func parse_sue(s []string) map[string]int {
	ret := make(map[string]int)

	for _, curr := range s {
		curr = strings.ReplaceAll(curr, ":", "")
		split := strings.Split(curr, " ")
		value, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		ret[split[0]] = value
	}

	return ret
}

// Example string : "Sue 1: children: 1, cars: 8, vizslas: 7"
func parse_input(s []string) []map[string]int {
	ret := make([]map[string]int, 0)

	for _, curr := range s {
		m := make(map[string]int)
		// remove all ":", "'", the first two words (Sue 1) and split our string into slice of words
		curr = strings.ReplaceAll(curr, ":", "")
		curr = strings.ReplaceAll(curr, ",", "")
		split := strings.Split(curr, " ")
		split = split[2:]

		for j := 0; j < len(split); j++ {
			// Check if current word is a number
			// If not, then it is the key we need to add to our map
			// Add its associated value at j+1
			_, err := strconv.Atoi(split[j])
			if err != nil {
				value, err := strconv.Atoi(split[j+1])
				if err != nil {
					panic(err)
				}
				m[split[j]] = value
				// skip next iteration
				j++
				continue
			}
		}
		ret = append(ret, m)
	}
	return ret
}

func whodunnit(sue map[string]int, clues []map[string]int) int {
	for i, clue := range clues {
		j := 0
		// Go through our clue and check if the values coincide with what we know of Sue.
		for key, value := range clue {
			if sue[key] != value {
				break
			}
			// If we manage to get to the last element...
			// Then we can return the index of the current sue (+ 1 because of indexing from 1 and not 0 as instructed)
			if j == len(clue)-1 {
				return i + 1
			}
			j++
		}
	}
	// Error
	return 0
}

func whodunnit_part2(sue map[string]int, clues []map[string]int) int {
	for i, clue := range clues {
		// Counter for our clue's index -- maps have no index in Go, only keys
		j := 0
		l := len(clue) - 1
		// Go through our clue and check if the values coincide with what we know of Sue.
		for key, value := range clue {
			if key == "cats" || key == "trees" {
				if sue[key] >= value {
					break
				}
			} else if key == "pomeranians" || key == "goldfish" {
				if sue[key] <= value {
					break
				}
			} else {
				if sue[key] != value {
					break
				}
			}
			// If we manage to get to the last element...
			// Then we can return the index of the current sue (+ 1 because of indexing from 1 and not 0 as instructed)
			if j == l {
				return i + 1
			}
			j++
		}
	}
	// Error
	return 0
}

func part1() {
	sue := parse_sue(utils.Input("input_sue.txt"))
	input := parse_input(utils.Input("input.txt"))
	n_sue := whodunnit(sue, input)

	fmt.Println("Part 1")
	fmt.Println("The culprit is sue number:", n_sue)
}

func part2() {
	sue := parse_sue(utils.Input("input_sue.txt"))
	input := parse_input(utils.Input("input.txt"))
	n_sue := whodunnit_part2(sue, input)

	fmt.Println("----------------")
	fmt.Println("Part 2")
	fmt.Println("The culprit is sue number:", n_sue)
}

func main() {
	part1()
	part2()
}
