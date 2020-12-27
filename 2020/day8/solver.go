package day8

import (
	"fmt"
	"strconv"
	"strings"
)

var accumulator = 0
var input string
var pos = 0
var runs = 0
var parsedLines = 0

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
	fmt.Printf("Runs: %d\nParsedLine: %d\n", runs, parsedLines)
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
	runs++
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
	parsedLines++
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

/*
--- Day 8: Handheld Halting ---
Your flight to the major airline hub reaches cruising altitude without incident. While you consider checking the in-flight menu for one of those drinks that come with a little umbrella, you are interrupted by the kid sitting next to you.

Their handheld game console won't turn on! They ask if you can take a look.

You narrow the problem down to a strange infinite loop in the boot code (your puzzle input) of the device. You should be able to fix it, but first you need to be able to run the code in isolation.

The boot code is represented as a text file with one instruction per line of text. Each instruction consists of an operation (acc, jmp, or nop) and an argument (a signed number like +4 or -20).

acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7 would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it is executed next.
jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below it, and jmp -20 would cause the instruction 20 lines above to be executed next.
nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
For example, consider the following program:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
These instructions are visited in this order:

nop +0  | 1
acc +1  | 2, 8(!)
jmp +4  | 3
acc +3  | 6
jmp -3  | 7
acc -99 |
acc +1  | 4
jmp -4  | 5
acc +6  |
First, the nop +0 does nothing. Then, the accumulator is increased from 0 to 1 (acc +1) and jmp +4 sets the next instruction to the other acc +1 near the bottom. After it increases the accumulator from 1 to 2, jmp -4 executes, setting the next instruction to the only acc +3. It sets the accumulator to 5, and jmp -3 causes the program to continue back at the first acc +1.

This is an infinite loop: with this sequence of jumps, the program will run forever. The moment the program tries to run any instruction a second time, you know it will never terminate.

Immediately before the program would run an instruction a second time, the value in the accumulator is 5.

Run your copy of the boot code. Immediately before any instruction is executed a second time, what value is in the accumulator?

Your puzzle answer was 1137.

--- Part Two ---
After some careful analysis, you believe that exactly one instruction is corrupted.

Somewhere in the program, either a jmp is supposed to be a nop, or a nop is supposed to be a jmp. (No acc instructions were harmed in the corruption of this boot code.)

The program is supposed to terminate by attempting to execute an instruction immediately after the last instruction in the file. By changing exactly one jmp or nop, you can repair the boot code and make it terminate correctly.

For example, consider the same program from above:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
If you change the first instruction from nop +0 to jmp +0, it would create a single-instruction infinite loop, never leaving that instruction. If you change almost any of the jmp instructions, the program will still eventually find another jmp instruction and loop forever.

However, if you change the second-to-last instruction (from jmp -4 to nop -4), the program terminates! The instructions are visited in this order:

nop +0  | 1
acc +1  | 2
jmp +4  | 3
acc +3  |
jmp -3  |
acc -99 |
acc +1  | 4
nop -4  | 5
acc +6  | 6
After the last instruction (acc +6), the program terminates by attempting to run the instruction below the last instruction in the file. With this change, after the program terminates, the accumulator contains the value 8 (acc +1, acc +1, acc +6).

Fix the program so that it terminates normally by changing exactly one jmp (to nop) or nop (to jmp). What is the value of the accumulator after the program terminates?
*/
