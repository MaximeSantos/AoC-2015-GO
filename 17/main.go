package main

import (
	"fmt"
	"strconv"

	"AoC/15/utils"
)

func parse(s []string) []int {
	ret := make([]int, 0, len(s))
	for _, curr := range s {
		n, err := strconv.Atoi(curr)
		if err != nil {
			panic(err)
		}
		ret = append(ret, n)
	}
	return ret
}

func checkSum(s []int, tar int) bool {
	total := 0
	for _, n := range s {
		total += n
		if total > tar {
			return false
		}
	}
	if total == tar {
		return true
	}
	return false
}

func sum(s []int) int {
	ret := 0
	for _, n := range s {
		ret += n
	}
	return ret
}

// Recursively create the powerset of given input
func pset(ps, tmp []int) [][]int {
	if len(ps) == 0 {
		return [][]int{tmp}
	}
	ret := [][]int{}
	for _, set := range pset(ps[1:], tmp) {
		ret = append(ret, set)
	}
	for _, set := range pset(ps[1:], append(tmp, ps[0])) {
		ret = append(ret, set)
	}
	return ret
}

func getPowerset(choices []int) [][]int {
	return pset(choices, []int{})
}

func validSets(ps [][]int, tar int) int {
	ret := 0
	for _, c := range ps {
		if sum(c) == tar {
			ret++
		}
	}
	return ret
}

func part1() {
	input := parse(utils.Input("input.txt"))
	ps := getPowerset(input)

	fmt.Println("Part 1")
	fmt.Println("Length of whole powerset:", len(ps))
	fmt.Println("Number of sets that meet our criteria:", validSets(ps, 150))
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

// package main
//
// import (
// 	"fmt"
// )
//
// var (
// 	totals     = map[int]int{}
// 	containers = []int{50,44,11,49,42,46,18,32,26,40,21,7,18,43,10,47,36,24,22,40}
// )
//
// func f(idx, used, n int) {
// 	if used == 150 {
// 		totals[n] = totals[n] + 1
// 	} else if used > 150 {
// 		return
// 	} else if idx >= len(containers) {
// 		return
// 	} else {
// 		f(idx+1, used+containers[idx], n+1)
// 		f(idx+1, used, n)
// 	}
// }
//
// func main() {
// 	f(0, 0, 0)
// 	minK, V, T := len(containers), 0, 0
// 	for k, v := range totals {
// 		if k < minK {
// 			minK = k
// 			V = v
// 		}
// 		T += v
// 	}
// 	fmt.Println("Part1", T)
// 	fmt.Println("Part2", minK, V)
// }
