package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
)

type Coord struct {
	R int
	C int
}

var dirs = []Coord{
	{-1, -1}, {-1, 0}, {-1, +1},
	{0, -1}, {0, +1},
	{+1, -1}, {+1, 0}, {+1, +1},
}

func inBounds(c Coord, m int, n int) bool {
	return (0 <= c.R && c.R < m) && (0 <= c.C && c.C < n)
}

func part1(graph [][]rune) int {
	m, n := len(graph), len(graph[0])
	ans := 0
	for i := range m {
		for j := range n {
			if graph[i][j] != '@' {
				continue
			}
			cur := 0
			for _, d := range dirs {
				c := Coord{i + d.R, j + d.C}
				if inBounds(c, m, n) && (graph[c.R][c.C] == '@') {
					cur += 1
				}
			}
			if cur < 4 {
				ans++
			}
		}
	}
	return ans
}

func part2(graph [][]rune) int {
	m, n := len(graph), len(graph[0])
	neiCount := make([][]int, m)
	for i := range m {
		neiCount[i] = make([]int, n)
	}
	for i := range m {
		for j := range n {
			if graph[i][j] != '@' {
				continue
			}
			for _, d := range dirs {
				c := Coord{i + d.R, j + d.C}
				if inBounds(c, m, n) && (graph[c.R][c.C] == '@') {
					neiCount[c.R][c.C]++
				}
			}
		}
	}

	q := dll.New()
	for i := range m {
		for j := range n {
			if graph[i][j] == '@' && neiCount[i][j] < 4 {
				q.Add(Coord{i, j})
				graph[i][j] = '.'
			}
		}
	}

	ans := 0
	for q.Size() > 0 {
		ci, _ := q.Get(0)
		c := ci.(Coord)
		q.Remove(0)
		ans++
		for _, d := range dirs {
			nc := Coord{c.R + d.R, c.C + d.C}
			if inBounds(nc, m, n) && (graph[nc.R][nc.C] == '@') {
				neiCount[nc.R][nc.C]--
				if neiCount[nc.R][nc.C] < 4 {
					q.Add(nc)
					graph[nc.R][nc.C] = '.'
				}
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
	graph := make([][]rune, 0)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		graph = append(graph, line)
	}

	// fmt.Println(part1(graph))
	fmt.Println(part2(graph))
}
