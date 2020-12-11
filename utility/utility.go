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

func Contains(s []int, searchTerm int) (contained bool) {
	for _, value := range s {
		if value == searchTerm {
			return true
		}
	}
	return false
}
