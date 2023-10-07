package main

import (
	"AoC/15/utils"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// All time units are in seconds

type Reindeer struct {
	Name     string
	Speed    int
	Time     int
	Rest     int
	Points   int
	Traveled int
}

func parse(s []string) []Reindeer {
	res := make([]Reindeer, len(s))

	for i, curr := range s {
		split := strings.Split(curr, " ")

		speed, err := strconv.Atoi(split[3])
		if err != nil {
			panic(err)
		}
		time, err := strconv.Atoi(split[6])
		if err != nil {
			panic(err)
		}
		rest, err := strconv.Atoi(split[13])
		if err != nil {
			panic(err)
		}

		res[i] = Reindeer{Name: split[0], Speed: speed, Time: time, Rest: rest}
	}

	return res
}

// Calculate the max distance a Reindeer can travel, given a certain amount of time
func distance(r Reindeer, t int) int {
	type Sprint struct {
		Distance int
		Time     int
		Number   int
	}

	sprint := Sprint{Distance: r.Speed * r.Time, Time: r.Time + r.Rest}
	sprint.Number = t / sprint.Time

	timeLeft := t - sprint.Time*sprint.Number

	var finalSprint int

	if timeLeft > r.Time {
		finalSprint = sprint.Distance
	} else {
		finalSprint = r.Speed * timeLeft
	}

	return sprint.Distance*sprint.Number + finalSprint
}

// Give 1 point to the reindeer furthest ahead
// If multiple, they all get 1 point
func givePoints(s []Reindeer) []Reindeer {
	highest := 0

	for i := range s {
		if i == 0 {
			highest = s[i].Traveled
			s[i].Points++
		} else if s[i].Traveled >= highest {
			s[i].Points++
		} else {
			break
		}
	}
	return s
}

func best_part1(s []Reindeer, t int) int {
	var ret int

	for _, curr := range s {
		dist := distance(curr, t)
		if dist > ret {
			ret = dist
		}
	}

	return ret
}

func best_part2(s []Reindeer, t int) int {
	for i := 1; i <= t; i++ {
		for j := range s {
			s[j].Traveled = distance(s[j], i)
		}
		// Sort slice by distance traveled
		slices.SortFunc(s, func(a, b Reindeer) int { return cmp.Compare(b.Traveled, a.Traveled) })

		s = givePoints(s)
	}

	// Sort slice by number of points
	slices.SortFunc(s, func(a, b Reindeer) int { return cmp.Compare(b.Points, a.Points) })

	return s[0].Points
}

func part1() {
	parsed := parse(utils.Input("input.txt"))
	dist := best_part1(parsed, 2503)

	fmt.Println("Part 1")
	fmt.Println("Longest traveled distance is :", dist)
}

func part2() {
	parsed := parse(utils.Input("input.txt"))
	points := best_part2(parsed, 2503)

	fmt.Println("----------------")
	fmt.Println("Part 2")
	fmt.Println("Number of points of the winning reindeer :", points)
}

func main() {
	part1()
	part2()

}
