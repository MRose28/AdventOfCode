package day2

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

const (
	ADD            = 1
	MULTIPLIKATION = 2
	HALT           = 99
)

func Solve() (p1, p2 int) {

	input := utility.InputAsStrArr(2019, 2, false)
	numbers := parseNumbers(input)
	p1 = getP1(numbers)

	return
}

func getP1(numbers []int) int {
	numbers[1] = 12
	numbers[2] = 2
	for i := 0; i < len(numbers); i += 4 {
		switch numbers[i] {
		case ADD:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
		case MULTIPLIKATION:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
		case HALT:
			return numbers[0]
		}
	}
	return -1
}

func parseNumbers(input []string) []int {
	n := make([]int, 0)
	for _, s := range strings.Split(input[0], ",") {
		number, _ := strconv.Atoi(s)
		n = append(n, number)
	}
	return n
}
