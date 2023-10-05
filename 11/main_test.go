package main

import "testing"

func TestLookForPassword(t *testing.T) {
	input := "abcdefgh"
	expected := "abcdffaa"

	test, err := lookForPassword(input)

	if err != nil || test != expected {
		t.Fatal("Error, did not find expected password :", test, err)
	}
}

func TestLookForPassword2(t *testing.T) {
	input := "ghijklmn"
	expected := "ghjaabcc"

	test, err := lookForPassword(input)

	if err != nil || test != expected {
		t.Fatal("Error, did not find expected password :", test, err)
	}
}
