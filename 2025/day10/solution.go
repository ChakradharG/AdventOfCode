package day10

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

type Joltage struct {
	lo uint64
	hi uint64
}

func (j *Joltage) Set(x uint64, b int) {
	if b < 5 {
		j.lo |= x << (10 * b)
	} else {
		j.hi |= x << (10 * (b - 5))
	}
}

func apply(jolts Joltage, toggle uint16) (Joltage, bool) {
	mask := uint64(0b1111111111)
	for b := range 10 {
		if (toggle>>b)&0b1 != 0 {
			if b < 5 {
				b10 := 10 * b
				if jolts.lo&(mask<<b10) > 0 {
					jolts.lo -= 0b1 << b10
				} else {
					return jolts, false
				}
			} else {
				b10 := 10 * (b - 5)
				if jolts.hi&(mask<<b10) > 0 {
					jolts.hi -= 0b1 << b10
				} else {
					return jolts, false
				}
			}
		}
	}
	return jolts, true
}

type Machine struct {
	state   uint16
	toggles []uint16
	jolts   Joltage
}

func NewMachine(line string) *Machine {
	m := Machine{
		parseState(line),
		parseToggles(line),
		parseJolts(line),
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

func parseJolts(line string) Joltage {
	r, _ := regexp.Compile(`{[\d,]+}`)
	match := r.FindString(line)
	jolts := Joltage{}
	for b, j := range strings.Split(match[1:len(match)-1], ",") {
		x, _ := strconv.Atoi(j)
		jolts.Set(uint64(x), b)
	}
	return jolts
}

func bfs1(machine Machine) uint64 {
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
				if !enqd[next] {
					q.Enqueue(next)
					enqd[next] = true
				}
			}
		}
		cnt++
	}
	return cnt
}

func bfs2(machine Machine) uint64 {
	/*
		This BFS explores the state space of all reachable joltage counter
		configurations. With up to 10 counters and typical joltage values in
		the hundreds, the number of non-negative integer vectors reachable by
		repeated subtraction is exponential in the number of counters. As a
		result, the BFS search space grows exponentially and is not a
		suitable algorithm for Part 2.

		See `2025/10/solution.py` for an ILP-based solution.
	*/
	q, enqd := llq.New(), map[Joltage]bool{}
	q.Enqueue(machine.jolts)
	enqd[machine.jolts] = true
	zero := Joltage{}
	cnt := uint64(0)
	for q.Size() > 0 {
		for i := q.Size(); i > 0; i-- {
			j, _ := q.Dequeue()
			jolts := j.(Joltage)
			if jolts == zero {
				return cnt
			}
			for _, toggle := range machine.toggles {
				next, possible := apply(jolts, toggle)
				if !possible {
					continue
				}
				if !enqd[next] {
					q.Enqueue(next)
					enqd[next] = true
				}
			}
		}
		cnt++
	}
	return uint64(0) // none found
}

func part1(machines []Machine) uint64 {
	ans := uint64(0)
	for _, machine := range machines {
		ans += bfs1(machine)
	}
	return ans
}

func part2(machines []Machine) uint64 {
	ans := uint64(0)
	for _, machine := range machines {
		ans += bfs2(machine)
	}
	return ans
}

func Run() {
	inp, err := os.Open("./day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	machines := []Machine{}
	for scanner.Scan() {
		machines = append(machines, *NewMachine(scanner.Text()))
	}

	fmt.Println("Part 1:", part1(machines))
	// fmt.Println("Part 2:", part2(machines))
	fmt.Println("Part 2:", "See `2025/10/solution.py` for an ILP-based solution.")
}
