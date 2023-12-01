package day12

import (
	"mrose.de/aoc/utility"
	"strings"
)

var steps int
var alphabet = "abcdefghijklmnopqrstuwvxyz"

var path []*node

func Solve() (part1, part2 int) {

	input := utility.StrArr(utility.TestInput(2022, 12))

	start := createMap(input)
	traverse(start)
	part1 = steps
	return
}

func traverse(start *node) {
	current := start
	path = make([]*node, 0)
	addNodeToPath(start)
	for {
		handleNeighbours(current)
		current.visited = true
	}
}

func handleNeighbours(current *node) {
	for _, n := range current.neighbours() {
		handleNeighbour(current, n)
	}
}

func handleNeighbour(current *node, neighbour *node) {
}

func getLastNode() *node {
	return path[len(path)-1]
}

func addNodeToPath(n *node) {
	path = append(path, n)
}

func removeLastNodeFromPath() {
	path = path[0 : len(path)-1]
}

func heightOK(c, t *node) bool {
	return t.height-c.height <= 1
}

var maxR int
var maxC int

func createMap(input []string) (start *node) {
	maxR = len(input)
	maxC = len(input[0])
	topRows := make([][]*node, maxR)
	for i, _ := range topRows {
		topRows[i] = make([]*node, maxC)
	}

	for r, line := range input {
		var lastLeft *node

		for c, h := range line {
			var height int
			if h != 'E' && h != 'S' {
				height = strings.Index(alphabet, string(h))
			}
			currentNode := &node{
				height: height,
			}
			if lastLeft != nil {
				lastLeft.right = currentNode
				currentNode.left = lastLeft
			}
			if r != 0 {
				topRows[r-1][c].down = currentNode
				currentNode.up = topRows[r-1][c]
			}
			if h == 'E' {
				currentNode.target = true
			}
			topRows[r][c] = currentNode
			lastLeft = currentNode
			if h == 'S' {
				start = currentNode
				start.steps = 0
			}
		}
	}
	return
}
