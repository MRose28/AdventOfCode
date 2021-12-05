package day5

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Line struct {
	Start *Coordinate
	End   *Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func Solve() (result int) {
	lines, diagonals := NewLines(strings.Split(utility.Input2021Day5(), "\n"))
	// go through all coordinates
	result = findLineCoordinates(lines, diagonals)
	return
}

func findLineCoordinates(lines []*Line, diagonals []*Line) (result int) {
	counterMap := make(map[Coordinate]int, 0)
	counterMap = part1(counterMap, lines)
	counterMap = part2(counterMap, diagonals)
	for _, v := range counterMap {
		if v > 1 {
			result++
		}
	}
	return
}

func part1(counterMap map[Coordinate]int, lines []*Line) map[Coordinate]int {
	for _, line := range lines {
		if line.Start.X == line.End.X && line.Start.Y != line.End.Y {
			min, max := utility.MinMax([]int{line.Start.Y, line.End.Y})
			for i := min; i <= max; i++ {
				c := Coordinate{
					X: line.Start.X,
					Y: i,
				}
				counterMap[c]++
			}
		} else if line.Start.Y == line.End.Y && line.Start.X != line.End.X {
			min, max := utility.MinMax([]int{line.Start.X, line.End.X})
			for i := min; i <= max; i++ {
				c := Coordinate{
					X: i,
					Y: line.Start.Y,
				}
				counterMap[c]++
			}
			// part 2
		}
	}
	return counterMap
}

func part2(counterMap map[Coordinate]int, diagonals []*Line) map[Coordinate]int {
	for _, line := range diagonals {

		if line.Start.X < line.End.X && line.Start.Y < line.End.Y {
			// to top right
			for i := 0; i <= line.End.X-line.Start.X; i++ {
				c := Coordinate{
					X: line.Start.X + i,
					Y: line.Start.Y + i,
				}
				counterMap[c]++
			}
		} else if line.Start.X > line.End.X && line.Start.Y > line.End.Y {
			// to bottom left
			for i := 0; i <= line.Start.X-line.End.X; i++ {
				c := Coordinate{
					X: line.Start.X - i,
					Y: line.Start.Y - i,
				}
				counterMap[c]++
			}
		} else if line.Start.X < line.End.X && line.Start.Y > line.End.Y {
			// to bottom right
			for i := 0; i <= line.End.X-line.Start.X; i++ {
				c := Coordinate{
					X: line.Start.X + i,
					Y: line.Start.Y - i,
				}
				counterMap[c]++
			}
		} else if line.Start.X > line.End.X && line.Start.Y < line.End.Y {
			// to top left
			for i := 0; i <= line.Start.X-line.End.X; i++ {
				c := Coordinate{
					X: line.Start.X - i,
					Y: line.Start.Y + i,
				}
				counterMap[c]++
			}
		} else if line.Start.X == line.Start.Y && line.End.X == line.End.Y {
			min, max := utility.MinMax([]int{line.Start.X, line.End.X})
			for i := min; i <= max; i++ {
				c := Coordinate{
					X: i,
					Y: i,
				}
				counterMap[c]++
			}

		}

	}
	return counterMap
}

func NewLines(input []string) (lines []*Line, diagonals []*Line) {
	lines = make([]*Line, 0)
	diagonals = make([]*Line, 0)
	for _, l := range input {
		line := NewLine(l)
		if (line.Start.Y == line.End.Y && line.Start.X != line.End.X) || (line.Start.X == line.End.X && line.Start.Y != line.End.Y) {
			lines = append(lines, line)
		} else {
			diagonals = append(diagonals, line)
		}
	}
	return
}

func NewLine(l string) (line *Line) {
	start := NewCoordinate(strings.Split(l, " -> ")[0])
	end := NewCoordinate(strings.Split(l, " -> ")[1])
	return &Line{
		Start: start,
		End:   end,
	}
}

func NewCoordinate(s string) *Coordinate {
	x, err := strconv.Atoi(strings.Split(s, ",")[0])
	if err != nil {
		fmt.Errorf("couldn't parse x value. error: %e", err)
	}
	y, err := strconv.Atoi(strings.Split(s, ",")[1])
	if err != nil {
		fmt.Errorf("couldn't parse y value. error: %e", err)
	}
	return &Coordinate{
		X: x,
		Y: y,
	}
}
