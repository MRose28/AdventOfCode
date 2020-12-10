package day10

import (
	"mrose.de/aoc/utility"
	"sort"
)

var iArr []int
var runsFromPosition = make(map[int]int, 0)

func Solve(input string) (result int) {
	prepareArray(input)
	result = possibilities(0)
	return
}

//Get input (+ 0 value and max value + 3) as sorted int array
func prepareArray(input string) {
	iArr = append(utility.FullArr(input), 0)
	_, max := utility.MinMax(iArr)
	iArr = append(iArr, max+3)
	sort.Ints(iArr)
}

//Find all possibilities of using the adapters. If an adapter can be changed all changes of later elements in the
//array have to be considered, so this function has to be used recursively for each element that is positioned later
//than your current position.
func possibilities(pos int) (result int) {
	//reached the end. nothing to do anymore
	if pos == len(iArr)-1 {
		return 1
	}

	//get value and a presence check for the position you are currently at from the iterMap.
	value, ok := runsFromPosition[pos]
	//if ok you are already did the work and just return the result
	if ok {
		return value
	}
	//check all elements that come after your current position
	for x := pos + 1; x < len(iArr); x++ {
		//if the difference between position x's value and your current position's value is <= 3,
		//you could change adapters
		if iArr[x]-iArr[pos] <= 3 {
			//if you change adapters at your current position, all steps from that point onward can be adapted.
			result += possibilities(x)
		} else {
			continue
		}
	}
	runsFromPosition[pos] = result
	return
}

//Solution Star 1
func Result(arr []int) (result int) {
	prev := 0
	jolt1 := 0
	jolt3 := 0
	for _, value := range arr {
		//find 1 step diffs
		if value-prev == 1 {
			jolt1++
			//find 2 step diffs
		} else if value-prev == 3 {
			jolt3++
		}
		prev = value
	}
	result = jolt1 * (jolt3)
	return
}
