package main

import (
	"bufio"
	"fmt"
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

func main() {
	inp, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
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

		L, err := strconv.ParseUint(left, 10, 64)
		if err != nil {
			panic(err)
		}
		R, err := strconv.ParseUint(right, 10, 64)
		if err != nil {
			panic(err)
		}
		intervals = append(intervals, interval{L, R})
	}
	for scanner.Scan() {
		line := scanner.Text()
		id, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			panic(err)
		}
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

	fmt.Println(part1(intervals, ids))
}
