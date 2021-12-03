package utility

import (
	"strconv"
	"strings"
)

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

//slice containing all number from the input text
func IntArr(input string, sep string) []int {
	sArr := strings.Split(input, sep)
	iArr := make([]int, 0)
	for _, value := range sArr {
		i, _ := strconv.Atoi(value)
		iArr = append(iArr, i)
	}
	return iArr
}

//slice containing all number from the input text
func StrArr(input string) []string {
	return StrArrCustom(input, "\n")
}

//slice containing all number from the input text
func StrArrCustom(input string, separator string) []string {
	return strings.Split(input, separator)
}

//slices have to have even entry length.
func CounterMap(inputArr []string, rowLength int) map[int]int {
	result := make(map[int]int)

	for i := 0; i < rowLength; i++ {
		result[i] = 0
	}

	for _, v := range inputArr {
		for i := 0; i < rowLength; i++ {
			if IntArr(v, "")[i] == 1 {
				result[i] += 1
			}
		}
	}

	return result
}

func ContainsInt(s []int, searchTerm int) (contained bool) {
	for _, value := range s {
		if value == searchTerm {
			return true
		}
	}
	return false
}

func ContainsString(s []string, searchTerm string) (contained bool) {
	for _, value := range s {
		if value == searchTerm {
			return true
		}
	}
	return false
}

func RemoveIndexStr(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndexInt(s []int, index int) []int {
	tmpArr := make([]int, 0)
	tmpArr = append(s[:index], s[index+1:]...)
	return tmpArr
}
