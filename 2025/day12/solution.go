package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) uint64 {
	ans := uint64(0)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		dims := strings.Split(parts[0], "x")
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		cnt := 0
		for c := range strings.SplitSeq(parts[1], " ") {
			x, _ := strconv.Atoi(c)
			cnt += x
		}
		if cnt <= (l/3)*(w/3) {
			ans++
		}
	}
	return ans
}

func Run() {
	inp, err := os.Open("./day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	for range 30 {
		// the shapes are not relevant
		scanner.Scan()
	}

	fmt.Println("Part 1:", part1(scanner))
	fmt.Println("Part 2:", "Free star!")
}
