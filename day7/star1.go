package day7

import (
	"strings"
)

func SolveStar1(input string) (result int) {
	input = strings.TrimSuffix(input, "\n")
	bagsContainingGold := make([]string, 0)

	//add bags that contain 'shiny gold' bags directly
	for _, line := range strings.Split(input, "\n") {
		bagDesc := strings.Replace(strings.Split(line, "contain")[0], " bags ", "", -1)
		if strings.Contains(strings.Split(
			line, "contain")[1], "shiny gold") && !contains(bagsContainingGold, bagDesc) {
			bagsContainingGold = append(bagsContainingGold, bagDesc)
		}
	}

	//add bags that contain 'shiny gold' indirectly
	changed := true
	for changed {
		changed, bagsContainingGold = updateGoldContainingBags(input, bagsContainingGold)
	}

	return len(bagsContainingGold)
}

func contains(s []string, searchTerm string) (contained bool) {
	for _, value := range s {
		if value == searchTerm {
			return true
		}
	}
	return false
}

func updateGoldContainingBags(input string, bagsContainingGold []string) (bool, []string) {
	changed := false
	for _, line := range strings.Split(input, "\n") {
		containedBags := strings.Trim(strings.Split(line, "contain")[1], " ")
		containingBag := strings.Replace(strings.Split(line, "contain")[0], " bags ", "", -1)
		for _, goldBag := range bagsContainingGold {
			if strings.Contains(containedBags, goldBag) {
				if contains(bagsContainingGold, containingBag) == false {
					bagsContainingGold = append(bagsContainingGold, containingBag)
					changed = true
				}
			}
		}
	}
	return changed, bagsContainingGold
}
