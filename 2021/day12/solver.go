package day12

import (
	"github.com/davecgh/go-spew/spew"
	"mrose.de/aoc/utility"
)

type Cave struct {
	Connections []*Cave
}

type SmallCave struct {
	Cave
	visited bool
}

func Solve() (result int) {
	input := utility.StrArr(utility.Input2021Day12())
	s := Cave{[]*Cave{}}

	spew.Dump(s)
	return len(input)
}
