package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	bh "github.com/emirpasic/gods/trees/binaryheap"
)

type Box struct{ X, Y, Z float64 }

func distance(b1, b2 Box) float64 {
	return math.Sqrt(math.Pow(b1.X-b2.X, 2) + math.Pow(b1.Y-b2.Y, 2) + math.Pow(b1.Z-b2.Z, 2))
}

type DSU struct {
	parent []int
	size   []int
}

func (d *DSU) Union(a int, b int) {
	aPar := d.Find(a)
	bPar := d.Find(b)
	if aPar != bPar {
		d.parent[bPar] = aPar
		d.size[aPar] += d.size[bPar]
	}
}

func (d *DSU) Find(a int) int {
	if d.parent[a] != a {
		d.parent[a] = d.Find(d.parent[a])
	}
	return d.parent[a]
}

type HeapElem struct {
	dist float64
	a    int
	b    int
}

func customComparator(a, b interface{}) int {
	d1 := a.(HeapElem).dist
	d2 := b.(HeapElem).dist

	switch {
	case d1 > d2:
		return -1
	case d1 < d2:
		return 1
	default:
		return 0
	}
}

func part1(boxes []Box, dsu *DSU) uint64 {
	n := len(boxes)
	heap := bh.NewWith(customComparator)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			h := HeapElem{distance(boxes[i], boxes[j]), i, j}
			heap.Push(h)
			if heap.Size() > 1000 {
				heap.Pop()
			}
		}
	}

	for _, v := range heap.Values() {
		a := v.(HeapElem).a
		b := v.(HeapElem).b
		dsu.Union(a, b)
	}

	m0, m1, m2 := 1, 1, 1
	for i := range n {
		if dsu.Find(i) != i {
			continue
		}
		sz := dsu.size[i]
		if sz > m0 {
			m0, m1, m2 = sz, m0, m1
		} else if sz > m1 {
			m1, m2 = sz, m1
		} else if sz > m2 {
			m2 = sz
		}

	}

	return uint64(m0 * m1 * m2)
}

func main() {
	inp, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	boxes := make([]Box, 0)
	dsu := DSU{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		X, _ := strconv.ParseFloat(parts[0], 64)
		Y, _ := strconv.ParseFloat(parts[1], 64)
		Z, _ := strconv.ParseFloat(parts[2], 64)
		box := Box{X, Y, Z}
		boxes = append(boxes, box)
		dsu.parent = append(dsu.parent, len(dsu.parent))
		dsu.size = append(dsu.size, 1)
	}

	fmt.Println(part1(boxes, &dsu))
}
