package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println("Go has encountered an error:", e)
		panic(e)
	}
}

func main() {
	// Read input text
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	// Turn input text into array - if file ends in newline remove empty string.
	calibrations := strings.Split(string(dat), "\n")
	if calibrations[len(calibrations)-1] == "" {
		calibrations = calibrations[:len(calibrations)-1]
	}

	fmt.Println("Part 1 Solution:", part_1(calibrations))
	fmt.Println("Part 2 Solution:", part_2(calibrations))
}

func part_1(calibrations []string) int {
	// Calculate Sum
	sum := 0
	for _, c := range calibrations {
		calibration, err := strconv.Atoi(c)
		check(err)
		sum += calibration
	}
	return sum
}

func part_2(calibrations []string) int {
	sums_set := make(map[int]bool)
	sum := 0
	sums_set[sum] = true

	for true {
		for _, c := range calibrations {
			calibration, err := strconv.Atoi(c)
			check(err)
			sum += calibration

			if _, empty := sums_set[sum]; empty {
				return sum
			}
			sums_set[sum] = true
		}
	}
	return 0 // We'll never get here
}
