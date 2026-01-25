package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	adj  map[string][]string
	memo map[string]uint64
)

func dfs(src string, dst string) uint64 {
	if src == dst {
		return uint64(1)
	}
	key := src + dst
	_, ok := memo[key]
	if !ok {
		for _, v := range adj[src] {
			memo[key] += dfs(v, dst)
		}
	}
	return memo[key]
}

func part1() uint64 {
	memo = make(map[string]uint64)
	return dfs("you", "out")
}

func part2() uint64 {
	return (dfs("svr", "dac")*dfs("dac", "fft")*dfs("fft", "out") +
		dfs("svr", "fft")*dfs("fft", "dac")*dfs("dac", "out"))
}

func Run() {
	inp, err := os.Open("./day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	adj = make(map[string][]string)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		u := parts[0]
		for v := range strings.SplitSeq(parts[1], " ") {
			adj[u] = append(adj[u], v)
		}
	}

	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
