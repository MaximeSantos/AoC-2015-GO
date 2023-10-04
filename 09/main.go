package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Travel struct {
	src  string
	dest string
	dist int
}

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

// Parse input into a slice of structs composed of origin (src), destination (dest), and distance (dist) of travel
func parse(str []string) []Travel {
	parsedInput := make([]Travel, len(str))

	for i, s := range str {
		splitInput := strings.Split(s, " ")

		src := splitInput[0]
		dest := splitInput[2]
		dist, err := strconv.Atoi(splitInput[4])
		if err != nil {
			panic(err)
		}

		parsedInput[i] = Travel{src, dest, dist}
	}

	return parsedInput
}

func getAllLocations(travels []Travel) []string {
	locations := make([]string, 0)

	for _, loc := range travels {
		if !slices.Contains(locations, loc.src) {
			locations = append(locations, loc.src)
		}
		if !slices.Contains(locations, loc.dest) {
			locations = append(locations, loc.dest)
		}
	}

	return locations
}

// Failed the challenge, for such a small sample of travels it was faster to just look for the solution with a pen & paper
// Gives a short route but not necessarily the shortest since it absolutely needs to build one route at a time
// To get the shortest, it would need to be able to build two or more routes at a time and join them at the end

// I could brute force the puzzle by getting all possible permutations of the unique locations
// Go through each, calculate and record the shortest and longest distances with or slice of Travels
// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
// https://en.wikipedia.org/wiki/Heap%27s_algorithm

func shortest(travels []Travel) int {
	distance := 0

	sortedToVisit := getAllLocations(travels)
	sort.Slice(sortedToVisit, func(i, j int) bool { return sortedToVisit[i] < sortedToVisit[j] })

	visited := make([]string, 0)

	fmt.Println("TO VISIT :", sortedToVisit)

loop:
	// go through each travel
	for i, curr := range travels {
		// Special case for first
		// We can go either direction after choosing our first shortest path
		if len(visited) == 0 {
			visited = append(visited, curr.src)
			visited = append(visited, curr.dest)
			distance += curr.dist
			// So we have to account for both possible directions for the second travel
		} else if len(visited) == 2 {
			// Check if current location contains one of our last visited location
			// If so, that's our next stop
			if slices.Contains(visited, curr.src) && !slices.Contains(visited, curr.dest) {
				visited = append(visited, curr.dest)
				distance += curr.dist
				goto loop
			} else if slices.Contains(visited, curr.dest) && !slices.Contains(visited, curr.src) {
				visited = append(visited, curr.src)
				distance += curr.dist
				goto loop
			}

			// Afterwards, check if either ends of our visited locations is a part of current location and that the associated destination haven't already been visited
			// The goal here is to create a path from one end to another of the list of locations we need to visit
		} else {
			if visited[len(visited)-1] == curr.src && !slices.Contains(visited, curr.dest) {
				visited = append(visited, curr.dest)
				distance += curr.dist
				goto loop
			} else if visited[len(visited)-1] == curr.dest && !slices.Contains(visited, curr.src) {
				visited = append(visited, curr.src)
				distance += curr.dist
				goto loop
			} else if visited[0] == curr.src && !slices.Contains(visited, curr.dest) {
				visited = append([]string{curr.dest}, visited...)
				distance += curr.dist
				goto loop
			} else if visited[0] == curr.dest && !slices.Contains(visited, curr.src) {
				visited = append([]string{curr.src}, visited...)
				distance += curr.dist
				goto loop
			}
		}

		fmt.Println("Line", i, "dist", distance, "visited", visited)

		// If we have visited all locations, break loop early
		sortedVisited := make([]string, len(visited))
		copy(sortedVisited, visited)
		sort.Slice(sortedVisited, func(i, j int) bool { return sortedVisited[i] < sortedVisited[j] })
		if slices.Compare(sortedToVisit, sortedVisited) == 0 {
			break
		}
	}

	return distance
}

func part1() {
	input := input()

	// Parse input into a slice of Travel structs
	parsedInput := parse(input)
	// Get a slice of all unique locations Santa needs to visit

	// Sort our parsed input base on travel distance, from shortest to longest
	sortedInput := make([]Travel, len(parsedInput))
	copy(sortedInput, parsedInput)
	sort.Slice(sortedInput, func(i, j int) bool { return sortedInput[i].dist < sortedInput[j].dist })

	// Get shortest overall travel distance by going through our shortest paths and visiting each locations
	shortestDistance := shortest(sortedInput)

	// fmt.Println("Original input :")
	// fmt.Println(input)
	fmt.Println("Sorted input :")
	fmt.Println(sortedInput)
	fmt.Println("Shortest distance is :")
	fmt.Println(shortestDistance)
}

func main() {
	part1()
}
