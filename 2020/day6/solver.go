package day6

import (
	"strings"
)

//Get the result to check against AOC
func Solve(input string) int {
	groups := groups(input)
	result := 0
	for _, group := range groups {
		group = strings.TrimSuffix(group, "\n")
		result += getGroupResult(group)
	}
	return result
}

//Separate the string blocks that represent groups
func groups(input string) []string {
	var groups []string
	groups = append(groups, strings.Split(input, "\n\n")...)
	return groups
}

//Return how many options have been answered with yes withing the given group.
func getGroupResult(group string) int {
	alphabet := "abcdefghijklmnopqrstuwvxyz"
	results := initializeMap()
	members := strings.Split(group, "\n")
	for _, member := range members {
		for _, letter := range alphabet {
			if strings.Contains(member, string(letter)) {
				results[string(letter)] += 1
			}
		}
	}

	//Star 2
	for _, letter := range alphabet {
		if results[string(letter)] == len(members) {
			results[string(letter)] = 1
		} else {
			results[string(letter)] = 0
		}

	}

	return getResultFromMap(results)
}

//Given a map representing all answers by a group, this returns the number of options that were chosen by every member.
func getResultFromMap(results map[string]int) int {
	result := 0
	for k := range results {
		if results[k] > 0 {
			result++
		}
	}
	return result
}

//initialize a Mao with keys representing the characters 'a-z'
func initializeMap() map[string]int {
	results := make(map[string]int)
	alphabet := "abcdefghijklmnopqrstuwvxyz"
	for _, letter := range alphabet {
		results[string(letter)] = 0
	}
	return results
}
