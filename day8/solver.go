package day8

import (
	"strconv"
	"strings"
)

var accumulator = 0
var input string
var pos = 0

//return the result
func Solve(inputStr string) int {
	return solveStar2(inputStr)
}

//Solve star 1
func solveStar1(inputStr string) int {
	input = inputStr
	input = strings.TrimSuffix(input, "\n")
	loopInput := input
	_ = run(loopInput, false)
	return accumulator
}

//Solve star 2
func solveStar2(inputStr string) int {
	input = inputStr
	input = strings.TrimSuffix(input, "\n")
	loopInput := input
	ok := false
	for !ok {
		accumulator = 0
		ok = run(loopInput, true)
		pos++
	}
	return accumulator
}

func makeLists(loopInput string) ([]string, []int) {
	instructList := strings.Split(loopInput, "\n")
	countList := make([]int, len(instructList))

	return instructList, countList
}

//Run all instructions. If 'star2' is true 'pos' is considered and parseLine instructions are changed
//to make the program find a way to run all the way through without using any instruction twice.
func run(loopInput string, star2 bool) bool {
	iArr, cArr := makeLists(loopInput)
	ok := true
	for i := 0; i < len(iArr); i++ {
		if cArr[i] != 0 {
			ok = false
			break
		} else {
			cArr[i]++
		}
		acc := 0
		jump := 0
		if star2 {
			if i == pos {
				acc, jump = parseLine(iArr[i], true)
			} else {
				acc, jump = parseLine(iArr[i], false)
			}
		} else {
			acc, jump = parseLine(iArr[i], false)

		}

		i += jump
		accumulator += acc
	}
	return ok
}

//Parse an instruction line. If star2 is true 'jmp' and 'nop' instructions are switched
func parseLine(line string, star2 bool) (int, int) {
	acc := 0
	jmp := 0
	if strings.Contains(line, "acc") {
		acc, _ = strconv.Atoi(strings.Split(line, " ")[1])
	}
	if !star2 {
		if strings.Contains(line, "jmp") {
			jmp, _ = strconv.Atoi(strings.Split(line, " ")[1])
			jmp--
		}
	} else {
		if strings.Contains(line, "nop") {
			jmp, _ = strconv.Atoi(strings.Split(line, " ")[1])
			jmp--
		}
	}
	return acc, jmp
}
