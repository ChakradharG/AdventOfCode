package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) int {
	cur, ans := 50, 0

	for scanner.Scan() {
		line := scanner.Text()

		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		steps %= 100

		if line[0] == 'L' {
			cur = (cur - steps + 100) % 100
		} else {
			cur = (cur + steps) % 100
		}

		if cur == 0 {
			ans++
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

	scanner := bufio.NewScanner(inp)

	fmt.Println(part1(scanner))
}
