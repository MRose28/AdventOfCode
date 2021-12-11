package day8

import (
	"mrose.de/aoc/utility"
	"strings"
)

type Display struct {
	InputValues, OutputValues []string
	Value                     int
	Connections               []*WireConnection
}

func (d *Display) decryptConnections() {

}

func NewDisplay(inputValues, outputValues []string) *Display {
	display := &Display{
		InputValues:  inputValues,
		OutputValues: outputValues,
		Connections:  NewConnections(),
	}

	display.decryptConnections()
	return display
}

type Position int

const (
	TOP Position = iota
	MIDDLE
	BOTTOM
	RIGHTTOP
	RIGHTBOTTOM
	LEFTTOP
	LEFTBOTTOM
)

type WireConnection struct {
	Position   Position
	PosString  string
	active     bool
	Candidates []string
}

func NewConnections() []*WireConnection {
	return []*WireConnection{
		{Position: TOP, Candidates: make([]string, 0)},
		{Position: BOTTOM, Candidates: make([]string, 0)},
		{Position: MIDDLE, Candidates: make([]string, 0)},
		{Position: RIGHTTOP, Candidates: make([]string, 0)},
		{Position: RIGHTBOTTOM, Candidates: make([]string, 0)},
		{Position: LEFTTOP, Candidates: make([]string, 0)},
		{Position: LEFTBOTTOM, Candidates: make([]string, 0)},
	}

}

func Solve() (result int) {

	input := utility.StrArr(utility.Input2021Day8())

	displayList := parseInput(input)

	//return countSimpleNumbers(outputValues)
	return len(displayList)
}

func countSimpleNumbers(values []string) (result int) {
	// length of known numbers
	one, four, seven, eight := 2, 4, 3, 7
	for _, value := range values {
		switch len(value) {
		case one, four, seven, eight:
			result++
		}
	}
	return
}

func parseInput(input []string) (displayList []*Display) {
	displayList = make([]*Display, 0)
	for _, line := range input {
		inputValues := make([]string, 0)
		outputValues := make([]string, 0)
		for _, value := range strings.Split(strings.Split(line, "|")[0], " ") {
			inputValues = append(inputValues, value)
		}
		for _, value := range strings.Split(strings.Split(line, "|")[1], " ") {
			outputValues = append(outputValues, value)
		}
		displayList = append(displayList, NewDisplay(inputValues, outputValues))
	}
	return displayList
}
