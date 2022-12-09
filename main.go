package main

import (
	"fmt"
	dayToSolve "mrose.de/aoc/2022/day9"
	"time"
)

// Solve the puzzle
func main() {
	defer elapsed("calculation")()
	part1, part2 := dayToSolve.Solve()
	fmt.Printf("\n\n%d\n\n%d\n\n", part1, part2)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
