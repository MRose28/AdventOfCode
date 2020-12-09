package day9

import (
	"strconv"
	"strings"
)

//Get the result to check against AOC
func Solve(input string) (result int) {
	input = strings.TrimSuffix(input, "\n")
	iArr := getFullArr(input)
	p25 := getPreamble(iArr)
	small, big := getAdditionParts(getInvalidNumber(iArr, p25), iArr)
	result = big + small
	return
}

//get the 25 number preamble from the input text as slice of int
func getPreamble(arr []int) (p25 []int) {
	p25 = make([]int, 0)
	for i := 0; i < 25; i++ {
		p25 = append(p25, arr[i])
	}
	return
}

//slice containing all number from the input text
func getFullArr(input string) []int {
	sArr := strings.Split(input, "\n")
	iArr := make([]int, 0)
	for _, value := range sArr {
		i, _ := strconv.Atoi(value)
		iArr = append(iArr, i)
	}
	return iArr
}

//the only number that does not conform to AOC rules.
//returns 0 if no invalid number was found
func getInvalidNumber(arr []int, p25 []int) int {
	value := 0
	present := false
	for i := 25; i < len(arr); i++ {
		value = arr[i]
		for _, p1 := range p25 {
			for _, p2 := range p25 {
				if p1 != p2 {
					if p1+p2 == value {
						present = true
					}
				}
			}
		}
		if !present {
			return value
		}
		present = false
		p25 = p25[1:]
		p25 = append(p25, value)
	}
	return 0
}

//get min and max numbers of the slice that sums up to the invalid number
func getAdditionParts(invalidNumber int, iArr []int) (big int, small int) {
	result := 0
	addArr := make([]int, 0)
	found := false
	for !found {
		for _, value := range iArr {
			result += value
			addArr = append(addArr, value)
			if result == invalidNumber {
				found = true
				return MinMax(addArr)
			}
			if result > invalidNumber {
				iArr = iArr[1:]
				addArr = make([]int, 0)
				result = 0
				break
			}
		}
	}
	return
}

//get min and max of slice
func MinMax(array []int) (min int, max int) {
	max = array[0]
	min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return
}
