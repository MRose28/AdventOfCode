package day18

import (
	"log"
	"mrose.de/aoc/utility"
	"strconv"
)

type SnailFishNumber struct {
	Left, Right           *SnailFishNumber
	LeftValue, RightValue int
	Depth                 int
}

func (sfn *SnailFishNumber) NewNumber(depth int) {
	sfn.Right = &SnailFishNumber{
		Depth: depth,

	}
}

func (sfn *SnailFishNumber) Split() {

}

func (sfn *SnailFishNumber) Explode() {

}

func Solve() (result int) {
	input := utility.StrArr(utility.Input(2021, 18))
	sfArr := parseSFNumbersFromInput(input)
	return len(sfArr)
}

func parseSFNumbersFromInput(input []string) []*SnailFishNumber {
	sfNumbers := make([]*SnailFishNumber, 0)

	for _, line := range input {
		sfNumbers = append(sfNumbers, NewSnailFishNumber(line))
	}

	return sfNumbers
}

func NewSnailFishNumber(input string) *SnailFishNumber {

	var current *SnailFishNumber
	var isLeft bool
	var depth int
	for i := 0; i < len(input); i++ {
		c := input[i : i+1]
		if c == "[" {
			isLeft = true
			if current == nil {
				current = &SnailFishNumber{}
				continue
			}
			depth++
			current.NewNumber(depth)
			current = current.Right
			current.Depth = depth
			continue
		}
		if c == "]" {
			isLeft = false
			current = current.Left

			continue
		}
		if c == "," {
			isLeft = false
			continue
		}
		startIndex := i
		for {
			_, err := strconv.Atoi(input[i+1 : i+2])
			if err == nil {
				i++
				continue
			}
			break
		}
		value, err := strconv.Atoi(input[startIndex:i+1])
		if err != nil {
			log.Fatalf("could not parse number.\ninput: %v", input[startIndex:i+1])
		}

		if isLeft {
			current.LeftValue = value
			continue
		}
		current.RightValue = value

	}

	return current
}
