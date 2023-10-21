package main

import (
	"fmt"
	"strings"

	"AoC/15/utils"
)

var board [][]bool

func parse(s []string) [][]bool {
	var ret [][]bool
	for y, l := range s {
		split := strings.Split(l, "")
		tmp := []bool{}
		for x, curr := range split {
			b := false
			if curr == "#" ||
				// Part 2 : light up the four corners during initial parsing
				(x == 0 && y == 0) ||
				(x == 0 && y == len(split[0])-1) ||
				(x == len(split)-1 && y == 0) ||
				(x == len(s)-1 && y == len(split)-1) {
				b = true
			}
			tmp = append(tmp, b)
		}
		ret = append(ret, tmp)
	}
	return ret
}

func check(x int, y int, s [][]bool) bool {
	// Part 2 : Make sure the four corners stay on during every step
	if (x == 0 && y == 0) || (x == len(s)-1 && y == 0) || (x == 0 && y == len(s[0])-1) || (x == len(s)-1 && y == len(s[0])-1) {
		return true
	}

	lit := 0
	// 	Count number of lit grid around curr
	for yy := -1; yy <= 1; yy++ {
		// Skip if we start on the top or bottom line of our grid
		if (y == 0 && yy == -1) || (y == len(s)-1 && yy == 1) {
			continue
		}
		for xx := -1; xx <= 1; xx++ {
			// Skip if we start on leftmost or rightmost column of our grid OR when we get to YY,XX == 0,0 since it's current pos
			if (x == 0 && xx == -1) || (x == len(s[0])-1 && xx == 1) || (xx == 0 && yy == 0) {
				continue
			}

			if s[y+yy][x+xx] {
				lit++
			}
		}
	}

	if lit == 3 || (s[y][x] && lit == 2) {
		return true
	}
	return false
}

func step(s [][]bool) [][]bool {
	// Deep copy
	ret := make([][]bool, len(s))
	for y := range s {
		ret[y] = make([]bool, len(s[y]))
		copy(ret[y], s[y])
	}

	// Go through every grid and get its next state
	for y := range s {
		for x := range s[y] {
			ret[y][x] = check(x, y, s)
		}
	}
	return ret
}

func count(s [][]bool) int {
	var ret int
	for y := range s {
		for x := range s[y] {
			if s[y][x] {
				ret++
			}
		}
	}
	return ret
}

func part1() {
	board = parse(utils.Input("input.txt"))

	for x := 0; x < 100; x++ {
		board = step(board)
	}
	lit := count(board)

	fmt.Println("Part 1")
	fmt.Println(lit, "lit grid after 100 steps")
	fmt.Println( /* board */ )
}

func part2() {
	fmt.Println("----------------")
	fmt.Println("Part 2")
	fmt.Println("World")
}

func main() {
	part1()
	part2()
}
