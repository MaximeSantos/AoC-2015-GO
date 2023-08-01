package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Open file and get the input by scanning it line by line
func input() [][]string {
	var parsed [][]string

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
		parsed = append(parsed, parse(line))
	}

	return parsed
}

// Parse line into a more usable format : "COMMAND INPUT OUTPUT"
func parse(line string) []string {
	s := strings.Split(line, " ")

	switch len(s) {
	case 3:
		// direct assignment  "187 -> ze"
		return []string{"TO", s[0], s[2]}
	case 4:
		// NOT gate "NOT cn -> co"
		return []string{"NOT", s[1], s[3]}
	case 5:
		// LSHIFT, RISHIFT, AND, OR gates "ez AND oe -> b"
		return []string{s[1], fmt.Sprintf("%s,%s", s[0], s[2]), s[4]}
		// Probably better to return an error with the empty slice | doesnt matter much here since our input is clean
	default:
		return []string{}
	}
}

// Order our parsed slice alphabetically. Single letters first, then double letters.
func order(parsed [][]string) [][]string {
	o := parsed
	sort.Slice(o, func(i, j int) bool {
		if len(o[i][2]) != len(o[j][2]) {
			return len(o[i][2]) < len(o[j][2])
		} else {
			return o[i][2] < o[j][2]
		}
	})
	// Bit of a cheat, but take first instruction to last position in slice since we know it's the instruction for the a wire
	x, o := o[0], o[1:]
	o = append(o, x)

	return o
}

// Go through the circuit by following the ordered instructions
//
// Returns the value of the a wire
func compute(o [][]string) int {
	// Map of our wires - Key is the name of the wire, Value is the 16 bit int signal
	m := make(map[string]int)

	for _, v := range o {
		switch v[0] {
		// In case of a direct assignment...
		case "TO":
			// If input value is a number, use its value directly
			// Else, look for input's value in the map and assign input
			if i, err := strconv.Atoi(v[1]); err == nil {
				m[v[2]] = i
			} else if _, ok := m[v[1]]; ok {
				m[v[2]] = m[v[1]]
			}
		case "NOT":
			// &^ is bitwise NOT in Go
			if i, err := strconv.Atoi(v[1]); err == nil {
				m[v[2]] = ^i
			} else if _, ok := m[v[1]]; ok {
				m[v[2]] = ^m[v[1]]
			}
		case "AND":
			s := signals(v[1], m)
			m[v[2]] = s[0] & s[1]
		case "OR":
			s := signals(v[1], m)
			m[v[2]] = s[0] | s[1]
		case "LSHIFT":
			s := signals(v[1], m)
			m[v[2]] = s[0] << s[1]
		case "RSHIFT":
			s := signals(v[1], m)
			m[v[2]] = s[0] >> s[1]
		}
	}

	if i, ok := m["a"]; ok {
		return i
	} else {
		return 0
	}
}

// Split our input string in two values.
//
// Check for numbers in string or look for value of wire in passed map
func signals(s string, m map[string]int) []int {
	var values = make([]int, 2)
	// handle operation
	slc := strings.Split(s, ",")
	for i, s := range slc {
		if signal, err := strconv.Atoi(s); err == nil {
			values[i] = signal
		} else if _, ok := m[s]; ok {
			values[i] = m[s]
		}
	}
	return values
}

// Take our parsed input and override value assigned to b with value passed in argument
func override(s [][]string, i int) [][]string {
	for _, v := range s {
		if v[2] == "b" {
			v[1] = fmt.Sprint(i)
		}
	}

	return s
}

func part1() int {
	parsed := (input())
	ordered := order(parsed)
	return compute(ordered)
}

func part2(i int) int {
	parsed := (input())
	overriden := override(parsed, i)
	ordered := order(overriden)
	return compute(ordered)
}

func main() {
	a := part1()
	fmt.Println(a)
	fmt.Println("-----------------------------")
	b := part2(a)
	fmt.Println(b)
}

// --- Day 7: Some Assembly Required ---

// This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

// Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

// The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

// For example:

//     123 -> x means that the signal 123 is provided to wire x.
//     x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
//     p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
//     NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.

// Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

// For example, here is a simple circuit:

// 123 -> x
// 456 -> y
// x AND y -> d
// x OR y -> e
// x LSHIFT 2 -> f
// y RSHIFT 2 -> g
// NOT x -> h
// NOT y -> i

// After it is run, these are the signals on the wires:

// d: 72
// e: 507
// f: 492
// g: 114
// h: 65412
// i: 65079
// x: 123
// y: 456

// In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?

// ! Part 1 : 16076

// --- Part Two ---

// Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?

// ! Part 2 : 2797
