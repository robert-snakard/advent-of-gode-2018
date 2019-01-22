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

	// Turn input text into array - if file ends in newline remove empty string
	boxids := strings.Split(string(dat), "\n")
	if boxids[len(boxids)-1] == "" {
		boxids = boxids[:len(boxids)-1]
	}

	fmt.Println("Part 1 Solution:", part_1(boxids))
}

func part_1(boxids []string) int {
	two_count := 0
	three_count := 0

	for _, id := range boxids {
		fmt.Println("ID:", id)
		var letter_count [26]int
		for _, letter := range id {
			letter_count[int(letter)-int('a')]++
		}

		two_flag := false
		three_flag := false
		for _, count := range letter_count {
			if !three_flag && count == 3 {
				three_count++
				three_flag = true
			} else if !two_flag && count == 2 {
				two_count++
				two_flag = true
			}

			if two_flag && three_flag {
				break
			}
		}
	}

	fmt.Println("two:", two_count, "three:", three_count)
	return three_count * two_count
}
