package day23

import "mrose.de/aoc/utility"

type Hike string

const (
	Path   Hike = "."
	Forest Hike = "#"
	Right  Hike = ">"
	Left   Hike = "<"
	Up     Hike = "^"
	Down   Hike = "v"
)

func Solve() (p1, p2 int) {
	input := utility.InputAsStrArr(2023, 23, true)

	p1 = len(input)
	return
}
