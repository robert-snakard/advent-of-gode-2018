package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

	data := strings.Split(string(dat), "\n")
	// From this point forward I've edited all input.txts so that it never ends in a newline

	fmt.Println("Part 1 Solution:", part_1(data))
}

func part_1(dat []string) int{
	return 3
}
