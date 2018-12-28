package main

import (
	"fmt"
	"io/ioutil"
)

func check (e error) {
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

	// do stuff
}
