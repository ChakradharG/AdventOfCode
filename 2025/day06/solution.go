package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	operands []uint64
	operator string
}

func part1(data []string) uint64 {
	problems := make([]Problem, 0)
	for _, line := range data {
		line = strings.TrimSpace(line)
		if line[0] == '+' || line[0] == '*' {
			i, j := 0, 0
			for i < len(line) {
				if line[i] != ' ' {
					problems[j].operator = string(line[i])
					j++
				}
				i++
			}
			break
		}
		for i, num := range strings.Fields(line) {
			if i == len(problems) {
				problems = append(problems, Problem{})
			}
			x, err := strconv.ParseUint(num, 10, 64)
			if err == nil {
				problems[i].operands = append(problems[i].operands, x)
			}
		}
	}

	ans := uint64(0)
	for _, problem := range problems {
		if problem.operator == "+" {
			cur := uint64(0)
			for _, operand := range problem.operands {
				cur += operand
			}
			ans += cur
		} else {
			cur := uint64(1)
			for _, operand := range problem.operands {
				cur *= operand
			}
			ans += cur
		}
	}
	return ans
}

func part2(data []string) uint64 {
	row := make([]uint64, 3722)
	ans := uint64(0)
	for _, line := range data {
		if line[0] == '+' || line[0] == '*' {
			op, cur := rune(line[0]), uint64(0)
			for j, ch := range line {
				if (j < len(line)-1) && (line[j+1] == '+' || line[j+1] == '*') {
					continue
				}
				if ch == '+' || ch == '*' {
					ans += cur
					cur = row[j]
					op = ch
				} else {
					if op == '+' {
						cur += row[j]
					} else {
						cur *= row[j]
					}
				}
			}
			ans += cur
		} else {
			for j, num := range line {
				if num == ' ' {
					continue
				}
				row[j] = 10*row[j] + uint64(num-'0')
			}
		}
	}
	return ans
}

func Run() {
	inp, err := os.Open("./day06/input.txt")
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
