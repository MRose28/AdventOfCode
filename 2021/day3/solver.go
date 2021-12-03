package day3

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
)

type Submarine struct {
	Gamma     int64
	Epsilon   int64
	co2Rating int64
	h2oRating int64
}

func Solve() (result int) {
	input := utility.Input2021Day3()
	arr := utility.StrArr(input)
	s := Submarine{}

	s.setup(arr)
	// Part One Solution
	// return int(s.Gamma * s.Epsilon)
	return int(s.h2oRating * s.co2Rating)
}

func calcPart1(input []string, rowLength int) (gammaString, epsilonString string) {
	counterMap := utility.CounterMap(input, rowLength)
	for i := 0; i < rowLength; i++ {
		if counterMap[i] > (len(input) / 2) {
			gammaString += "1"
			epsilonString += "0"
			continue
		}
		gammaString += "0"
		epsilonString += "1"
	}
	return
}

func calcPart2(input []string, rowLength int) (co2String, h2oString string) {
	h2oString = calcH2o(input, rowLength)
	co2String = calcCo2(input, rowLength)
	return
}

func calcH2o(input []string, rowLength int) string {
	selectionValue := 0
	counterMap := make(map[int]int)
	for i := 0; i < rowLength; i++ {
		nextList := make([]string, 0)
		counterMap = utility.CounterMap(input, rowLength)

		if counterMap[i]*10 >= len(input)*10/2 {
			selectionValue = 1
		} else {
			selectionValue = 0
		}

		for _, v := range input {
			if utility.IntArr(v, "")[i] == selectionValue {
				nextList = append(nextList, v)
			}
		}
		input = nextList
		if len(nextList) == 1 {
			break
		}
	}
	return input[0]
}

func calcCo2(input []string, rowLength int) string {
	selectionValue := 0
	counterMap := make(map[int]int)
	for i := 0; i < rowLength; i++ {
		nextList := make([]string, 0)
		counterMap = utility.CounterMap(input, rowLength)

		if counterMap[i]*10 < len(input)*10/2 {
			selectionValue = 1
		} else {
			selectionValue = 0
		}

		for _, v := range input {
			if utility.IntArr(v, "")[i] == selectionValue {
				nextList = append(nextList, v)
			}
		}
		input = nextList
		if len(nextList) == 1 {
			break
		}
	}
	return input[0]
}

func (s *Submarine) setup(arr []string) {
	gammaString, epsilonString := calcPart1(arr, len(utility.StrArrCustom(arr[0], "")))

	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		fmt.Printf("Could not parse gamma value. Error:\n%e", err)
	}
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		fmt.Printf("Could not parse epsilon value. Error:\n%e", err)
	}
	co2String, h2oString := calcPart2(arr, len(utility.StrArrCustom(arr[0], "")))

	co2, err := strconv.ParseInt(co2String, 2, 64)
	if err != nil {
		fmt.Printf("Could not parse co2 value. Error:\n%e", err)
	}
	h2o, err := strconv.ParseInt(h2oString, 2, 64)
	if err != nil {
		fmt.Printf("Could not parse h2o value. Error:\n%e", err)
	}
	s.Gamma = gamma
	s.Epsilon = epsilon
	s.co2Rating = co2
	s.h2oRating = h2o
}
