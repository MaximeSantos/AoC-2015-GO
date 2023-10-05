package main

import "testing"

func TestIncrement1(t *testing.T) {
	input := "aaaaaaaa"
	expected := "aaaaaaab"

	test := increment(input)

	if test != expected {
		t.Fatal("Error, expected output not met for increment function")
	}
}

func TestIncrement2(t *testing.T) {
	input := "zzzzzzzz"
	expected := "aaaaaaaa"

	test := increment(input)

	if test != expected {
		t.Fatal("Error, expected output not met for increment function")
	}
}

func TestIncrement3(t *testing.T) {
	input := "hijklmmn"
	expected := "hijklmmo"

	test := increment(input)

	if test != expected {
		t.Fatal("Error, expected output not met for increment function")
	}
}
func TestIncrement4(t *testing.T) {
	input := "bbbbaazz"
	expected := "bbbbabaa"

	test := increment(input)

	if test != expected {
		t.Fatal("Error, expected output not met for increment function")
	}
}
