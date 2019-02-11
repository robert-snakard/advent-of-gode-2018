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
	sleep_intervals, err := GetIntervalsFromEvents(events)
	check.Check(err)


	fmt.Println("Part 1 Solution:", part_1(sleep_intervals))
	fmt.Println("Part 2 Solution:", part_2(sleep_intervals))
}

func part_1(sleep_intervals map[int][]Interval) int{
	sleep_sum := make(map[int]int)
	sleep_counters := make(map[int]*[60]int)
	// Initialize sleep_counters. Have to use pointers b/c Go's maps don't play nicely when
	// holding objects that are larger than 'size_t' bytes. The other strategy would be to
	// use a temp variable on each assignment but this becomes more inefficient the more times
	// we access each object
	for key, _ := range sleep_intervals {
		cur_id := key
		sleep_counters[cur_id] = &[60]int{}
	}

	var max_sum, max_id int
	// Find guard that's asleep the longest (max_id) and which minute he's asleep the longest for (sleep_counters)
	for key, val := range sleep_intervals {
		cur_id := key
		for _, interval := range val {
			for i := interval.start; i < interval.end; i++ {
				sleep_sum[cur_id]++
				(*sleep_counters[cur_id])[i]++

				if max_sum < sleep_sum[cur_id] {
					max_id = cur_id
					max_sum = sleep_sum[cur_id]
				}

			}
		}
	}

	var max_idx int
	// Where max_idx is the minute the guard most commonly fell asleep at
	for idx, minutes := range sleep_counters[max_id] {
		if minutes > sleep_counters[max_id][max_idx] {
			max_idx = idx
		}
	}

	return max_id * max_idx
}

func part_2(sleep_intervals map[int][]Interval) int{
	sleep_counters := make(map[int]*[60]int)
	// Initialize sleep_counters. Have to use pointers b/c Go's maps don't play nicely when
	// holding objects that are larger than 'size_t' bytes. The other strategy would be to
	// use a temp variable on each assignment but this becomes more inefficient the more times
	// we access each object
	for key, _ := range sleep_intervals {
		cur_id := key
		sleep_counters[cur_id] = &[60]int{}
	}

	var max_val int
	var max_minute int
	var max_id int
	// find the minute, guard combo where sleep happens most often
	for key, val := range sleep_intervals {
		cur_id := key
		for _, interval := range val {
			for i := interval.start; i < interval.end; i++ {
				(*sleep_counters[cur_id])[i]++
				if (*sleep_counters[cur_id])[i] > max_val {
					max_val = (*sleep_counters[cur_id])[i]
					max_minute = i
					max_id = cur_id
				}
			}
		}
	}

	return max_id * max_minute
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
	return events, nil
}

func GetIntervalsFromEvents(events []Event) (map[int][]Interval, error){
	// Key is guards id, array is all the time intervals that guard fell asleep
	sleep_intervals := make(map[int][]Interval)

	cur_id := -1
	var cur_state State
	// Fill sleep intervals map
	for _, evt := range events {
		if evt.state == NewGuard {
			cur_id = evt.guard_id
			cur_state = NewGuard
			continue
		}

		if cur_id < 0 {
			return make(map[int][]Interval), errors.New("invalid guard id")
		}

		// If the guard falls asleep add a new interval. If he wakes up complete the previous interval
		switch evt.state {
			case FallsAsleep:
				sleep_intervals[cur_id] = append(sleep_intervals[cur_id], Interval{evt.time.Minute(), 0})
				cur_state = FallsAsleep
			case WakesUp:
				if cur_state != FallsAsleep {
					return make(map[int][]Interval), errors.New("woke up without falling asleep first")
				}
				sleep_intervals[cur_id][len(sleep_intervals[cur_id])-1].end = evt.time.Minute()
			default:
				return make(map[int][]Interval), errors.New("invalid state")
		}
	}
	return sleep_intervals, nil
}

// Helper Functions for sorting
type byDateTime []Event
func (s byDateTime) Len() int {
	return len(s)
}
func (s byDateTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byDateTime) Less(i, j int) bool {
	return s[i].time.Before(s[j].time)
}
