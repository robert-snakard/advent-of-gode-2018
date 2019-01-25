package main

import (
	"errors"
	"fmt"
	"io/ioutils"
	"strings"
	"strconv"

	"../shared"
)

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check.Check(err)

	data := strings.Split(string(dat), "\n")
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}

	fmt.Println("Part 1 Solution:", part_1(data))
}

func part_1(data []string) int{
	return 4
}
