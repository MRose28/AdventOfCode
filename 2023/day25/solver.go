package day25

import (
	"mrose.de/aoc/utility"
	"strings"
)

func Solve() (p1, p2 int) {
	input := utility.InputAsStrArr(2023, 25, true)
	components := parseComponents(input)
	p1 = len(components)

	return
}

// rzs: qnr cmg lsr rsh
func parseComponents(input []string) map[string][]string {
	compMap := make(map[string][]string)
	for _, line := range input {
		currentName := strings.Split(line, ":")[0]
		otherNames := strings.Split(strings.Split(line, ": ")[1], " ")
		if compMap[currentName] == nil {
			compMap[currentName] = make([]string, 0)
		}

		for _, name := range otherNames {
			if compMap[name] == nil {
				compMap[name] = make([]string, 0)
			}
			if !utility.ContainsString(compMap[currentName], name) {
				compMap[currentName] = append(compMap[currentName], name)
			}
			if !utility.ContainsString(compMap[name], currentName) {
				compMap[name] = append(compMap[name], currentName)
			}
		}
	}
	return compMap
}
