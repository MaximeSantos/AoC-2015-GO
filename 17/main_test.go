package main

import "testing"

func TestCheckSum(t *testing.T) {
	expected := true
	input := []int{1, 2, 3, 5}
	test := checkSum(input, 11)

	if test != expected {
		t.Fatal("Error. Expected", expected, "found", test)
	}

	input2 := []int{1, 2, 3, 5, 11, 11}
	test2 := checkSum(input2, 33)

	if test2 != expected {
		t.Fatal("Error. Expected", expected, "found", test2)
	}

	input3 := []int{40, 5, 5, 10, 10}
	test3 := checkSum(input3, 70)

	if test3 != expected {
		t.Fatal("Error. Expected", expected, "found", test3)
	}
}

func TestCombin(t *testing.T) {
	expected := 4
	input := []int{20, 15, 10, 5, 5}
	test := validSets(getPowerset(input), 25)
	if test != expected {
		t.Fatal("Error. Expected", expected, "found", test)
	}
}
