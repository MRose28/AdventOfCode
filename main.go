package main

import (
	"fmt"
	dayToSolve "mrose.de/aoc/2022/day7"
	"time"
)

// Solve the puzzle
func main() {
	defer elapsed("calculation")()
	result := dayToSolve.Solve()
	fmt.Printf("\n\n%d\n\n", result)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
