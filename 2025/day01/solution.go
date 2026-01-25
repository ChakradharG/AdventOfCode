package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(data []string) int {
	cur, ans := 50, 0

	for _, line := range data {
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

func part2(data []string) int {
	cur, ans := 50, 0

	for _, line := range data {
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

func Run() {
	inp, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
