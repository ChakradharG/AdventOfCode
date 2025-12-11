package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var M, N uint64

type Point struct{ X, Y uint64 }

type Rect struct{ leftBorder, rightBorder, topBorder, bottomBorder, area uint64 }

type Edge struct {
	isHorizontal bool
	c0, c1, c2   uint64
}

func createRect(p0 Point, p1 Point) Rect {
	return Rect{
		min(p0.X, p1.X), max(p0.X, p1.X),
		min(p0.Y, p1.Y), max(p0.Y, p1.Y),
		area(p0, p1),
	}
}

func createEdge(p0 Point, p1 Point) Edge {
	if p0.X == p1.X {
		return Edge{false, p0.X, min(p0.Y, p1.Y), max(p0.Y, p1.Y)}
	} else {
		return Edge{true, p0.Y, min(p0.X, p1.X), max(p0.X, p1.X)}
	}
}

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

func intersects(e Edge, rect Rect) bool {
	if e.isHorizontal {
		yCoordOk := (rect.topBorder < e.c0) && (e.c0 < rect.bottomBorder)
		cutsLeft := (e.c1 <= rect.leftBorder) && (rect.leftBorder <= e.c2)
		cutsRight := (e.c1 <= rect.rightBorder) && (rect.rightBorder <= e.c2)
		return yCoordOk && (cutsLeft || cutsRight)
	} else {
		xCoordOk := (rect.leftBorder < e.c0) && (e.c0 < rect.rightBorder)
		cutsTop := (e.c1 <= rect.topBorder) && (rect.topBorder <= e.c2)
		cutsBottom := (e.c1 <= rect.bottomBorder) && (rect.bottomBorder <= e.c2)
		return xCoordOk && (cutsTop || cutsBottom)
	}
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

func part2(points []Point) uint64 {
	/*
		Finds the largest rectangle such that no edge intersects it
		and does not check if it lies inside the polygon formed by the edges.
		This works for the AoC problem because of the shape of the polygon, but might
		fail on arbitrary inputs.
		For example, given the following input:
			1,1
			10,1
			10,9
			9,9
			9,2
			1,2
		this algorithm will ouput the rectangle between points (1,2) and (9, 9)
		which lies outside the polygon.
	*/
	count := len(points)
	rects := []Rect{}
	edges := []Edge{}
	for i := range points {
		for j := i + 1; j < count; j++ {
			rects = append(rects, createRect(points[i], points[j]))
		}
		edges = append(edges, createEdge(points[i], points[(i-1+count)%count])) // last point is connect to the first point
	}
	sort.Slice(rects, func(i, j int) bool {
		return rects[j].area < rects[i].area
	})

	for _, rect := range rects {
		empty := true
		for _, e := range edges {
			if intersects(e, rect) {
				empty = false
				break
			}
		}
		if empty {
			return rect.area
		}
	}
	return uint64(0) // none found
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

	// fmt.Println(part1(points))
	fmt.Println(part2(points))
}
