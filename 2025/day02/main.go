package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalid(sId string) bool {
	n := len(sId)
	if n%2 != 0 {
		return false
	}
	mid := int(n / 2)
	return sId[:mid] == sId[mid:]
}

func part1(line string) uint64 {
	ans := uint64(0)
	for _, rng := range strings.Split(line, ",") {
		lstr, rstr, _ := strings.Cut(rng, "-")
		l, _ := strconv.ParseUint(lstr, 10, 64)
		r, _ := strconv.ParseUint(rstr, 10, 64)
		for id := l; id <= r; id++ {
			if isInvalid(strconv.FormatUint(id, 10)) {
				ans += id
			}
		}
	}
	return ans
}

func main() {
	inp, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	line := strings.TrimSpace(string(inp))

	fmt.Println(part1(line))
}
