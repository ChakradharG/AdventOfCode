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
		line := string(scanner.Text())
		n := len(line)
		row0 := [12]string{}
		row1 := [12]string{}
		row1[0] = string(line[n-1])
		k := 2
		for i := n - 2; i >= 0; i-- {
			row0[0] = max(row1[0], string(line[i]))
			for j := 1; j < k; j++ {
				row0[j] = max(
					string(line[i])+row1[j-1],
					row1[j],
				)
			}
			k = min(12, k+1)
			row0, row1 = row1, row0
		}
		cur, _ := strconv.Atoi(row1[11])
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
