package main

import (
	"AoC/15/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Happiness struct {
	guest  string
	nextTo string
	change int
}

func parse(in []string) []Happiness {
	slice := make([]Happiness, len(in))

	for i, line := range in {
		split := strings.Split(line, " ")

		// Store string number that represents happiness change as an int
		n, err := strconv.Atoi(split[3])
		if err != nil {
			panic("Error, input unconsistent - no number at index 3")
		}

		// add our Happiness to our slice, only guest & nextTo value for now
		slice[i] = Happiness{guest: split[0], nextTo: strings.TrimSuffix(split[10], ".")}

		// Check if the guest gains or loses value from being sitted next to the other person
		// Add the the corresponding positive or negative value to our Happiness struct
		if split[2] == "gain" {
			slice[i].change = n
		} else if split[2] == "lose" {
			slice[i].change = -n
		} else {
			panic("Error, input unconsistent - no lose or gain at index 2")
		}
	}

	return slice
}

// Return a list of unique guests
func getAllGuests(s []Happiness) []string {
	var list []string

	for _, l := range s {
		if !slices.Contains(list, l.guest) {
			list = append(list, l.guest)
		}
	}

	return list
}

// Get all possible combinations for a given array of strings
func permutate(s []string) [][]string {
	var res [][]string
	var f func(i int)

	perm := make([]string, len(s))
	indexInUse := make([]bool, len(s))

	f = func(i int) {
		if i >= len(s) {
			arr := make([]string, len(s))
			copy(arr, perm)
			res = append(res, arr)
			return
		}
		for j := 0; j < len(s); j++ {
			if !indexInUse[j] {
				indexInUse[j] = true
				perm[i] = s[j]
				f(i + 1)
				indexInUse[j] = false
			}
		}
	}

	f(0)
	return res
}

// Calculate the change in happiness for a given sit configuration
func getHappiness(sit []string, list []Happiness) int {
	sum := 0

	for i := range sit {
		if i == 0 {
			idxRight := slices.IndexFunc(list, func(h Happiness) bool {
				return sit[i] == h.guest && sit[i+1] == h.nextTo
			})
			idxLeft := slices.IndexFunc(list, func(h Happiness) bool {
				return sit[i] == h.guest && sit[len(sit)-1] == h.nextTo
			})

			sum += list[idxLeft].change
			sum += list[idxRight].change

		} else if i == len(sit)-1 {
			idxRight := slices.IndexFunc(list, func(h Happiness) bool {
				return sit[i] == h.guest && sit[0] == h.nextTo
			})
			idxLeft := slices.IndexFunc(list, func(h Happiness) bool {
				return sit[i] == h.guest && sit[i-1] == h.nextTo
			})

			sum += list[idxLeft].change
			sum += list[idxRight].change
		} else {
			idxRight := slices.IndexFunc(list, func(h Happiness) bool {
				return sit[i] == h.guest && sit[i+1] == h.nextTo
			})
			idxLeft := slices.IndexFunc(list, func(h Happiness) bool {
				return sit[i] == h.guest && sit[i-1] == h.nextTo
			})

			sum += list[idxLeft].change
			sum += list[idxRight].change
		}
	}

	return sum
}

// We permutate our list of guests to get all sits configurations
// Then calculate the optimal happiness change output out of all of those
func brute(list []Happiness) int {
	best := 0
	guests := getAllGuests(list)
	permGuests := permutate(guests)

	for _, p := range permGuests {
		currSum := getHappiness(p, list)
		if currSum > best {
			best = currSum
		}
	}

	return best
}

func part1() {
	input := utils.Input("input.txt")
	parsed := parse(input)
	bestChange := brute(parsed)

	fmt.Println("Part 1")
	fmt.Println("Best happiness change is", bestChange)
}

func part2() {
	input := utils.Input("input2.txt")
	parsed := parse(input)
	bestChange := brute(parsed)

	fmt.Println("-------------")
	fmt.Println("Part 2")
	fmt.Println("Best happiness change is", bestChange)
}

func main() {
	part1()
	part2()
}
