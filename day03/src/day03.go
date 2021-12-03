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

	p1Answer, oneMap, zeroMap := commonBits(puzzleInput)
	lifeSupport(puzzleInput, oneMap, zeroMap)

	fmt.Println("P1 Answer =", p1Answer) //, "\nP2 Answer  =", p2Answer)

}

func commonBits(cbInput []string) (int, map[int]int, map[int]int) {
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

	return int(gammaCalc) * int(epsilonCalc), onebitsMap, zerobitsMap
}

func lifeSupport(lsInput []string, lsoneMap map[int]int, lszeroMap map[int]int) {
	oxygenSlice := make([]string, 0)
	oxygentempSlice := make([]string, 0)
	co2Slice := make([]string, 0)
	co2tempSlice := make([]string, 0)
	//keepLooping := true

	// loop through runes of string
	for i := 0; i < len(lsInput[0]); i++ {
		// check if our slice has one entry
		if len(oxygenSlice) != 1 {
			// loop through same rune on each line of input
			for j := 0; j < len(lsInput); j++ {
				// check map for most common
				if lsoneMap[j] >= lszeroMap[j] {
					if lsInput[j][0] == '1' {
						oxygentempSlice = append(oxygentempSlice, lsInput[j])
					}
				} else {
					if lsInput[j][0] == '0' {
						oxygentempSlice = append(oxygentempSlice, lsInput[j])
					}
				}
			}

		}
		fmt.Println("oxslice", oxygenSlice)
		oxygenSlice = append(oxygenSlice, oxygentempSlice[0])
		oxygentempSlice = nil
	}

	// loop through runes of string
	for i := 0; i < len(lsInput[0]); i++ {
		// check if our slice has one entry
		if len(co2Slice) != 1 {
			// loop through same rune on each line of input
			for j := 0; j < len(lsInput); j++ {
				// check map for most common
				if lszeroMap[j] >= lsoneMap[j] {
					if lsInput[j][0] == '0' {
						co2tempSlice = append(co2tempSlice, lsInput[j])
					}
				} else {
					if lsInput[j][0] == '1' {
						co2tempSlice = append(co2tempSlice, lsInput[j])
					}
				}
			}

		}
		fmt.Println("co2slice", co2Slice)
		co2Slice = append(co2Slice, co2tempSlice[0])
		co2tempSlice = nil
	}

	//oxygenCalc, _ := strconv.Atoi(oxygenSlice[0])
	//co2Calc, _ := strconv.Atoi(co2Slice[0])

	//return oxygenCalc * co2Calc
}
