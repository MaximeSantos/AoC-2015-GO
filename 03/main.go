package main

import "fmt"

type Coordinates struct {
	x int
	y int
}

// Returns true if value exists in slice ; else return false
func includes(slice []Coordinates, value Coordinates) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Change the location location given according to the direction & add it to list of visited locations if it is a new location
func changeLocation(pVisitedCoordinates *[]Coordinates, cLocation *Coordinates, direction rune) {
	switch string(direction) {
	case "^":
		cLocation.y++
	case "v":
		cLocation.y--
	case ">":
		cLocation.x++
	case "<":
		cLocation.x--
	}

	if !includes(*pVisitedCoordinates, *cLocation) {
		*pVisitedCoordinates = append(*pVisitedCoordinates, *cLocation)
	}
}

func main() {
	// Santa & Robo's locations
	santaLocation, roboLocation := Coordinates{}, Coordinates{}
	isSantaTurnToMove := true
	// Slice of all uniquely visited locations
	visitedCoordinates := []Coordinates{}

	// For each input, change current location of correct character
	// Then if it is a new location, add it to slice of visited location
	for _, c := range input {
		if isSantaTurnToMove {
			changeLocation(&visitedCoordinates, &santaLocation, c)
		} else {
			changeLocation(&visitedCoordinates, &roboLocation, c)
		}
		isSantaTurnToMove = !isSantaTurnToMove
	}

	fmt.Printf("Visited coordinates : %v\n", visitedCoordinates)
	fmt.Printf("Number of unique visited houses : %v\n", len(visitedCoordinates))
}

//

// --- Day 3: Perfectly Spherical Houses in a Vacuum ---

// Santa is delivering presents to an infinite two-dimensional grid of houses.

// He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

// However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

// For example:

//     > delivers presents to 2 houses: one at the starting location, and one to the east.
//     ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
//     ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

// ! 2565

// --- Part Two ---

// The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

// Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

// This year, how many houses receive at least one present?

// For example:

//     ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
//     ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
//     ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.

// ! 2639
