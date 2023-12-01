package day12

//
//import "mrose.de/aoc/utility"
//
//type Tree struct {
//	x, y   int
//	height string
//	blocked int
//}
//
//var alphabet map[string]int
//var forest map[*Tree]int
//var startTree *Tree
//
//func Solve() (part1, part2 int) {
//
//	inputArr := utility.StrArr(utility.TestInput(2022, 12))
//	createAlphabet()
//	createForest(inputArr)
//	part1 = shortestPath()
//	return
//}
//
//func shortestPath() int {
//
//}
//
//func createForest(arr []string) {
//	forest = make(map[*Tree]int, 0)
//	for y, line := range arr {
//		for x, r := range line {
//			newTree := &Tree{
//				x:      x,
//				y:      y,
//				height: string(r),
//			}
//			forest[newTree] = alphabet[newTree.height]
//			if newTree.height == "S" {
//				startTree = newTree
//			}
//		}
//	}
//}
//
//func createAlphabet() {
//	alphabetChars := "abcdefghijklmnopqrstuwvxyz"
//	alphabet = make(map[string]int, 0)
//	for i, r := range alphabetChars {
//		alphabet[string(r)] = i
//	}
//}
