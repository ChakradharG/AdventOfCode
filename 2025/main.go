package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChakradharG/AdventOfCode/2025/day01"
	"github.com/ChakradharG/AdventOfCode/2025/day02"
	"github.com/ChakradharG/AdventOfCode/2025/day03"
	"github.com/ChakradharG/AdventOfCode/2025/day04"
	"github.com/ChakradharG/AdventOfCode/2025/day05"
	"github.com/ChakradharG/AdventOfCode/2025/day06"
	"github.com/ChakradharG/AdventOfCode/2025/day07"
	"github.com/ChakradharG/AdventOfCode/2025/day08"
	"github.com/ChakradharG/AdventOfCode/2025/day09"
	"github.com/ChakradharG/AdventOfCode/2025/day10"
	"github.com/ChakradharG/AdventOfCode/2025/day11"
	"github.com/ChakradharG/AdventOfCode/2025/day12"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a day to run (e.g., '01', '02').")
	}
	day := os.Args[1]

	switch day {
	case "01":
		day01.Run()
	case "02":
		day02.Run()
	case "03":
		day03.Run()
	case "04":
		day04.Run()
	case "05":
		day05.Run()
	case "06":
		day06.Run()
	case "07":
		day07.Run()
	case "08":
		day08.Run()
	case "09":
		day09.Run()
	case "10":
		day10.Run()
	case "11":
		day11.Run()
	case "12":
		day12.Run()
	default:
		fmt.Printf("Day %s not found.\n", day)
	}
}
