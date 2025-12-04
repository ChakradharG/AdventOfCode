package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) int {
	ans := 0
	for scanner.Scan() {
		line := string(scanner.Text())
		mx1, mx2 := 0, 0
		for i := range len(line) - 1 {
			d, _ := strconv.Atoi(string(line[i]))
			if d > mx1 {
				mx2 = 0
				mx1 = d
			} else if d > mx2 {
				mx2 = d
			}
		}
		temp, _ := strconv.Atoi(string(line[len(line)-1]))
		mx2 = max(mx2, temp)
		ans += 10*mx1 + mx2
	}
	return ans
}

func part2(scanner *bufio.Scanner) int {
	ans := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		n := len(line)
		stack := [12]rune{}
		for i := range n {
			ch := line[i]
			for j := max(0, i+12-n); j < 12; j++ {
				if stack[j] < ch {
					stack[j] = ch
					for k := j + 1; k < 12; k++ {
						stack[k] = '0'
					}
					break
				}
			}
		}
		cur := 0
		for j := range 12 {
			cur = 10*cur + int(stack[j]-'0')
		}
		ans += cur
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

	// fmt.Println(part1(scanner))
	fmt.Println(part2(scanner))
}
