package day12

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Direction string
type Distance int
type Degree int

type ship struct {
	northPos  int
	eastPos   int
	direction Direction
	wayPoint  wayPoint
}

type wayPoint struct {
	northTarget int
	eastTarget  int
}

func newShip() ship {
	return ship{northPos: 0, eastPos: 0, direction: "E", wayPoint: wayPoint{northTarget: 1, eastTarget: 10}}
}

type instruction struct {
	action string
	number int
}

func newInstruction(input string) instruction {
	arr := strings.SplitN(input, "", 2)
	number, _ := strconv.Atoi(arr[1])
	inst := instruction{action: arr[0], number: number}
	return inst
}

func Solve() (result int) {
	instructions := instructions()
	ship := newShip()
	runInstructions(instructions, &ship)
	normalizeShip(&ship)
	result = ship.northPos + ship.eastPos
	return
}

/*The ship only uses north and east coordinates. A south coordinate is a negative north coordinate in this model.
The task is to sum up the absolute values, so we have to make sure positive values of are used.
This is ok for the task. But usually the operation should not be done on the ship, but on the sum up operation in
the Solve method.*/
func normalizeShip(ship *ship) {
	if ship.eastPos < 0 {
		ship.eastPos = ship.eastPos * -1
	}
	if ship.northPos < 0 {
		ship.northPos = ship.northPos * -1
	}
}

//Build an array of instructions to run.
func instructions() (instructions []instruction) {
	iArr := utility.StrArr(utility.Input2020Day12())
	instructions = make([]instruction, 0)
	for _, value := range iArr {
		instructions = append(instructions, newInstruction(value))
	}
	return
}

//Run all instructions.
func runInstructions(instructions []instruction, ship *ship) {
	for _, instruction := range instructions {
		if strings.Contains("NSEW", instruction.action) {
			move(ship, instruction)
		} else if strings.Contains("RL", instruction.action) {
			turn(ship, instruction)
		} else if instruction.action == "F" {
			forward(ship, instruction)
		}
	}
}

/*------------------Star2--------------------------*/
//Just move the waypoint according to the instruction. The ship stays put.
func move(ship *ship, instruction instruction) {
	switch instruction.action {
	case "N":
		ship.wayPoint.northTarget += instruction.number
	case "E":
		ship.wayPoint.eastTarget += instruction.number
	case "S":
		ship.wayPoint.northTarget -= instruction.number
	case "W":
		ship.wayPoint.eastTarget -= instruction.number
	default:
		fmt.Print("something went wrong. couldnt read instruction for movement:\n")
	}

}

//turn the waypoint around the ship. Values between waypoint coordinates must be switched accordingly.
func turn(ship *ship, instruction instruction) {
	directions := []string{
		"N", "E", "S", "W",
	}
	north := ship.wayPoint.northTarget
	east := ship.wayPoint.eastTarget
	switch instruction.action {
	case "R":
		steps := (instruction.number / 90) % len(directions)
		switch steps {
		case 0:
			break
		case 1:
			ship.wayPoint.northTarget = -1 * east
			ship.wayPoint.eastTarget = north
		case 2:
			ship.wayPoint.eastTarget *= -1
			ship.wayPoint.northTarget *= -1
		case 3:
			ship.wayPoint.northTarget = east
			ship.wayPoint.eastTarget = -1 * north
		default:
			fmt.Printf("couldnt handle steps: %d", steps)

		}

	case "L":
		steps := (instruction.number / 90) % len(directions)
		switch steps {
		case 0:
			break
		case 1:
			ship.wayPoint.northTarget = east
			ship.wayPoint.eastTarget = -1 * north
		case 2:
			ship.wayPoint.eastTarget *= -1
			ship.wayPoint.northTarget *= -1
		case 3:
			ship.wayPoint.northTarget = -1 * east
			ship.wayPoint.eastTarget = north
		default:
			fmt.Printf("couldnt handle steps: %d", steps)

		}
	}
}

//Move the ship to the waypoint x times.
func forward(ship *ship, instruction instruction) {
	for i := 0; i < instruction.number; i++ {
		ship.northPos += ship.wayPoint.northTarget
		ship.eastPos += ship.wayPoint.eastTarget
	}
}

/*--------------------Star1---------------------------*/
//func move(ship *ship, instruction instruction) {
//	switch instruction.action {
//	case "N":
//		ship.northPos += instruction.number
//	case "E":
//		ship.eastPos += instruction.number
//	case "S":
//		ship.northPos -= instruction.number
//	case "W":
//		ship.eastPos -= instruction.number
//	default:
//		fmt.Print("something went wrong. couldnt read instruction for movement:\n")
//	}
//
//}
//
//func turn(ship *ship, instruction instruction) {
//	directions := []string{
//		"N", "E", "S", "W",
//	}
//	switch instruction.action {
//	case "R":
//		steps := (instruction.number / 90) % len(directions)
//		for i, v := range directions {
//			if ship.direction == direction(v) {
//				ship.direction = direction(directions[(i+steps)%len(directions)])
//				break
//			}
//		}
//
//	case "L":
//		steps := (instruction.number / 90) % len(directions)
//		for i, v := range directions {
//			if ship.direction == direction(v) {
//				index := i - steps
//				if steps > i {
//					index = len(directions) - (steps - i)
//				}
//				if index < 0 {
//					index = index * -1
//				}
//				ship.direction = direction(directions[index])
//				break
//			}
//		}
//	}
//}
//
//func forward(ship *ship, instruction instruction) {
//	switch ship.direction {
//	case "N":
//		ship.northPos += instruction.number
//	case "E":
//		ship.eastPos += instruction.number
//	case "S":
//		ship.northPos -= instruction.number
//	case "W":
//		ship.eastPos -= instruction.number
//	default:
//		fmt.Print("something went wrong. couldnt read instruction for movement:\n")
//	}
//}
