package day21

import (
	"mrose.de/aoc/utility"
)

type NodeType string

const (
	garden NodeType = "."
	rock   NodeType = "#"
)

type Node struct {
	Type       NodeType
	N, E, S, W *Node
	start      bool
}

var remainingSteps = 64

func Solve() (p1, p2 int) {

	input := utility.StrArr(utility.Input(2023, 21))

	startNode := parseNodes(input)
	p1 = numOfPossibleNodes(startNode)

	return
}

func numOfPossibleNodes(node *Node) int {
	currentNodes := []*Node{node}
	for remainingSteps != 0 {
		currentNodes = possibilities(currentNodes)
		remainingSteps--
	}
	return len(currentNodes)
}

func parseNodes(input []string) *Node {
	var start, current, prev *Node
	prevLine := make([]*Node, len(input[0]), len(input[0]))
	currentLine := make([]*Node, len(input[0]), len(input[0]))

	for _, line := range input {
		for j, r := range line {
			prev = current
			current = &Node{
				Type:  nodeType(r),
				N:     prevLine[j],
				W:     prev,
				start: r == 'S',
			}
			if prev != nil {
				prev.E = current
			}
			if prevLine[j] != nil {
				prevLine[j].S = current
			}
			currentLine[j] = current
			if current.start {
				start = current
			}
		}
		prevLine = currentLine
		currentLine = make([]*Node, len(input[0]), len(input[0]))
	}

	return start
}

func parseNodesWithInfMap(input []string) *Node {
	var start, current, prev *Node
	prevLine := make([]*Node, len(input[0]), len(input[0]))
	currentLine := make([]*Node, len(input[0]), len(input[0]))
	borderN := make([]*Node, len(input[0]), len(input[0]))
	borderE := make([]*Node, len(input), len(input))
	borderS := make([]*Node, len(input[0]), len(input[0]))
	borderW := make([]*Node, len(input), len(input))
	for y, line := range input {
		for x, r := range line {
			prev = current
			current = &Node{
				Type:  nodeType(r),
				N:     prevLine[x],
				W:     prev,
				start: r == 'S',
			}
			if prev != nil {
				prev.E = current
			}
			if prevLine[x] != nil {
				prevLine[x].S = current
			}
			currentLine[x] = current
			if current.start {
				start = current
			}

			if y == 0 {
				borderN[x] = current
			}
			if y == len(input)-1 {
				borderS[x] = current
			}
			if x == 0 {
				borderW[y] = current
			}
			if x == len(line)-1 {
				borderE[y] = current
			}
		}
		prevLine = currentLine
		currentLine = make([]*Node, len(input[0]), len(input[0]))
	}

	for i, node := range borderN {
		node.N = borderS[i]
		borderS[i] = node
	}

	for i, node := range borderE {
		node.E = borderW[i]
		borderW[i] = node
	}

	return start
}

func nodeType(r rune) NodeType {
	if r == '.' {
		return garden
	}
	if r == '#' {
		return rock
	}

	return garden
}

func possibilities(nodes []*Node) []*Node {
	p := make([]*Node, 0)
	for _, node := range nodes {
		if node.N.Type == garden {
			if !containsNode(p, node.N) {
				p = append(p, node.N)
			}
		}
		if node.E.Type == garden {
			if !containsNode(p, node.E) {
				p = append(p, node.E)
			}
		}
		if node.S.Type == garden {
			if !containsNode(p, node.S) {
				p = append(p, node.S)
			}
		}
		if node.W.Type == garden {
			if !containsNode(p, node.W) {
				p = append(p, node.W)
			}
		}
	}
	return p
}

func containsNode(nodes []*Node, current *Node) bool {
	for _, node := range nodes {
		if node == current {
			return true
		}
	}
	return false
}

/*
--- Day 21: Step Counter ---
You manage to catch the airship right as it's dropping someone else off on their all-expenses-paid trip to Desert Island! It even helpfully drops you off near the gardener and his massive farm.

"You got the sand flowing again! Great work! Now we just need to wait until we have enough sand to filter the water for Snow Island and we'll have snow again in no time."

While you wait, one of the Elves that works with the gardener heard how good you are at solving problems and would like your help. He needs to get his steps in for the day, and so he'd like to know which garden plots he can reach with exactly his remaining 64 steps.

He gives you an up-to-date map (your puzzle input) of his starting position (S), garden plots (.), and rocks (#). For example:

...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
The Elf starts at the starting position (S) which also counts as a garden plot. Then, he can take one step north, south, east, or west, but only onto tiles that are garden plots. This would allow him to reach any of the tiles marked O:

...........
.....###.#.
.###.##..#.
..#.#...#..
....#O#....
.##.OS####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
Then, he takes a second step. Since at this point he could be at either tile marked O, his second step would allow him to reach any garden plot that is one step north, south, east, or west of any tile that he could have reached after the first step:

...........
.....###.#.
.###.##..#.
..#.#O..#..
....#.#....
.##O.O####.
.##.O#...#.
.......##..
.##.#.####.
.##..##.##.
...........
After two steps, he could be at any of the tiles marked O above, including the starting position (either by going north-then-south or by going west-then-east).

A single third step leads to even more possibilities:

...........
.....###.#.
.###.##..#.
..#.#.O.#..
...O#O#....
.##.OS####.
.##O.#...#.
....O..##..
.##.#.####.
.##..##.##.
...........
He will continue like this until his steps for the day have been exhausted. After a total of 6 steps, he could reach any of the garden plots marked O:

...........
.....###.#.
.###.##.O#.
.O#O#O.O#..
O.O.#.#.O..
.##O.O####.
.##.O#O..#.
.O.O.O.##..
.##.#.####.
.##O.##.##.
...........
In this example, if the Elf's goal was to get exactly 6 more steps today, he could use them to reach any of 16 garden plots.

However, the Elf actually needs to get 64 steps today, and the map he's handed you is much larger than the example map.

Starting from the garden plot marked S on your map, how many garden plots could the Elf reach in exactly 64 steps?
*/
