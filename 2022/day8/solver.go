package day8

import (
	"mrose.de/aoc/utility"
	"strconv"
)

type tree struct {
	top, bot, left, right *tree
	height                int
	visible               bool
	scenicScore           int
}

var maxR int
var maxC int

func Solve() (part1, part2 int) {
	input := utility.Input(2022, 8)
	rootTree := createForest(input)
	part1 = visibleTrees(rootTree)
	part2 = scenicScore(rootTree)
	return
}

func scenicScore(t *tree) int {
	maxScore := 0

	for {
		buildScenicScores(t)
		if t.scenicScore > maxScore {
			maxScore = t.scenicScore
		}
		if t.right != nil {
			t = t.right
		} else {
			for {
				if t.left == nil {
					break
				}
				t = t.left
			}
			if t.bot == nil {
				break
			}
			t = t.bot
		}

	}

	return maxScore
}

func buildScenicScores(t *tree) {
	var currentTree *tree
	var scoreRight, scoreLeft, scoreTop, scoreBot int
	if t.right != nil {
		currentTree = t.right
		for {
			if currentTree.height < t.height {
				scoreRight++
				if currentTree.right == nil {
					break
				}
				currentTree = currentTree.right
				continue
			}
			if t.height <= currentTree.height {
				scoreRight++
				break
			}
		}
	}

	if t.left != nil {
		currentTree = t.left
		for {
			if currentTree.height < t.height {
				scoreLeft++
				if currentTree.left == nil {
					break
				}
				currentTree = currentTree.left
				continue
			}
			if t.height <= currentTree.height {
				scoreLeft++
				break
			}
		}
	}

	if t.bot != nil {
		currentTree = t.bot
		for {
			if currentTree.height < t.height {
				scoreBot++
				if currentTree.bot == nil {
					break
				}
				currentTree = currentTree.bot
				continue
			}
			if t.height <= currentTree.height {
				scoreBot++
				break
			}
		}
	}

	if t.top != nil {
		currentTree = t.top
		for {
			if currentTree.height < t.height {
				scoreTop++
				if currentTree.top == nil {
					break
				}
				currentTree = currentTree.top
				continue
			}
			if t.height <= currentTree.height {
				scoreTop++
				break
			}
		}
	}
	t.scenicScore = scoreBot * scoreTop * scoreLeft * scoreRight
}

func createForest(input string) *tree {
	inputArr := utility.StrArr(input)
	maxR = len(inputArr)
	maxC = len(inputArr[0])

	var rootTree *tree
	topTrees := make([][]*tree, len(inputArr))
	for i, _ := range topTrees {
		topTrees[i] = make([]*tree, len(inputArr[0]))
	}
	for r, line := range inputArr {
		var lastLeft *tree

		for c, h := range line {
			height, _ := strconv.Atoi(string(h))
			currentTree := &tree{
				height: height,
			}
			if lastLeft != nil {
				lastLeft.right = currentTree
				currentTree.left = lastLeft
			}
			if r != 0 {
				topTrees[r-1][c].bot = currentTree
				currentTree.top = topTrees[r-1][c]
			}
			topTrees[r][c] = currentTree
			lastLeft = currentTree
			if c == 0 && r == 0 {
				rootTree = currentTree
			}

		}
	}
	return rootTree
}

func visibleTrees(t *tree) int {
	horizontal(t)
	vertical(t)

	return findVisibleTrees(t)
}

func findVisibleTrees(t *tree) (result int) {
	for i := 0; i < maxR; i++ {
		for j := 0; j < maxC; j++ {
			if t.visible {
				result++
			}
			if j == maxC-1 {
				continue
			}
			t = t.right
		}
		for {
			if t.left != nil {
				t = t.left
			} else {
				break
			}
			t = t.left
		}
		if i == maxR-1 {
			continue
		}
		t = t.bot
	}
	return
}

func horizontal(t *tree) {
	maxHeight := 0
	for i := 0; i < maxR; i++ {
		for j := 0; j < maxC; j++ {
			if t.left == nil {
				t.visible = true
				maxHeight = t.height
				t = t.right
				continue
			}

			if maxHeight < t.height {
				t.visible = true
				maxHeight = t.height
			}

			if j == maxC-1 {
				continue
			}
			t = t.right
		}

		for j := 0; j < maxC; j++ {
			if t.right == nil {
				t.visible = true
				maxHeight = t.height
				t = t.left
				continue
			}

			if maxHeight < t.height {
				t.visible = true
				maxHeight = t.height
			}
			if j == maxC-1 {
				continue
			}
			t = t.left
		}
		if i == maxR-1 {
			continue
		}
		t = t.bot
	}
}

func vertical(t *tree) {
	maxHeight := 0
	for i := 0; i < maxC; i++ {
		for j := 0; j < maxR; j++ {
			if t.top == nil {
				t.visible = true
				maxHeight = t.height
				t = t.bot
				continue
			}

			if maxHeight < t.height {
				t.visible = true
				maxHeight = t.height
			}

			if j == maxR-1 {
				continue
			}
			t = t.bot
		}

		for j := 0; j < maxR; j++ {
			if t.bot == nil {
				t.visible = true
				maxHeight = t.height
				t = t.top
				continue
			}

			if maxHeight < t.height {
				t.visible = true
				maxHeight = t.height
			}
			if j == maxR-1 {
				continue
			}
			t = t.top
		}
		if i == maxC-1 {
			continue
		}
		t = t.right
	}
}
