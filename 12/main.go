package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Open the input.json and return the sum of all numbers
func input() interface{} {
	var res interface{}

	// Open file
	file, err := os.Open("input.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Convert file to a Go interface
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(byteValue), &res)

	return res
}

func findAndAdd(m interface{}) int {
	var sum int

	if n, ok := m.(float64); ok {
		sum += int(n)
	} else if arr, ok := m.([]interface{}); ok {
		for _, v := range arr {
			sum += findAndAdd(v)
		}
	} else if mm, ok := m.(map[string]interface{}); ok {
		for _, v := range mm {
			// Part 2 : discard the object & its children if it contains the string "red"
			if s, ok := v.(string); ok && s == "red" {
				return 0
			}
			sum += findAndAdd(v)
		}
	}

	return sum
}

func main() {
	input := input()
	fmt.Println(findAndAdd(input))
}
