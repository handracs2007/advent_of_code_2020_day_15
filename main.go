// Solution to Advent of Code 2020 day 15: https://adventofcode.com/2020/day/15/
package main

import (
	"fmt"
)

func main() {

	input := []int{1, 20, 8, 12, 0, 14}
	maxTurn := 30000000
	posMap := make(map[int][2]int)
	lastNumber := -1

	for currentTurn := 1; currentTurn <= maxTurn; currentTurn++ {
		if currentTurn <= len(input) {
			// Get the number from the input list
			number := input[currentTurn-1]

			// Build the positions string.
			regPos, _ := posMap[number]
			regPos[0], regPos[1] = currentTurn, currentTurn

			// Sets the positions string of a number.
			posMap[number] = regPos
			lastNumber = number

			continue
		}

		// We do not have more input, now we will process based on the registered lastNumber value.
		// Get the positions string of the last number.
		pos := posMap[lastNumber]

		// Gets the new number based on the given scenario. Set it straight away as the last number.
		lastNumber = pos[1] - pos[0]

		// Build the positions string.
		regPos, ok := posMap[lastNumber]
		if ok {
			regPos[0] = regPos[1]
		}
		regPos[1] = currentTurn
		if !ok {
			// If first number, we make both values the same.
			regPos[0] = regPos[1]
		}

		// Sets the positions string of a number.
		posMap[lastNumber] = regPos
	}

	//fmt.Println(posMap)
	fmt.Println(lastNumber)
}
