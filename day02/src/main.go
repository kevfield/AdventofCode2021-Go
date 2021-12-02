package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// get file location from user
	fileLoc, _ := inputFlags()

	puzzleInput, _ := readFile(fileLoc)

	//calculateDepth(puzzleInput)
	fmt.Println("P1 Answer", calculateDepth(puzzleInput))

	fmt.Println("P2 Answer", calculateAim(puzzleInput))

}

func calculateDepth(cdInput []string) int {
	var finalPosition, horizontalPos, depthPos, currdirectionamountInt int

	for i := 0; i < len(cdInput); i++ {

		moveCommand := strings.Split(cdInput[i], " ")

		if moveCommand[0] == "forward" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			horizontalPos = horizontalPos + currdirectionamountInt
		} else if moveCommand[0] == "down" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			depthPos = depthPos + currdirectionamountInt
		} else if moveCommand[0] == "up" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			depthPos = depthPos - currdirectionamountInt
		}
	}
	finalPosition = depthPos * horizontalPos

	return finalPosition
}

func calculateAim(cdInput []string) int {
	var finalPosition, horizontalPos, depthPos, aim, currdirectionamountInt int

	for i := 0; i < len(cdInput); i++ {

		moveCommand := strings.Split(cdInput[i], " ")

		if moveCommand[0] == "forward" {
			// forward X does two things:
			// It increases your horizontal position by X units.
			// It increases your depth by your aim multiplied by X.
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			horizontalPos = horizontalPos + currdirectionamountInt
			depthPos = depthPos + aim*currdirectionamountInt
		} else if moveCommand[0] == "down" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			aim = aim + currdirectionamountInt
		} else if moveCommand[0] == "up" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			aim = aim - currdirectionamountInt
		}
	}
	finalPosition = depthPos * horizontalPos

	return finalPosition
}
