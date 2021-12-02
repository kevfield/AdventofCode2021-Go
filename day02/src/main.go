package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// get file location
	fileLoc, _ := inputFlags()
	// get puzzle input, add to slice of strings
	puzzleInput, _ := readFile(fileLoc)

	fmt.Println("P1 Answer =", calculateDepth(puzzleInput), "\nP2 Answer =", calculateAim(puzzleInput))
}

func calculateDepth(cdInput []string) int {
	var horizontalPos, depthPos, currdirectionamountInt int

	for i := 0; i < len(cdInput); i++ {

		moveCommand := strings.Split(cdInput[i], " ")

		if moveCommand[0] == "forward" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			horizontalPos += currdirectionamountInt
		} else if moveCommand[0] == "down" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			depthPos += currdirectionamountInt
		} else if moveCommand[0] == "up" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			depthPos -= currdirectionamountInt
		}
	}
	return depthPos * horizontalPos
}

func calculateAim(cdInput []string) int {
	var horizontalPos, depthPos, aim, currdirectionamountInt int

	for i := 0; i < len(cdInput); i++ {

		moveCommand := strings.Split(cdInput[i], " ")

		if moveCommand[0] == "forward" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			horizontalPos += currdirectionamountInt
			depthPos += aim * currdirectionamountInt
		} else if moveCommand[0] == "down" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			aim = aim + currdirectionamountInt
		} else if moveCommand[0] == "up" {
			currdirectionamountInt, _ = strconv.Atoi(moveCommand[1])
			aim -= currdirectionamountInt
		}
	}
	return depthPos * horizontalPos
}
