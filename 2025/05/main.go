package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	L uint64
	R uint64
}

func part1(intervals []interval, ids []uint64) int {
	ans := 0
	i, j := 0, 0
	for (i < len(intervals)) && (j < len(ids)) {
		if ids[j] < intervals[i].L {
			j++
		} else if ids[j] <= intervals[i].R {
			ans++
			j++
		} else {
			i++
		}
	}
	return ans
}

func part2(intervals []interval) uint64 {
	r := intervals[0].R
	ans := r - intervals[0].L + 1
	for i := 1; i < len(intervals); i++ {
		R := intervals[i].R
		if r <= R {
			L := max(intervals[i].L-1, r)
			ans += (R - L)
			r = R
		}
	}
	return ans
}

func main() {
	inp, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	intervals := make([]interval, 0)
	ids := make([]uint64, 0)

	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		line := scanner.Text()
		left, right, flag := strings.Cut(line, "-")
		if !flag {
			break
		}

		L, _ := strconv.ParseUint(left, 10, 64)
		R, _ := strconv.ParseUint(right, 10, 64)
		intervals = append(intervals, interval{L, R})
	}
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.ParseUint(line, 10, 64)
		ids = append(ids, id)
	}
	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i].L == intervals[j].L {
			return intervals[i].R < intervals[j].R
		} else {
			return intervals[i].L < intervals[j].L
		}
	})
	sort.Slice(ids, func(i int, j int) bool {
		return ids[i] < ids[j]
	})

	// fmt.Println(part1(intervals, ids))
	fmt.Println(part2(intervals))
}
