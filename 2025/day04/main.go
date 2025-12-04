package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getVal(i int, j int, graph []string) int {
	if (0 <= i && i < len(graph)) && (0 <= j && j < len(graph[0])) {
		if graph[i][j] == '@' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func part1(graph []string) int {
	m, n := len(graph), len(graph[0])
	ans := 0
	for i := range m {
		for j := range n {
			if graph[i][j] == '.' {
				continue
			}
			cur := (getVal(i-1, j-1, graph) + getVal(i-1, j, graph) + getVal(i-1, j+1, graph) +
				getVal(i, j-1, graph) + getVal(i, j+1, graph) +
				getVal(i+1, j-1, graph) + getVal(i+1, j, graph) + getVal(i+1, j+1, graph))
			if cur < 4 {
				ans++
			}
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
	graph := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		graph = append(graph, line)
	}

	fmt.Println(part1(graph))
}
