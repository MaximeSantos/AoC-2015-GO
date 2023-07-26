package main

func parseInts(box string) Box {
	currentBox := Box{}
	// point to the current Box value, starting with length
	currentValue := &currentBox.length

	// Iterate through each character
	for _, char := range box {
		// Check if character is x and if it is, point to next value
		if string(char) == "x" {
			switch currentValue {
			case &currentBox.length:
				currentValue = &currentBox.width
			case &currentBox.width:
				currentValue = &currentBox.height
			}
			continue
		}
		// Current character get added to currentValue
		// We *10 to get to skip to append the current digit to the right
		// int(char - '0') allows us to get an int of the Unicode representation of the rune
		// we assume here that the input is sanitized and only has digits & x
		*currentValue = *currentValue*10 + int(char-'0')
	}
	// Print & return each box
	// fmt.Printf("Length: %d, Width: %d, Height: %d\n", currentBox.length, currentBox.width, currentBox.height)
	return currentBox
}

func getSmallestFace(box Box) int {
	face1 := box.length * box.width
	face2 := box.width * box.height
	face3 := box.height * box.length

	// Slice of our faces
	s := []int{face1, face2, face3}

	// We assume face1 is smallest as a starter
	smallestFace := face1

	// Go through each value of our box and set it as smallestFace if smaller than previous one
	for _, v := range s {
		if v < smallestFace {
			smallestFace = v
		}
	}
	return smallestFace
}

func getTwoSmallestSides(box Box) [2]int {
	// We assume the length & width are smallest by default
	smallestSides := [2]int{box.length, box.width}

	// if box height is longer than both length & width, then skip, else
	if box.height > box.length && box.height > box.width {
	} else {
		// change the value of longest side (length or width) by the value of height
		if box.length < box.width {
			smallestSides[1] = box.height
		} else {
			smallestSides[0] = box.height
		}
	}
	return smallestSides
}
