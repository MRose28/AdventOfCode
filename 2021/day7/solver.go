package day7

import (
	"mrose.de/aoc/utility"
)

func Solve() (result int) {
	input := utility.IntArr(utility.Input2021Day7(), ",")

	fuelMap := NewFuelMap(input)

	return getSmallest(fuelMap)
}

func getSmallest(fuelMap map[int]int) (result int) {
	for _, v := range fuelMap {
		if result == 0 {
			result = v
			continue
		}
		if v<result {
			result = v
		}
	}
	return
}

func findPosition(diffMap map[int]float64) (result int, diff float64) {
	diff = 0
	for k, v := range diffMap {
		if result == 0 {
			diff = v
			result = k
			continue
		}
		if v < diff {
			diff = v
			result = k
		}
	}
	return
}

func NewDiffMap(input []int) (diffMap map[int]float64) {
	_, max := utility.MinMax(input)
	diffMap = make(map[int]float64, 0)
	for i := 0; i < max; i++ {
		diffSum := 0
		for _, v := range input {
			min, max := utility.MinMax([]int{i, v})
			diffSum += max - min
		}
		diffMap[i] = float64(diffSum) / float64(len(input))
	}
	return
}

func NewFuelMap(input []int) (diffMap map[int]int) {
	_, max := utility.MinMax(input)
	diffMap = make(map[int]int, 0)
	for i := 0; i <= max; i++ {
		fuelBurnt := 0
		for _, v := range input {
			min, max := utility.MinMax([]int{i, v})
			fuelIncrease := 0
			for j := min; j < max; j++ {
				fuelBurnt++
				fuelBurnt += fuelIncrease
				fuelIncrease++
			}
		}
		diffMap[i] = fuelBurnt
	}
	return
}
