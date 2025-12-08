package main

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

func part1(line string) uint64 {
	ans := uint64(0)
	for rng := range strings.SplitSeq(line, ",") {
		lstr, rstr, _ := strings.Cut(rng, "-")
		l, _ := strconv.ParseUint(lstr, 10, 64)
		r, _ := strconv.ParseUint(rstr, 10, 64)
		for id := l; id <= r; id++ {
			if isRepeating(strconv.FormatUint(id, 10), 2) {
				ans += id
			}
		}
	}
	return ans
}

func part2(line string) uint64 {
	ans := uint64(0)
	for rng := range strings.SplitSeq(line, ",") {
		lstr, rstr, _ := strings.Cut(rng, "-")
		l, _ := strconv.ParseUint(lstr, 10, 64)
		r, _ := strconv.ParseUint(rstr, 10, 64)
		for id := l; id <= r; id++ {
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

func main() {
	inp, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSpace(string(inp))

	// fmt.Println(part1(line))
	fmt.Println(part2(line))
}
