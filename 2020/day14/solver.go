package day14

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type instructionSet struct {
	mask         string
	instructions []addressValuePair
}

type addressValuePair struct {
	memAddress int
	value      int
}

func Solve() (result int) {
	instructions := instructions()
	result = runInstructions(instructions)
	return
}

//Instruction Runner
func runInstructions(instructionSets []instructionSet) (result int) {
	memory := initializeMemoryMap(instructionSets)

	for _, v := range instructionSets {
		for _, instruction := range v.instructions {
			updateMemory(memory[instruction.memAddress], v.mask, instruction.value)
		}
	}
	result = calculateResult(memory)
	return
}

//Memory Initializer
func initializeMemoryMap(instructionSets []instructionSet) (memory map[int]map[int]bool) {
	memory = make(map[int]map[int]bool, 0)
	for iSets, v := range instructionSets {
		for iIns, instruction := range v.instructions {
			entry := make(map[int]bool, 0)
			for i := 0; i < 36; i++ {
				entry[i] = false
			}

			if instruction.memAddress == 0 {
				fmt.Printf("\nSet index: %d\nInstruction index: %d\n\n", iSets, iIns)
			}
			memory[instruction.memAddress] = entry
		}
	}
	return

}

//Sum of remaining values
func calculateResult(memory map[int]map[int]bool) (result int) {
	valueStr := ""
	for _, entry := range memory {
		for i := 0; i < len(entry); i++ {
			if entry[i] {
				valueStr += "1"
			} else {
				valueStr += "0"
			}
		}
		increment, err := strconv.ParseInt(valueStr, 2, 64)
		if err != nil {
			fmt.Print(err)
		}
		//if increment > 0 {
		//fmt.Printf("\nBinary(%d): %d",i , increment)
		//}
		result += int(increment)
		valueStr = ""
	}
	//fmt.Print("\n\n")

	return
}

//Update memory according to given instructions
func updateMemory(memEntry map[int]bool, mask string, value int) {
	binaryValueArray := binaryArray(value, len(strings.Split(mask, "")))
	for i := 0; i < len(binaryValueArray); i++ {
		binaryValue := binaryValueArray[len(binaryValueArray)-1-i]
		if binaryValue == "0" {
			(memEntry)[len(memEntry)-1-i] = false
		} else {
			(memEntry)[len(memEntry)-1-i] = true
		}
	}
	memEntry = applyMask(memEntry, mask)
}

func binaryArray(value int, maskLength int) (result []string) {
	arr := strings.Split(strconv.FormatInt(int64(value), 2), "")
	result = make([]string, 0)
	for i := 0; i < maskLength-len(arr); i++ {
		result = append(result, "0")
	}
	result = append(result, arr...)
	return
}

//Overwrite values according to the mask
func applyMask(entry map[int]bool, mask string) map[int]bool {
	maskContentArray := strings.Split(mask, "")
	for i, rule := range maskContentArray {
		if rule == "X" {
			continue
		}
		if rule == "0" {
			(entry)[i] = false
		}
		if rule == "1" {
			(entry)[i] = true
		}
	}
	return entry

}

//Create an array of instructions from the input string
func instructions() (arr []instructionSet) {
	arr = make([]instructionSet, 0)
	instructions := make([]addressValuePair, 0)
	stringArr := ModdedStringArray()
	var mask string
	var valuePair addressValuePair
	counter := 0
	for i, v := range stringArr {
		if strings.Split(v, "|")[0] == "mask" {
			mask = strings.Split(v, "|")[1]
			counter++
			continue
		} else {
			memAddress, _ := strconv.Atoi(strings.Split(v, "|")[0])
			value, _ := strconv.Atoi(strings.Split(v, "|")[1])
			valuePair = addressValuePair{
				memAddress: memAddress,
				value:      value,
			}
		}
		instructions = append(instructions, valuePair)
		if i == len(stringArr)-1 || strings.Split(stringArr[i+1], "|")[0] == "mask" {
			arr = append(arr, instructionSet{mask: mask, instructions: instructions})
			instructions = make([]addressValuePair, 0)
		}
	}
	return
}

//Replace some characters to make parsing easier.
func ModdedStringArray() (arr []string) {
	input := utility.Input2020Day14()
	input = strings.Replace(input, " = ", "|", -1)
	input = strings.Replace(input, "mem[", "", -1)
	input = strings.Replace(input, "]", "", -1)

	arr = utility.StrArr(input)
	return
}
