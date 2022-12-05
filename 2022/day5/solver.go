package day5

import (
	"fmt"
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type stack []string
type stacks []stack
type instruction struct {
	amount, from, to int
}
type instructions []instruction

func Solve() (result int) {

	input := utility.Input(2022, 5)
	inputStacks := createStacks(strings.Split(input, "\n\n")[0])
	manual := createInstructions(strings.Split(input, "\n\n")[1])
	strResult := moveElements2(inputStacks, manual)
	print("\n\n\n" + strResult)

	return len(inputStacks) + len(manual)
}

// star 1
func moveElements(inputStacks stacks, manual instructions) string {
	print("\n\n")
	for _, inst := range manual {
		if inst.from == 2 || inst.to == 2 {
			fmt.Printf("%+v\n", inst)
		}
	}
	for j, inputStack := range inputStacks {
		fmt.Printf("%v: ", j+1)
		for _, s := range inputStack {
			fmt.Printf(" %v", s)
		}
		fmt.Print("\n")
	}
	for i2, instruction := range manual {
		if i2 == 55 {
			fmt.Printf("")
		}
		fmt.Printf("\nINSTRUCTION NR %v:\nmove %v from %v to %v",
			i2+1, instruction.amount, instruction.from, instruction.to)
		for i := 0; i < instruction.amount; i++ {
			if len(inputStacks[instruction.from-1]) > 0 {
				e := inputStacks[instruction.from-1][0]
				inputStacks[instruction.to-1] = append(stack{e}, inputStacks[instruction.to-1]...)
				inputStacks[instruction.from-1] = inputStacks[instruction.from-1][1:]
			} else {
				log.Fatalln("stack empty")
			}
		}
		fmt.Printf("\n\nSTACKS:\n")
		for i, inputStack := range inputStacks {
			fmt.Printf("%v: %v\n", i+1, inputStack)
		}
	}

	resultString := ""
	for _, inputStack := range inputStacks {
		resultString += inputStack[0]
	}

	return resultString
}

func moveElements2(inputStacks stacks, manual instructions) string {
	print("\n\n")
	for _, inst := range manual {
		if inst.from == 2 || inst.to == 2 {
			fmt.Printf("%+v\n", inst)
		}
	}
	for j, inputStack := range inputStacks {
		fmt.Printf("%v: ", j+1)
		for _, s := range inputStack {
			fmt.Printf(" %v", s)
		}
		fmt.Print("\n")
	}
	for _, instruction := range manual {
		newItems := make([]string, 0)
		oldItems := make([]string, 0)
		for _, t := range inputStacks[instruction.from-1][0:instruction.amount] {
			newItems = append(newItems, t)
		}
		for _, t := range inputStacks[instruction.to-1] {
			oldItems = append(oldItems, t)
		}
		inputStacks[instruction.to-1] = append(newItems, oldItems...)
		inputStacks[instruction.from-1] = inputStacks[instruction.from-1][instruction.amount:]

	}

	resultString := ""
	for _, inputStack := range inputStacks {
		resultString += inputStack[0]
	}

	return resultString
}

func createInstructions(s string) (instructionList instructions) {
	instructionList = make(instructions, 0)
	for _, line := range utility.StrArr(s) {
		amount, _ := strconv.Atoi(strings.Split(strings.Split(line, "move ")[1], " ")[0])
		from, _ := strconv.Atoi(strings.Split(strings.Split(line, "from ")[1], " ")[0])
		to, _ := strconv.Atoi(strings.Split(strings.Split(line, "to ")[1], " ")[0])
		instruction := instruction{
			amount: amount,
			from:   from,
			to:     to,
		}
		instructionList = append(instructionList, instruction)
	}
	return
}

func createStacks(s string) stacks {
	strArr := utility.StrArr(s)
	idxLine := strArr[len(strArr)-1]
	strArr = strArr[0 : len(strArr)-1]
	stacksList := make(stacks, 0)
	idxs := make([]int, 0)

	counter := 1
	for true {
		idx := strings.Index(idxLine, strconv.Itoa(counter))

		if idx == -1 {
			break
		}

		idxs = append(idxs, idx)
		counter++
	}

	for i := 0; i < len(idxs); i++ {
		stacksList = append(stacksList, make(stack, 0))
	}

	for _, e := range strArr {
		for i, idx := range idxs {
			if len(e) > idx {
				element := string(e[idx])
				if element != " " && element != "[" && element != "]" {
					stacksList[i] = append(stacksList[i], element)
				}
			}
		}
	}
	print(strArr)
	return stacksList
}
