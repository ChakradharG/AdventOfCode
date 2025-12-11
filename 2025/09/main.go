package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct{ X, Y uint64 }

var M, N uint64

func manhattanDist(p Point, x uint64, y uint64) uint64 {
	x0 := min(x, p.X)
	x1 := max(x, p.X)
	y0 := min(y, p.Y)
	y1 := max(y, p.Y)

	return (x1 - x0) + (y1 - y0)
}

func area(p0 Point, p1 Point) uint64 {
	x0 := min(p0.X, p1.X)
	x1 := max(p0.X, p1.X)
	y0 := min(p0.Y, p1.Y)
	y1 := max(p0.Y, p1.Y)

	return (x1 - x0 + 1) * (y1 - y0 + 1)
}

func part1(points []Point) uint64 {
	tl, tld := points[0], manhattanDist(points[0], 0, 0)
	tr, trd := points[0], manhattanDist(points[0], 0, N)
	bl, bld := points[0], manhattanDist(points[0], M, 0)
	br, brd := points[0], manhattanDist(points[0], M, N)

	for _, point := range points {
		if manhattanDist(point, 0, 0) < tld {
			tl, tld = point, manhattanDist(point, 0, 0)
		}
		if manhattanDist(point, 0, N) < trd {
			tr, trd = point, manhattanDist(point, 0, N)
		}
		if manhattanDist(point, M, 0) < bld {
			bl, bld = point, manhattanDist(point, M, 0)
		}
		if manhattanDist(point, M, N) < brd {
			br, brd = point, manhattanDist(point, M, N)
		}
	}

	return max(
		area(tl, tr), area(tl, br), area(tl, bl),
		area(tr, br), area(tr, bl),
		area(br, bl),
	)
}

func main() {
	inp, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	points := []Point{}
	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.ParseUint(parts[0], 10, 64)
		y, _ := strconv.ParseUint(parts[1], 10, 64)
		points = append(points, Point{x, y})
		M = max(M, x+1)
		N = max(N, y+1)
	}

	fmt.Println(part1(points))
}
