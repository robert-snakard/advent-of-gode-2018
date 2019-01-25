package main

import (
//	"errors"
	"fmt"
	"io/ioutil"
	"strings"
//	"strconv"

	"../shared"
)

type interval struct {
	start int
	end int
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check.Check(err)

	data := strings.Split(string(dat), "\n")
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}
	sleep_intervals, err := ParseData(data)
	check.Check(err)

	fmt.Println("Part 1 Solution:", part_1(sleep_intervals))
}

func part_1(sleep_intervals map[int][]interval) int{
	return 4
}

func ParseData(data []string) (map[int][]interval, error){
	sleep_intervals := make(map[int][]interval)
	sleep_intervals[1] = []interval{}

	// Use old style for loop so we can increment i inside the loop
	for i := 0; i < 5; {
		fmt.Println("looping:", i)
		i++
	}

	return sleep_intervals, nil
}
