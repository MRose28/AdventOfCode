package main

import (
	"fmt"
	"mrose.de/aoc/2021/day4"
	"time"
)
//Solve the puzzle
func main() {
	defer elapsed("calculation")()
	result := day4.Solve()
	fmt.Printf("\n\n%d\n\n", result)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
