package day13

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Dot struct {
	X, Y   int
	Marked bool
}

type Direction int

const (
	UP Direction = iota
	LEFT
)

type Instruction struct {
	Direction Direction
	Value     int
}

func Solve() (result int) {
	input := utility.Input(2021, 13)
	dots := parseDots(strings.Split(input, "\n\n")[0])
	instructions := parseInstructions(strings.Split(input, "\n\n")[1])

	dots = fold(dots, instructions)
	printDots(dots)
	return len(dots)
}

func printDots(dots []*Dot) {
	maxY := 0
	maxX := 0
	for _, dot := range dots {
		if dot.Y > maxY {
			maxY = dot.Y
		}
		if dot.X > maxX {
			maxX = dot.X
		}
	}
	arr := make([]string, maxY+1)


	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			exists := false
			for _, dot := range dots {

				if dot.X == x && dot.Y == y {
					arr[y] += "X"
					exists = true
					break
				}

			}
			if !exists {
				arr[y] += "."
			}
		}
	}

	for _, s := range arr {
		fmt.Printf("%s\n", s)
	}

}

func fold(dots []*Dot, instructions []*Instruction) []*Dot {
	for _, instruction := range instructions {
		switch instruction.Direction {
		case UP:
			dots = foldUp(dots, instruction)
		case LEFT:
			dots = foldLeft(dots, instruction)
		default:
			spew.Dump(instruction)
			log.Fatalf("\nCannot read direction")
		}
	}
	return dots
}

func foldLeft(dots []*Dot, instruction *Instruction) (dotsFolded []*Dot) {
	dotsFolded = make([]*Dot, 0)
	dotsRemove := make([]*Dot, 0)

	for _, dot := range dots {
		if dot.X == instruction.Value {
			continue
		}
		if dot.X < instruction.Value {
			dotsFolded = append(dotsFolded, dot)
		} else {
			dotsRemove = append(dotsRemove, dot)
		}
	}

	newDots := make([]*Dot, 0)
	for _, dot := range dotsRemove {
		exists := false
		for _, d := range dotsFolded {
			if d.Y == dot.Y && instruction.Value-(dot.X-instruction.Value) == d.X {
				exists = true
			}
		}

		if !exists {
			newDots = append(newDots, &Dot{
				X:      instruction.Value - (dot.X - instruction.Value),
				Y:      dot.Y,
				Marked: true,
			})
		}
	}
	dotsFolded = append(dotsFolded, newDots...)
	return
}

func foldUp(dots []*Dot, instruction *Instruction) (dotsFolded []*Dot) {
	dotsFolded = make([]*Dot, 0)
	dotsRemove := make([]*Dot, 0)

	for _, dot := range dots {
		if dot.Y == instruction.Value {
			continue
		}
		if dot.Y < instruction.Value {
			dotsFolded = append(dotsFolded, dot)
		} else {
			dotsRemove = append(dotsRemove, dot)
		}
	}

	for _, dot := range dotsRemove {
		exists := false
		for _, d := range dotsFolded {
			if d.Y == instruction.Value-(dot.Y-instruction.Value) && d.X == dot.X {
				exists = true
			}
		}

		if !exists {
			dotsFolded = append(dotsFolded, &Dot{
				X:      dot.X,
				Y:      instruction.Value - (dot.Y - instruction.Value),
				Marked: true,
			})
		}
	}
	return
}

func parseInstructions(s string) (instructions []*Instruction) {
	s = strings.Replace(s, "fold along ", "", -1)
	lines := strings.Split(s, "\n")
	instructions = make([]*Instruction, len(lines))

	for i, line := range lines {
		direction := UP
		if line[0:1] == "x" {
			direction = LEFT
		}
		v, _ := strconv.Atoi(strings.Split(line, "=")[1])
		instruction := &Instruction{
			Direction: direction,
			Value:     v,
		}

		instructions[i] = instruction
	}
	return
}

func parseDots(s string) (dots []*Dot) {
	lines := strings.Split(s, "\n")
	dots = make([]*Dot, 0)

	for _, line := range lines {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		dots = append(dots,
			&Dot{
				X:      x,
				Y:      y,
				Marked: true,
			},
		)
	}
	return
}
