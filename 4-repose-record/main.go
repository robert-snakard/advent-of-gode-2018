package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"strconv"
	"time"

	"../shared"
)

type Event struct {
	time time.Time
	state State
	guard_id int
}

type Interval struct {
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
	events, err := ParseIntoEvents(data)
	check.Check(err)

	fmt.Println("Part 1 Solution:", part_1(events))
}

func part_1(events []Event) int{
	return 4
}

type State int
const (
	NewGuard    State = 0
	FallsAsleep State = 1
	WakesUp     State = 2
)

func ParseIntoEvents(data []string) ([]Event, error){
	var events []Event
	for _, line := range data {
		var err error
		var cur_event Event

		// Line halves if of the form
		//     [0] = date and time
		//     [1] = string that describes state
		line_halves := strings.Split(line, "]")

		time_layout := "[2006-01-02 15:04"
		cur_event.time, err = time.Parse(time_layout, line_halves[0])
		check.Check(err)

		if strings.Contains(line_halves[1], "begins shift") {
			cur_event.state = NewGuard
		} else if strings.Contains(line_halves[1], "falls asleep") {
			cur_event.state = FallsAsleep
		} else if strings.Contains(line_halves[1], "wakes up") {
			cur_event.state = WakesUp
		} else {
			return []Event{}, errors.New(fmt.Sprintf("Error parsing input: Unable to",
			                                         "determine state of %s", line))
		}

		// Split to get id is of the form
		//     [0] - ""
		//     [1] - "Guard"
		//     [2] - "#1234"
		//     [3] - "begins"
		//     [4] - "shift"
		if cur_event.state == NewGuard {
			split_to_get_id := strings.Split(line_halves[1], " ")
			cur_event.guard_id, err = strconv.Atoi(split_to_get_id[2][1:])
			check.Check(err)
		}

		events = append(events, cur_event)
	}

	sort.Sort(byDateTime(events))
	return events nil
}

// Helper Functions for sorting
type byDateTime []ParsedLine
func (s byDateTime) Len() int {
	return len(s)
}
func (s byDateTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byDateTime) Less(i, j int) bool {
	return s[i].time.Before(s[j].time)
}
