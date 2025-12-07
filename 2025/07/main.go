package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) int {
	beam := map[int]bool{}
	ans := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
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

func main() {
	inp, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)

	fmt.Println(part1(scanner))
}
