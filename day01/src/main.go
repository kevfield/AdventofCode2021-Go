package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// get file location from user
	fileLoc, _ := inputFlags()

	puzzleInput, _ := readFile(fileLoc)

	p1Answer := higherorLower(puzzleInput)

	fmt.Println("p1 Answer", p1Answer)

	p2Answer := sumofThree(puzzleInput)
	fmt.Println("p2 answer", p2Answer)

}

// read a file from an input and return into a slice of strings
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//lines = append(lines, "")
	return lines, scanner.Err()
}

// use flags to get user input
func inputFlags() (string, string) {

	// variables declaration
	var flagfile string
	var partid string

	// flags declaration using flag package
	flag.StringVar(&flagfile, "file", "input.txt", "filename of the input data eg: input.txt")
	flag.StringVar(&partid, "part", "a", "use either a or b")

	flag.Parse() // after declaring flags we need to call it
	return flagfile, partid
}

func higherorLower(holInput []string) int {

	var higherFound int
	var valOne int
	var valTwo int

	for i := 0; i < len(holInput)-1; i++ {
		valOne, _ = strconv.Atoi(holInput[i+1])
		valTwo, _ = strconv.Atoi(holInput[i])
		if valOne > valTwo {
			higherFound++
		}
	}

	return higherFound
}

func sumofThree(sotInput []string) int {

	lengthofSlice := len(sotInput)
	noFurther := lengthofSlice - 2
	var totalFound int
	var previousSum int
	var currentSum int

	for i := 0; i < len(sotInput); i++ {
		if i < noFurther {
			valOne, _ := strconv.Atoi(sotInput[i])
			valTwo, _ := strconv.Atoi(sotInput[i+1])
			valThree, _ := strconv.Atoi(sotInput[i+2])

			currentSum = valOne + valTwo + valThree
		}
		if i != 0 {
			if currentSum > previousSum {
				totalFound++
			}

		} else {
			previousSum = currentSum
		}
		previousSum = currentSum
	}

	return totalFound
}
