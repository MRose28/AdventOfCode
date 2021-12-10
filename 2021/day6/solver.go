package day6

import (
	"mrose.de/aoc/utility"
)

// Each lantern fish creates a new fish every 7 days
// This is not synchronized
// First generation starts its timer to 8. The fish who created the new one has its timer reset to 6.

var resetValue = 6

type LanternFish struct {
	Timer int
}

func NewLanternFish() *LanternFish {
	return &LanternFish{
		Timer: 8,
	}
}

func StartingLanternFish(timer int) *LanternFish {
	return &LanternFish{Timer: timer}
}

func (l *LanternFish) reset() {
	l.Timer = resetValue
}

func Solve() (result int) {
	input := utility.IntArr(utility.Input2021Day6(), ",")
	fishArr := getOriginalFish(input)
	fishArrAfterGrowth := simulateGrowth(fishArr, 80)
	return len(fishArrAfterGrowth)
}

func simulateGrowth(arr []*LanternFish, days int) []*LanternFish {
	var fish *LanternFish
	for i := 0; i < days; i++ {
		for i := len(arr)-1; i >= 0 ; i-- {
			fish = arr[i]
			if fish.Timer==0 {
				arr = append(arr, NewLanternFish())
				fish.reset()
				continue
			}
			fish.Timer--
		}
	}
	return arr
}

func getOriginalFish(input []int) []*LanternFish {
	resultArr := make([]*LanternFish, 0)

	for _, v := range input {
		resultArr = append(resultArr, StartingLanternFish(v))
	}
	return resultArr
}
