package main

import (
	"fmt"
	"strconv"
)

func main() {

	// get file location
	fileLoc, _ := inputFlags()
	// get puzzle input, add to slice of strings
	puzzleInput, _ := readFile(fileLoc)

	fmt.Println("P1 Answer =", commonBits(puzzleInput))

}

func commonBits(cbInput []string) int {
	// make a map to count the bits
	onebitsMap := make(map[int]int)
	zerobitsMap := make(map[int]int)

	// j is the index
	// character is the rune
	for i := 0; i < len(cbInput); i++ {
		for j := range cbInput[i] {
			if cbInput[i][j] == '1' {
				onebitsMap[j]++
			} else {
				zerobitsMap[j]++
			}
		}
	}

	gammaSlice := make([]byte, len(cbInput[0]))
	epsilonSlice := make([]byte, len(cbInput[0]))
	for i := 0; i < len(cbInput[i]); i++ {
		if onebitsMap[i] > zerobitsMap[i] {
			gammaSlice[i] = '1'
			epsilonSlice[i] = '0'
		} else {
			gammaSlice[i] = '0'
			epsilonSlice[i] = '1'
		}
	}

	gammaCalc, _ := strconv.ParseInt(string(gammaSlice), 2, 64)
	epsilonCalc, _ := strconv.ParseInt(string(epsilonSlice), 2, 64)

	return int(gammaCalc) * int(epsilonCalc)
}
