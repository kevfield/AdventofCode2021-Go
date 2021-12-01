package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	// get file location from user
	fileLoc, _ := inputFlags()

	puzzleInput, _ := readFile(fileLoc)

	p1Answer := higherorLower(puzzleInput)

	fmt.Println("p1 Answer", p1Answer)

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

func higherorLower(input []string) int {

	var higherFound int
	// run loop with -2 to catch edge case and avoid out of range
	for i := 0; i < len(input)-2; i++ {
		if input[i+1] > input[i] {
			higherFound++
		}
	}

	// edge case
	sliceLength := len(input)
	if input[sliceLength-1] > input[sliceLength-2] {
		higherFound++
	}

	return higherFound
}
