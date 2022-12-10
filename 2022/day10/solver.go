package day10

import (
	"fmt"
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type instruction struct {
	add   bool
	value int
}

type clockCircuit struct {
	cycle int
	x     int
}

type crt struct {
	cPos, rPos int
}

var CRT *crt

var signalStrength int

// Part 2 is printed in the console this time
func Solve() (part1, part2 int) {
	input := utility.Input(2022, 10)
	instructions := createInstructions(input)
	runCycles(instructions)
	part1 = signalStrength

	return
}

func runCycles(instructions []instruction) (result int) {
	circuit := &clockCircuit{x: 1}
	CRT = &crt{}
	for _, instruction := range instructions {
		addCycles(circuit, instruction)
	}

	return
}

func addCycles(circuit *clockCircuit, i instruction) {
	if i.add {
		for i := 0; i < 2; i++ {
			circuit.cycle++
			checkSignalStrength(circuit)
			printCRT(circuit)
		}
		addX(circuit, i)
		return
	}
	circuit.cycle++
	checkSignalStrength(circuit)
	printCRT(circuit)
}

func addX(circuit *clockCircuit, i instruction) {
	circuit.x += i.value
}

func checkSignalStrength(circuit *clockCircuit) {
	if circuit.cycle == 20 || (circuit.cycle-20)%40 == 0 {
		signalStrength += circuit.cycle * circuit.x
	}
}

func printCRT(circuit *clockCircuit) {
	if circuit.x == CRT.cPos || circuit.x+1 == CRT.cPos || circuit.x-1 == CRT.cPos {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	CRT.cPos++
	if int(circuit.cycle/40) > CRT.rPos {
		CRT.rPos++
		CRT.cPos = 0
		fmt.Print("\n")
	}
}

func createInstructions(input string) (instructions []instruction) {
	arr := utility.StrArr(input)
	instructions = make([]instruction, 0)
	for _, s := range arr {
		if strings.Split(s, " ")[0] == "noop" {
			instructions = append(instructions, instruction{})
		} else {
			value, err := strconv.Atoi(strings.Split(s, " ")[1])
			if err != nil {
				log.Fatalf("error parsing input data. on instruction 'addx' a number is expected"+
					" in the second part. Got %v", strings.Split(s, " ")[1])
			}
			instructions = append(instructions, instruction{
				add:   true,
				value: value,
			})
		}
	}
	return
}
