package main

import (
	"fmt"
	"strconv"
	"strings"

	"AoC/15/utils"
)

type Ingredient struct {
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func parse(s []string) []Ingredient {
	var ret []Ingredient

	for _, curr := range s {
		// Clean up the string by removing ponctuation & split it
		// Example string : "Sprinkles: capacity 5, durability -1, flavor 0, texture 0, calories 5"
		curr = strings.ReplaceAll(curr, ":", "")
		curr = strings.ReplaceAll(curr, ",", "")
		split := strings.Split(curr, " ")

		capacity, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}
		durability, err := strconv.Atoi(split[4])
		if err != nil {
			panic(err)
		}
		flavor, err := strconv.Atoi(split[6])
		if err != nil {
			panic(err)
		}
		texture, err := strconv.Atoi(split[8])
		if err != nil {
			panic(err)
		}
		calories, err := strconv.Atoi(split[10])
		if err != nil {
			panic(err)
		}

		ing := Ingredient{Capacity: capacity, Durability: durability, Flavor: flavor, Texture: texture, Calories: calories}
		ret = append(ret, ing)
	}

	return ret
}

// Returns a list of all combinations for an array of certain size, up to a maximum integer
// Copied from https://github.com/mkst/aoc/blob/master/2015/15.go
// Could not figure out the algorithm on my own
func combin(size int, max int) [][]int {
	var ret [][]int
	var queue [][]int

	// Get all combinations
	for i := 0; i < max; i++ {
		item := []int{i}
		queue = append(queue, item)
	}

	if size == 1 {
		return queue
	}

	for {
		if len(queue) == 0 {
			break
		}

		var item []int
		item, queue = queue[0], queue[1:]

		sum := 0
		for _, v := range item {
			sum += v
		}

		if len(item) == size-1 {
			ret = append(ret, append(item, max-sum))
			continue
		}

		for i := 0; i <= max-sum; i++ {
			queue = append(queue, append(item, i))
		}
	}

	return ret
}

func brute(s []Ingredient) (int, int) {
	var ret, ret_part2 int

	combin := combin(len(s), 100)

	for _, curr := range combin {
		var (
			capacity   int
			durability int
			flavor     int
			texture    int
			calories   int
		)

		for i, n := range curr {
			capacity += n * s[i].Capacity
			durability += n * s[i].Durability
			flavor += n * s[i].Flavor
			texture += n * s[i].Texture
			calories += n * s[i].Calories
		}

		if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
			continue
		}

		sum := capacity * durability * flavor * texture

		if sum > ret {
			ret = sum
		}
		if sum > ret_part2 && calories == 500 {
			ret_part2 = sum
		}
	}

	return ret, ret_part2
}

func part1_2() {
	parsed := parse(utils.Input("input.txt"))
	best, best_part2 := brute(parsed)

	fmt.Println("Part 1")
	fmt.Println("Highest score is :", best)
	fmt.Println("----------")
	fmt.Println("Part 2")
	fmt.Println("Highest score for a 500 calories cookie is :", best_part2)
}

func main() {
	part1_2()
}
