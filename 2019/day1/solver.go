package day1

import (
	"mrose.de/aoc/utility"
)

var totalFuel = 0

func Solve() (result int) {

	iArr := utility.IntArr(utility.Input2019Day1(), "\n")
	updateTotalFuel(iArr)
	result = totalFuel
	return
}

func updateTotalFuel(arr []int) {
	for _, mass := range arr {
		temp := 0
		temp += mass/3 - 2
		totalFuel += temp
		totalFuel += moduleFuel(temp)
	}
}

func moduleFuel(mass int) (result int) {
	result = mass/3 - 2
	if result > 0 {
		result += moduleFuel(result)
		return
	}
	return 0
}
