package day7

import (
	"strconv"
	"strings"
)

//Solver
func SolveStar2(input string) (result int) {
	input = strings.TrimSuffix(input, "\n")
	ruleMap := ruleMap(strings.Split(input, "\n"))
	return amountOfContainedBags(ruleMap, "shiny gold")
}

//Total amount of contained bags
func amountOfContainedBags(ruleMap map[string]map[string]int, ownerName string) int {
	resultMap := allTenants(ruleMap, ownerName)
	result := 0
	for _, v := range resultMap {
		result += v
	}
	return result
}

//Get all tenants and their quantity in a map
func allTenants(ruleMap map[string]map[string]int, name string) map[string]int {
	changed := true
	currentTenants := make([]string, 0)
	futureTenants := make([]string, 0)
	lowerTenants := make(map[string]int, 0)
	currentTenants = append(currentTenants, name)
	resultMap := make(map[string]int, 0)

	for changed {
		if len(currentTenants) == 0 {
			changed = false
		}
		for _, tenantName := range currentTenants {
			lowerTenants = tenants(ruleMap, tenantName)
			if lowerTenants != nil {
				for k, v := range lowerTenants {
					for i:=0; i < v; i++ {
						futureTenants = append(futureTenants, k)
					}
					resultMap[k] += v
				}
			}
		}
		currentTenants = futureTenants
		futureTenants = make([]string, 0)
	}
	return resultMap
}

//A map containing the rules
func ruleMap(input []string) map[string]map[string]int {
	ruleMap := make(map[string]map[string]int, 0)
	for _, line := range input {
		owner, tenants := parseRule(line)
		ruleMap[owner] = tenants
	}
	return ruleMap
}

//Get the tenants of one owner
func tenants(ruleMap map[string]map[string]int, owner string) map[string]int {
	return ruleMap[owner]
}

//Parse a single rule from the ruleMap
func parseRule(line string) (string, map[string]int) {
	owner := strings.Split(line, " bags contain ")[0]
	tenants := make(map[string]int, 0)

	for _, tenant := range strings.Split(strings.Split(line, " bags contain ")[1], ", ") {
		tenantArr := strings.Split(tenant, " ")
		if tenantArr[0] != "no" {
			tenants[tenantArr[1]+" "+tenantArr[2]], _ = strconv.Atoi(tenantArr[0])
		}
	}
	return owner, tenants
}
