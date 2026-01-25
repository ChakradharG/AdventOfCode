package day02

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isRepeating(sId string, x int) bool {
	n := len(sId)
	if n%x != 0 {
		return false
	}
	ln := n / x
	target := sId[:ln]
	for i := ln; i < n; i += ln {
		if target != sId[i:i+ln] {
			return false
		}
	}
	return true
}

func part1(rngs [][2]uint64) uint64 {
	ans := uint64(0)
	for _, rng := range rngs {
		for id := rng[0]; id <= rng[1]; id++ {
			if isRepeating(strconv.FormatUint(id, 10), 2) {
				ans += id
			}
		}
	}
	return ans
}

func part2(rngs [][2]uint64) uint64 {
	ans := uint64(0)
	for _, rng := range rngs {
		for id := rng[0]; id <= rng[1]; id++ {
			sId := strconv.FormatUint(id, 10)
			for x := 2; x <= len(sId); x++ {
				if isRepeating(sId, x) {
					ans += id
					break
				}
			}
		}
	}
	return ans
}

func Run() {
	inp, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rngs := [][2]uint64{}
	for rng := range strings.SplitSeq(strings.TrimSpace(string(inp)), ",") {
		lstr, rstr, _ := strings.Cut(rng, "-")
		l, _ := strconv.ParseUint(lstr, 10, 64)
		r, _ := strconv.ParseUint(rstr, 10, 64)
		rngs = append(rngs, [2]uint64{l, r})
	}

	fmt.Println("Part 1:", part1(rngs))
	fmt.Println("Part 2:", part2(rngs))
}
