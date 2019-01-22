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

/*******
Research for this problem includes the following research paper and stackoverflow question.
This research was for implementing a solution for part 2. that did not take O(n^2k) time.

- Directory Matching nad Indexing with Errors and Don't Cares
	https://cs.nyu.edu/~adi/CGL04.pdf
- Data structure or algorithm for quickly finding differences between strings
	https://cs.stackexchange.com/questions/93467/data-structure-or-algorithm-for-quickly-finding-differences-between-strings

The current solution does not use this research and is a simple solution that doea take O(n^2k) time.
******/

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
	fmt.Println("Part 2 Solution:", part_2(boxids))
}

func part_1(boxids []string) int {
	two_count := 0
	three_count := 0

	for _, id := range boxids {
		// Count letters in box id
		var letter_count [26]int
		for _, letter := range id {
			letter_count[int(letter)-int('a')]++
		}

		// Figure out if id has a double letter and/or a triple letter
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

	return three_count * two_count
}

func part_2(boxids []string) string {
	var id_solution string
	var comp_solution string
	Outer: for _, id := range boxids {
		for _, comparison := range boxids {
			diffs := 0
			for i := 0; i < len(id); i++ {
				if id[i] != comparison[i] {
					diffs++
				}
			}
			if diffs == 1 {
				id_solution = id
				comp_solution = comparison
				break Outer
			}
		}
	}

	final_solution := ""
	for i := 0; i < len(id_solution); i++ {
		if id_solution[i] == comp_solution[i] {
			final_solution += string(id_solution[i])
		}
	}

	return final_solution
}
