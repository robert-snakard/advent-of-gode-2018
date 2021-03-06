package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"

	"../shared"
)

type claim struct {
	id int
	loff int
	toff int
	width int
	height int
}

func main() {
	// Read input text
	dat, err := ioutil.ReadFile("input.txt")
	check.Check(err)

	// Turn input text into array of claim structs
	data := strings.Split(string(dat), "\n")
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}
	claims, err := ParseData(data)
	check.Check(err)

	fmt.Println("Part 1 Solution:", part_1(claims))
	fmt.Println("Part 2 Solution:", part_2(claims))
}

func part_1(claims []claim) int{
	var fabric [1000][1000]int

	// Fill fabric array with claims
	for _, claim := range claims {
		for i := claim.toff; i < claim.toff + claim.height; i++ {
			for j := claim.loff; j < claim.loff + claim.width; j++ {
				fabric[i][j]++
			}
		}
	}

	// count # of square inches covered by more than one claim
	var inch_count int
	for _, row := range fabric {
		for _, claims_on_sqr_in := range row {
			if claims_on_sqr_in > 1 {
				inch_count++
			}
		}
	}

	return inch_count
}

func part_2(claims []claim) claim{
	var fabric [1000][1000]int

	// Label the sq_in as -1 if it contains overlapping claims
	for _, claim := range claims {
		for i := claim.toff; i < claim.toff + claim.height; i++ {
			for j := claim.loff; j < claim.loff + claim.width; j++ {
				if fabric[i][j] == 0 {
					fabric[i][j] = claim.id
				} else {
					fabric[i][j] = -1
				}
			}
		}
	}

	// For each claim, check that no sq_in is overlapping (set to -1)
	ClaimLoop: for _, claim := range claims {
		for i := claim.toff; i < claim.toff + claim.height; i++ {
			for j := claim.loff; j < claim.loff + claim.width; j++ {
				if fabric[i][j] == -1 {
					continue ClaimLoop
				}
			}
		}

		// We'll only get here if we don't continue (the claim does not ever overlap)
		return claim
	}

	return claim{}
}

// The function ParseData assumes that the input is all ascii and nothing has to be
// done to account for utf8's byte/rune discrepancy
func ParseData(data []string) ([]claim, error) {
	var claims []claim
	for idx, line := range data {
		var cur_claim claim
		var err error

		// split_by_spaces if of the form
		//     [0] = "#<id>"
		//     [1] = "@"
		//     [2] = "loff,toff:"
		//     [3] = "widthxheight"
		split_by_spaces := strings.Split(line, " ")

		if !strings.HasPrefix(split_by_spaces[0], "#") {
			return []claim{}, errors.New(fmt.Sprintf("idx %d does not start with #", idx))
		}
		cur_claim.id, err = strconv.Atoi(split_by_spaces[0][1:])
		check.Check(err)

		if split_by_spaces[1] != "@" {
			return []claim{}, errors.New(fmt.Sprintf("idx %d does not contain an '@'", idx))
		}

		offsets := strings.Split(split_by_spaces[2][:len(split_by_spaces[2])-1], ",")
		if len(offsets) != 2 {
			return []claim{}, errors.New(fmt.Sprintf("idx %d has its offsets formatted incorrectly - %+v", idx, offsets))
		}
		cur_claim.loff, err = strconv.Atoi(offsets[0])
		check.Check(err)
		cur_claim.toff, err = strconv.Atoi(offsets[1])
		check.Check(err)

		dimensions := strings.Split(split_by_spaces[3], "x")
		if len(dimensions) != 2 {
			return []claim{}, errors.New(fmt.Sprintf("idx %d has its dimensions formatted incorrectly - %+v", idx, dimensions))
		}
		cur_claim.width, err = strconv.Atoi(dimensions[0])
		check.Check(err)
		cur_claim.height, err = strconv.Atoi(dimensions[1])
		check.Check(err)

		claims = append(claims, cur_claim)
	}

	return claims, nil
}
