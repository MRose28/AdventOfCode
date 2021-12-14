package day14

import (
	"mrose.de/aoc/utility"
	"strings"
)

// Input targetElement when condition is met

type Rule struct {
	Condition     string
	TargetElement string
}

var ruleMap map[string]string

type ElementNode struct {
	Prev  *ElementNode
	Next  *ElementNode
	Value string
}

func NewElementNode(value string) *ElementNode {
	return &ElementNode{
		Value: value,
	}
}

func (e *ElementNode) AddNewNodeAfter(value string) {
	newElement := NewElementNode(value)
	newElement.Next = e.Next
	newElement.Prev = e
	e.Next = newElement
	e.Next.Prev = newElement
}

func Solve() (result int) {
	input := utility.Input2021Day14()
	headElement, _ := NewElementList(strings.Split(input, "\n\n")[0])
	ruleMap = parseRules(strings.Split(input, "\n\n")[1])
	nodeMap := runTemplate(headElement, ruleMap, 40)
	result = getBigResult(nodeMap)
	return
}

func getBigResult(nodeMap map[string]int) int {
	resultMap := make(map[string]int, 0)
	for k, v := range nodeMap {
		resultMap[strings.Split(k, "")[0]] += v
		resultMap[strings.Split(k, "")[1]] += v
	}
	min := 0
	max := 0
	for _, v := range resultMap {
		if min == 0 {
			min = v/2
			if v%2!=0 {
				min+=1
			}
		}
		if min > v/2 {
			min = v/2
			if v%2!=0 {
				min+=1
			}
		}
		if max < v/2 {
			max = v/2
			if v%2!=0 {
				max+=1
			}
		}
	}

	return max - min
}

func NewElementList(s string) (*ElementNode, *ElementNode) {
	var head *ElementNode
	var prev *ElementNode
	var current *ElementNode
	for i := 0; i < len(s); i++ {
		current = NewElementNode(s[i : i+1])
		if prev != nil {
			prev.Next = current
			current.Prev = prev
		} else {
			head = current
		}
		prev = current
	}
	return head, current
}

func runTemplate(element *ElementNode, rules map[string]string, runs int) map[string]int {

	// contains total number of node pairs
	nodeMap := initNodeMap(element)

	for run := 0; run < runs; run++ {
		// contains the amount that should be added per node
		diffMap := make(map[string]int, 0)

		for k := range nodeMap {
			diffMap[k] = nodeMap[k]
		}
		for k, v := range rules {
			nodeMap[k] -= diffMap[k]
			a := strings.Split(k, "")[0] + v
			nodeMap[a] += diffMap[k]
			b := v + strings.Split(k, "")[1]
			nodeMap[b] += diffMap[k]
		}
	}

	return nodeMap
}

func initNodeMap(element *ElementNode) map[string]int {
	resultMap := make(map[string]int, 0)
	for element.Next != nil {
		resultMap[element.Value+element.Next.Value]++
		element = element.Next
	}
	return resultMap
}

func parseRules(s string) (rules map[string]string) {
	arr := utility.StrArr(s)
	rules = make(map[string]string, 0)
	var c string
	var e string
	var splitS []string
	for _, s := range arr {
		splitS = strings.Split(s, " -> ")
		c = splitS[0]
		e = splitS[1]
		rules[c] = e
	}
	return rules
}


// Part One
//func runTemplate(element string, rules []*Rule, runs int) string {
//	for run := 0; run < runs; run++ {
//		for i := 0; i < len(element)-1; i++ {
//			for _, rule := range rules {
//				conditionsString := element[i : i+2]
//				if conditionsString == rule.Condition {
//					element = element[:i+1] + rule.TargetElement + element[i+1:]
//					i++
//					break
//				}
//			}
//		}
//	}
//
//	return element
//}

// Part One
//func getResult(element string) (result int) {
//
//	counterMap := make(map[string]int)
//	for i := 0; i < len(element); i++ {
//		counterMap[element[i:i+1]]++
//	}
//
//	min := 0
//	max := 0
//	for _, v := range counterMap {
//		if min == 0 {
//			min = v
//		}
//		if min > v {
//			min = v
//		}
//		if max < v {
//			max = v
//		}
//	}
//
//	result = max - min
//	return
//}