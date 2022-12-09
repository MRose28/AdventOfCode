package day9

import (
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type point struct {
	x, y           int
	previous, next *point
}

type move struct {
	direction     string
	numberOfMoves int
}

func Solve() (part1, part2 int) {

	inputArr := utility.StrArr(utility.Input(2022, 9))

	moves := createMoves(inputArr)

	//part1 = executeMoves(moves)
	part2 = executeMovesForPart2(moves)
	return
}

func createMoves(arr []string) (moves []move) {
	moves = make([]move, 0)
	for _, s := range arr {
		direction := strings.Split(s, " ")[0]
		numberOfMoves, err := strconv.Atoi(strings.Split(s, " ")[1])
		if err != nil {
			log.Fatalf("error parsing input. could not parse %s, expected format is '{string} {number}'", s)
		}
		moves = append(moves, move{
			direction:     direction,
			numberOfMoves: numberOfMoves,
		})
	}
	return
}

// Part1
func executeMoves(moves []move) (result int) {
	visitedPoints := make([]point, 0)
	head := &point{}
	tail := &point{}
	visitedPoints = append(visitedPoints, point{
		x: 0,
		y: 0,
	})
	for _, m := range moves {
		for i := 0; i < m.numberOfMoves; i++ {

			executeHeadMove(head, m)
			moveParts(head, tail)

			exists := false
			for _, p := range visitedPoints {
				if p == *tail {
					exists = true
					break
				}
			}
			if !exists {
				visitedPoints = append(visitedPoints, *tail)
			}

		}
	}
	result = len(visitedPoints)
	return
}

// Part2
func executeMovesForPart2(moves []move) (result int) {
	visitedPoints := make([]point, 0)
	head, tail := createLinkedPoints()
	visitedPoints = append(visitedPoints, point{
		x: 0,
		y: 0,
	})

	for _, m := range moves {
		for i := 0; i < m.numberOfMoves; i++ {
			executeHeadMove(head, m)
			currentNode := head
			for {
				if currentNode.next != nil {
					moved := moveParts(currentNode, currentNode.next)
					if !moved {
						break
					}
					currentNode = currentNode.next
				} else {
					exists := false
					for _, p := range visitedPoints {
						if p == *tail {
							exists = true
							break
						}
					}
					if !exists {
						visitedPoints = append(visitedPoints, *tail)
					}
					break
				}
			}
		}
	}

	return len(visitedPoints)
}

func createLinkedPoints() (head, tail *point) {
	var previous *point
	for i := 0; i < 10; i++ {
		newPoint := &point{
			x:        0,
			y:        0,
			previous: nil,
			next:     nil,
		}
		if previous != nil {
			previous.next = newPoint
			newPoint.previous = previous
		}
		if i == 0 {
			head = newPoint
		}
		if i == 9 {
			tail = newPoint
		}
		previous = newPoint
	}
	return
}

func executeHeadMove(head *point, m move) {
	switch m.direction {
	case "R":
		head.x++
	case "L":
		head.x--
	case "U":
		head.y++
	case "D":
		head.y--
	}
}

func moveParts(ahead, behind *point) (moved bool) {
	previousPoint := *behind
	if ahead.y == behind.y || ahead.x == behind.x {
		if ahead.y == behind.y {
			switch ahead.x - behind.x {
			case 2:
				behind.x++
			case -2:
				behind.x--
			}
		}
		if ahead.x == behind.x {
			switch ahead.y - behind.y {
			case 2:
				behind.y++
			case -2:
				behind.y--
			}
		}
	} else {
		if ahead.x-behind.x == 2 {
			behind.x++
			if ahead.y > behind.y {
				behind.y++
			} else {
				behind.y--
			}
		} else if ahead.x-behind.x == -2 {
			behind.x--
			if ahead.y > behind.y {
				behind.y++
			} else {
				behind.y--
			}
		}

		if ahead.y-behind.y == 2 {
			behind.y++
			if ahead.x > behind.x {
				behind.x++
			} else {
				behind.x--
			}
		} else if ahead.y-behind.y == -2 {
			behind.y--
			if ahead.x > behind.x {
				behind.x++
			} else {
				behind.x--
			}
		}
	}
	if previousPoint != *behind {
		moved = true
	}
	return
}
