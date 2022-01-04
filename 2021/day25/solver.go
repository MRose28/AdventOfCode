package day25
//
//import "mrose.de/aoc/utility"
//
//const (
//	EMPTY Kind = iota
//	EAST
//	SOUTH
//)
//
//type Kind int
//
//type Cucumber struct {
//	Kind               Kind
//	BotNext, RightNext *Cucumber
//	Moved              bool
//}
//
//func Solve() (result int) {
//	inputArr := utility.StrArr(utility.Input(2021, 25))
//	cucumbers := parseCucumbers(inputArr)
//	return simulateMovement(cucumbers)
//}
//
//func simulateMovement(cucumbers [][]*Cucumber) int {
//	steps := 0
//	for simulateStep(cucumbers) {
//		steps++
//		resetChangedStates(cucumbers)
//	}
//	return steps
//}
//
//func resetChangedStates(cucumbers [][]*Cucumber) {
//
//	for _, row := range cucumbers {
//		for _, cucumber := range row {
//			cucumber.Moved = false
//		}
//	}
//}
//
//func simulateStep(cucumbers [][]*Cucumber) (changed bool) {
//	for rowIndex := 0; rowIndex < len(cucumbers); rowIndex++ {
//		row := cucumbers[rowIndex]
//		cucumber := cucumbers[rowIndex][0]
//		for i := 0; i < len(row); i++ {
//			if cucumber.Kind != EAST {
//				continue
//			}
//			if cucumber.Moved {
//				continue
//			}
//			if cucumber.RightNext.Kind == EMPTY {
//				cucumber.RightNext.Kind = EAST
//				cucumber.Kind = EMPTY
//				changed = true
//				cucumber.RightNext.Moved = true
//			}
//			cucumber = cucumber.RightNext
//		}
//	}
//
//	for rowIndex := 0; rowIndex < len(cucumbers); rowIndex++ {
//		row := cucumbers[rowIndex]
//		cucumber := cucumbers[rowIndex][0]
//		for i := 0; i < len(row); i++ {
//			if cucumber.Kind != SOUTH {
//				continue
//			}
//			if cucumber.Moved {
//				continue
//			}
//			if cucumber.BotNext.Kind == EMPTY {
//				cucumber.BotNext.Kind = SOUTH
//				cucumber.Kind = EMPTY
//				changed = true
//				cucumber.BotNext.Moved = true
//			}
//		}
//	}
//	return
//}
//
//func parseCucumbers(input []string) []*Cucumber {
//	cucumberRowHeads := make([]*Cucumber, len(input))
//	cucumberColHeads := make([]*Cucumber, len(input[0]))
//	prevTopRow := make([]*Cucumber, len(input[0]))
//	for rowIndex, row := range input {
//		var left *Cucumber
//		var firstInRow *Cucumber
//		for i := 0; i < len(row); i++ {
//			value := row[i : i+1]
//			newCucumber := &Cucumber{}
//			switch value {
//			case "v":
//				newCucumber.Kind = SOUTH
//			case ">":
//				newCucumber.Kind = EAST
//			case ".":
//				newCucumber.Kind = EMPTY
//			}
//
//			if i == len(row)-1 {
//				newCucumber.RightNext = firstInRow
//			}
//			if left != nil {
//				left.RightNext = newCucumber
//			}
//			left = newCucumber
//			if i == 0 {
//				firstInRow = newCucumber
//			}
//			if rowIndex == 0 {
//				cucumberColHeads = append(cucumberColHeads, newCucumber)
//			}
//		}
//		cucumberRowHeads[rowIndex] = firstInRow
//	}
//
//	var prevTop *Cucumber
//	var current *Cucumber
//	for rowIndex, head := range cucumberRowHeads {
//		current = head
//		for colIndex := 0; colIndex < len(input[0]); colIndex++ {
//			if rowIndex < len(cucumberRowHeads)-1 {
//
//			}
//			current = current.RightNext
//		}
//	}
//	return cucumberRowHeads
//}
