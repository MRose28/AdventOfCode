package day1

import (
	"mrose.de/aoc/utility"
)

func Solve() (result int) {
	input := utility.StrArrCustom(utility.Input(2022, 1), "\n\n")
	return MaxTopThree(input)
}

func Max(input []string) (result int) {
	max := 0
	for _, elv := range input {
		inv := utility.IntArr(elv, "\n")
		totalCInInv := 0
		for _, calories := range inv {
			totalCInInv += calories
		}
		if totalCInInv > max {
			max = totalCInInv
		}
	}
	return max
}

// Star 2
func MaxTopThree(input []string) (result int) {
	max1 := 0
	max2 := 0
	max3 := 0

	for _, elv := range input {
		inv := utility.IntArr(elv, "\n")
		totalCInInv := 0
		for _, calories := range inv {
			totalCInInv += calories
		}
		if totalCInInv > max1 {
			max2 = max1
			max3 = max2
			max1 = totalCInInv
		} else if totalCInInv < max1 && totalCInInv > max2 {
			max3 = max2
			max2 = totalCInInv
		} else if totalCInInv < max1 && totalCInInv < max2 && totalCInInv > max3 {
			max3 = totalCInInv
		}
	}
	result = max1 + max2 + max3
	return
}
