package main

import "testing"

func TestCheckForStraightTrue(t *testing.T) {
	input := "efgjaaaa"
	expected := true

	if checkForStraight(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForStraightTrue2(t *testing.T) {
	input := "abdjaabc"
	expected := true

	if checkForStraight(input) != expected {
		t.Fatal("Error when straight is at the end of string")
	}
}

func TestCheckForStraightTrue3(t *testing.T) {
	input := "xyzabdja"
	expected := true

	if checkForStraight(input) != expected {
		t.Fatal("Error with straight xyz")
	}
}

func TestCheckForStraightTrue4(t *testing.T) {
	input := "abdjaxyz"
	expected := true

	if checkForStraight(input) != expected {
		t.Fatal("Error when straight xyz is at the end of string")
	}
}

func TestCheckForStraightFalse(t *testing.T) {
	input := "abdkjkja"
	expected := false

	if checkForStraight(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForForbiddenTrue(t *testing.T) {
	input := "abcdefjk"
	expected := true

	if checkForForbidden(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForForbiddenFalse(t *testing.T) {
	input := "abcdlmnp"
	expected := false

	if checkForForbidden(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForForbiddenFalse2(t *testing.T) {
	input := "abcdimnp"
	expected := false

	if checkForForbidden(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForForbiddenFalse3(t *testing.T) {
	input := "abcdomnp"
	expected := false

	if checkForForbidden(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForDoubleTrue(t *testing.T) {
	input := "abbceffg"
	expected := true

	if checkForDouble(input) != expected {
		t.Fatal("Error with example string")
	}
}

func TestCheckForDoubleTrue2(t *testing.T) {
	input := "aabcdeff"
	expected := true

	if checkForDouble(input) != expected {
		t.Fatal("Error when double is at the end of string")
	}
}

func TestCheckForDoubleTrue3(t *testing.T) {
	input := "ghjaabcc"
	expected := true

	if checkForDouble(input) != expected {
		t.Fatal("Error with example string")
	}
}

func TestCheckForDoubleFalse(t *testing.T) {
	input := "aabcdefj"
	expected := false

	if checkForDouble(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForDoubleFalse2(t *testing.T) {
	input := "abbcegjk"
	expected := false

	if checkForDouble(input) != expected {
		t.Fatal("Error")
	}
}

func TestCheckForStraightAndForbidden(t *testing.T) {
	input := "hijklmmn"
	expected := false

	if (checkForStraight(input) && checkForForbidden(input)) != expected {
		t.Fatal("Error, fail on example string")
	}
}
