package day13

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type bus struct {
	ID int
}

func Solve() (result int) {
	result = Star2()
	return
}

func Star1() (result int) {
	iArr := utility.StrArr(utility.Input2020Day13())
	time, _ := strconv.Atoi(iArr[0])
	buses := buses(iArr[1])
	bus, waitTime := findBestBus(time, buses)
	result = bus.ID * waitTime
	return
}

func Star2() (result int) {
	iArr := utility.StrArr(utility.Input2020Day13())
	result = startOfEarliestSequence(buses(iArr[1]))
	return
}

func startOfEarliestSequence(buses []bus) (timestamp int) {
	timestamp = buses[0].ID
	maxI := 0
	skip := 1
	for {
		skip = 1

		for i, bus := range buses {
			if (timestamp+i)%bus.ID != 0 {
				break
			}

			/*For partial results we can calculate when the alignment will happen again for all buses with
			index <=currentIndex, multiplying their departure cycles (in this case their IDs).
			At this point the departure cycles aligned in the correct way. Example:
			We found correctly aligning timestamps for the buses 7, 9 and the current bus 20. The next time this will
			happen is at timestamp += 7*9*20.*/
			skip *= bus.ID
			maxI = i
		}

		//Check for success
		if maxI == len(buses)-1 {
			return
		}

		//Increment timestamp with skip time, calculated from the product of the traversed buses. See comment above.
		timestamp += skip
	}
}

func busesStar1(input string) (buses []bus) {
	buses = make([]bus, 0)
	arr := strings.Split(input, ",")
	for _, id := range arr {
		if id == "x" {
			continue
		}
		intID, _ := strconv.Atoi(id)
		buses = append(buses, bus{ID: intID})
	}
	return
}

func buses(input string) (buses []bus) {
	buses = make([]bus, 0)
	arr := strings.Split(input, ",")
	intID := 1
	for _, id := range arr {
		intID = 1
		if id != "x" {
			intID, _ = strconv.Atoi(id)
		}
		buses = append(buses, bus{ID: intID})
	}
	return
}

func rest(time int, id int) (waitTime int) {
	waitTime = ((time/id + 1) * id) - time
	return
}

func findBestBus(time int, buses []bus) (bestBus bus, waitTime int) {
	for _, bus := range buses {
		if waitTime == 0 {
			waitTime = rest(time, bus.ID)
			bestBus = bus
		} else {
			currentRest := rest(time, bus.ID)
			if currentRest < waitTime {
				waitTime = currentRest
				bestBus = bus
			}
		}
	}
	return
}
