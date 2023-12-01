package day16

import (
	"log"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

var time int = 30
var totalPressureRelease int = 0

var valves map[string]*Valve

func Solve() (part1, part2 int) {
	arr := utility.StrArr(utility.TestInput(2022, 16))
	createValves(arr)
	createConnections(arr)
	addDistances()
	calcMaxRelease()
	part1 = totalPressureRelease

	return
}

func calcMaxRelease() (result int) {
	current := getStartNode()
	for time > 0 {
		current = getNextValve(current)
	}
	return
}

func getNextValve(current *Valve) (next *Valve) {
	maxRelease := 0
	distanceToNext := 0
	for _, target := range valves {
		if target.name != current.name && !target.open && target.flowrate > 0 {
			release := (time - 1 - current.distances[target.name]) * target.flowrate

			if next == nil {
				maxRelease = release
				next = target
				distanceToNext = current.distances[target.name]
			} else {
				if release > maxRelease {
					if maxRelease-release < (distanceToNext-current.distances[target.name])*target.flowrate {
						maxRelease = release
						next = target
						distanceToNext = current.distances[target.name]
					}
				} else if current.distances[target.name] < distanceToNext {
					if maxRelease-release < (distanceToNext-current.distances[target.name])*target.flowrate {
						maxRelease = release
						next = target
						distanceToNext = current.distances[target.name]
					}
				}
			}
		}
	}
	if next != nil {
		totalPressureRelease += maxRelease
		time -= distanceToNext - 1
		next.openValve()
	}

	return
}

func getStartNode() (start *Valve) {
	for _, valve := range valves {
		if valve.name == "AA" {
			start = valve
			break
		}
	}
	return
}

func addDistances() {
	for _, valve := range valves {
		valve.getDistances()
	}
}

func createConnections(arr []string) {

	for _, s := range arr {
		name := strings.Split(strings.Split(s, ";")[0], " ")[1]
		connInput := strings.Split(s, " ")
		for i := 9; i < len(connInput); i++ {
			connName := strings.Trim(connInput[i], ",")
			valves[name].addConnection(valves[connName])
		}
	}
}

func createValves(arr []string) {
	valves = make(map[string]*Valve, 0)
	for _, s := range arr {
		name := strings.Split(strings.Split(s, ";")[0], " ")[1]
		flow, err := strconv.Atoi(strings.Split(strings.Split(strings.Split(s, ";")[0], " ")[4], "=")[1])
		if err != nil {
			log.Fatalf("could not parse flow: %v", strings.Split(strings.Split(strings.Split(s, ";")[0], " ")[4], "=")[1])
		}
		valves[name] = NewValve(name, flow)
	}
}

// only open valves if flow rate > 0
// check distance to current valve to calc potential release -> travel node with the highest release until time is over
