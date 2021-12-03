package day2

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type direction int

const (
	FORWARD direction = iota
	DOWN
	UP
	UNKNOWN
)

type Command struct {
	Direction direction
	Value     int
}

type Submarine struct {
	PosHor      int
	Depth       int
	Aim         int
	CommandList []Command
}

func Solve() (result int) {
	sub := Submarine{
		CommandList: loadCommands(utility.Input2021Day2()),
	}
	sub.Navigate()

	return sub.getProduct()
}

func loadCommands(input string) (commandList []Command) {
	arr := utility.StrArr(input)
	commandList = make([]Command, 0)
	var direction direction
	for _, v := range arr {
		line := strings.Split(v, " ")
		direction = getDirection(line[0])
		value, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(fmt.Errorf("%e", err))
		}
		command := Command{direction, value}
		commandList = append(commandList, command)
	}
	return
}

func getDirection(input string) direction {

	switch input {
	case "forward":
		return FORWARD
	case "up":
		return UP
	case "down":
		return DOWN
	}

	return UNKNOWN
}

func (s *Submarine) NavigateOneStar() {

	for _, v := range s.CommandList {
		switch v.Direction {
		case FORWARD:
			s.PosHor += v.Value
		case UP:
			s.Depth -= v.Value
		case DOWN:
			s.Depth += v.Value
		}
	}

}

// TwoStar
func (s *Submarine) Navigate() {

	for _, v := range s.CommandList {
		switch v.Direction {
		case FORWARD:
			s.PosHor += v.Value
			s.Depth += s.Aim * v.Value
		case UP:
			s.Aim -= v.Value
		case DOWN:
			s.Aim += v.Value
		}
	}

}

func (s *Submarine) getProduct() (result int) {
	return s.PosHor * s.Depth
}
