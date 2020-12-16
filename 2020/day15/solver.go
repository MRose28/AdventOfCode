package day15

import (
	"fmt"
)

type priorNumber struct {
	value         int
	lastPositions []int
}

func newSpokenNumber(value int, position int) *priorNumber {
	return &priorNumber{value: value, lastPositions: []int{position}}
}

func Solve() (result int) {
	input := []int{19,0,5,1,10,13}
	target := 300000000
	result = countToTarget(input, target)

	return
}

func countToTarget(numbers []int, targetNumber int) (result int) {
	priorNumbers := make(map[int]priorNumber, 0)
	position := 1
	var priorNumber priorNumber
	for _, v := range numbers {
		priorNumber = *newSpokenNumber(v, position)
		priorNumbers[priorNumber.value] = priorNumber
		position++
	}

	for {
		var currentNumber int
		if len(priorNumber.lastPositions) == 1 {
			currentNumber = 0
		} else {
			currentNumber = priorNumber.lastPositions[len(priorNumber.lastPositions)-1] -
				priorNumber.lastPositions[len(priorNumber.lastPositions)-2]
		}
		v, ok := priorNumbers[currentNumber]
		if ok {
			v.lastPositions = append(priorNumbers[currentNumber].lastPositions, position)
			priorNumber = v
			priorNumbers[currentNumber] = priorNumber
		} else {
			priorNumber = *newSpokenNumber(currentNumber, position)
			priorNumbers[currentNumber] = priorNumber
		}

		if position == 2020 {
			fmt.Printf("Solution Star1: %d\n\nSolution Star2: ", priorNumber.value)
		}

		if position == targetNumber {
			return priorNumber.value
		}
		position++
	}
}

