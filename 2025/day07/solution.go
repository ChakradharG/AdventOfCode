package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(data []string) int {
	beam := map[int]bool{}
	ans := 0
	for _, line := range data {
		for j, c := range line {
			if c == 'S' {
				beam[j] = true
			} else if (c == '^') && beam[j] {
				ans++
				beam[j-1] = true
				beam[j] = false
				beam[j+1] = true
			}
		}
	}
	return ans
}

func part2(data []string) int {
	beam := map[int]int{}
	ans := 0
	for _, line := range data {
		for j, c := range line {
			if c == 'S' {
				ans++
				beam[j]++
			} else if (c == '^') && (beam[j] > 0) {
				beam[j-1] += beam[j]
				beam[j+1] += beam[j]
				ans += beam[j]
				beam[j] = 0
			}
		}
	}
	return ans
}

func Run() {
	inp, err := os.Open("./day07/input.txt")
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
