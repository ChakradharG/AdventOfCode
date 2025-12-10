package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

type Machine struct {
	state   uint16
	toggles []uint16
	// jolts   []uint16
}

func NewMachine(line string) *Machine {
	m := Machine{
		parseState(line),
		parseToggles(line),
		// parseJolts(line),
	}
	return &m
}

func parseState(line string) uint16 {
	r, _ := regexp.Compile(`[#\.]+`)
	state := uint16(0)
	for b, c := range r.FindString(line) {
		if c == '#' {
			state |= 0b1 << b
		}
	}
	return state
}

func parseToggles(line string) []uint16 {
	r, _ := regexp.Compile(`(\((\d+,)*\d+\))+`)
	matches := r.FindAllString(line, -1)
	toggles := make([]uint16, len(matches))
	for i, match := range matches {
		toggle := uint16(0)
		for c := range strings.SplitSeq(match[1:len(match)-1], ",") {
			b, _ := strconv.Atoi(c)
			toggle |= 0b1 << b
		}
		toggles[i] = toggle
	}
	return toggles
}

// func parseJolts(line string) {

// }

func bfs(machine Machine) uint64 {
	q, enqd := llq.New(), map[uint16]bool{}
	q.Enqueue(machine.state)
	enqd[machine.state] = true
	cnt := uint64(0)
	for q.Size() > 0 {
		for i := q.Size(); i > 0; i-- {
			st, _ := q.Dequeue()
			state := st.(uint16)
			if state == 0 {
				return cnt
			}
			for _, toggle := range machine.toggles {
				next := state ^ toggle
				_, ok := enqd[next]
				if !ok {
					q.Enqueue(next)
					enqd[next] = true
				}
			}
		}
		cnt++
	}
	return cnt
}

func part1(machines []Machine) uint64 {
	ans := uint64(0)
	for _, machine := range machines {
		ans += bfs(machine)
	}
	return ans
}

// func part2(machines []Machine) uint64 {
// 	return uint64(0)
// }

func main() {
	inp, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	machines := []Machine{}
	for scanner.Scan() {
		machines = append(machines, *NewMachine(scanner.Text()))
	}

	fmt.Println(part1(machines))
	// fmt.Println(part2(machines))
}
