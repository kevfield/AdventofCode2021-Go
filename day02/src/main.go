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
