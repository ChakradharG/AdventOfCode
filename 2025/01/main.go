package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) int {
	cur, ans := 50, 0

	for scanner.Scan() {
		line := scanner.Text()

		steps, _ := strconv.Atoi(line[1:])
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

func part2(scanner *bufio.Scanner) int {
	cur, ans := 50, 0

	for scanner.Scan() {
		line := scanner.Text()

		steps, _ := strconv.Atoi(line[1:])
		ans += steps / 100 // crosses zero
		steps %= 100

		if line[0] == 'L' {
			if cur != 0 && steps >= cur {
				ans++
			}
			cur = (cur - steps + 100) % 100
		} else {
			if cur != 0 && steps >= (100-cur) {
				ans++
			}
			cur = (cur + steps) % 100
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

	scanner := bufio.NewScanner(inp)

	// fmt.Println(part1(scanner))
	fmt.Println(part2(scanner))
}
