package main

import (
	"fmt"
	"time"

	dayToSolve "mrose.de/aoc/2024/day3"
)

// Solve the puzzle
func main() {
	defer elapsed("calculation")()
	part1, part2 := dayToSolve.Solve()
	fmt.Printf("\npart1: %d\npart2: %d\n", part1, part2)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
