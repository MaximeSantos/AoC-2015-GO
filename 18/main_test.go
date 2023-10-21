package main

import (
	"fmt"
	"testing"

	"AoC/15/utils"
)

// Prints the 2D slice in a more readable format
func printA(s [][]bool) {
	fmt.Println("Lit :", count(s))
	tmp := [][]string{}
	for y := range s {
		line := []string{}
		for _, curr := range s[y] {
			if curr {
				line = append(line, "#")
			} else {
				line = append(line, ".")
			}
		}
		tmp = append(tmp, line)
	}
	for _, curr := range tmp {
		fmt.Println(curr)
	}
	fmt.Println("---------------------------")
}

func TestMain(t *testing.T) {
	expected := 17
	tBoard := parse(utils.Input("test.txt"))
	printA(tBoard)
	for x := 0; x < 5; x++ {
		tBoard = step(tBoard)
		printA(tBoard)
	}
	test := count(tBoard)

	if test != expected {
		t.Fatal("Error. Expected", expected, "found", test)
	}
}
