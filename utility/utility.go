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
func IntArr(input string) []int {
	sArr := strings.Split(input, "\n")
	iArr := make([]int, 0)
	for _, value := range sArr {
		i, _ := strconv.Atoi(value)
		iArr = append(iArr, i)
	}
	return iArr
}

//slice containing all number from the input text
func StrArr(input string) []string {
	return strings.Split(input, "\n")
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

func RemoveIndex(s []string, index int) []string {
		return append(s[:index], s[index+1:]...)
}
