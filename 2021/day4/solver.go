package day4

import (
	"fmt"
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type Game struct {
	drawList    []string
	Boards      []*Board
	WinnerBoard *Board
}

func Solve() (result int) {
	input := utility.Input2021Day4()
	drawList := strings.Split(strings.Split(input, "\n\n")[0], ",")
	boards := NewBoards(strings.Split(input, "\n\n")[1:])
	game := Game{
		Boards:   boards,
		drawList: drawList,
	}

	return game.Play()
}

func (g *Game) Play() (result int) {
	bingoBoard := &Board{}
	lastValue := ""
	for _, s := range g.drawList {
		for _, board := range g.Boards {
			for _, value := range board.Values {
				if value.Value == s {
					value.Checked = true
				}
			}
		}

		for _, board := range g.Boards {
			bingoBoard = board.isBingo()
			if bingoBoard != nil {
				if !board.Finished {
					board.Finished = true
					g.WinnerBoard = bingoBoard
				}
			}
		}
		if bingoBoard != nil {
			allFinished := true
			for _, board := range g.Boards {
				if !board.Finished {
					allFinished = false
				}
			}
			if allFinished {
				result = getResult(g.WinnerBoard, s)
				return
			}
		}
	}

	return
}

func getResult(board *Board, value string) int {
	sum := 0
	lastNumber, err := strconv.Atoi(value)
	if err != nil {
		_ = fmt.Errorf("could not parse value: %v. error: %e", value, err)
	}
	for _, v := range board.Values {
		if !v.Checked {
			intValue, err := strconv.Atoi(v.Value)
			if err != nil {
				_ = fmt.Errorf("could not parse value: %v. error: %e", value, err)
			}
			sum += intValue
		}
	}
	return sum * lastNumber
}

type BoardValue struct {
	Value   string
	Checked bool
}

type Board struct {
	Values   []*BoardValue
	Rows     [][]*BoardValue
	Cols     [][]*BoardValue
	Finished bool
}

func NewBoard(input string) *Board {
	rowArr := strings.Split(input, "\n")
	board := &Board{
		Rows: make([][]*BoardValue, 0),
		Cols: make([][]*BoardValue, 0),
	}
	initializeRows(board, rowArr)
	initializeCols(board)
	return board
}

func initializeCols(board *Board) {
	for i := 0; i < len(board.Rows[0]); i++ {
		board.Cols = append(board.Cols, make([]*BoardValue, 0))
	}
	for _, row := range board.Rows {
		for i, value := range row {
			board.Cols[i] = append(board.Cols[i], value)
		}
	}
}

func initializeRows(board *Board, rowArr []string) {
	for _, row := range rowArr {
		valueArr := strings.Split(row, " ")
		rowValues := make([]*BoardValue, 0)
		for _, value := range valueArr {
			if value == "" {
				continue
			}
			boardValue := &BoardValue{
				Value:   value,
				Checked: false,
			}
			board.Values = append(board.Values, boardValue)
			rowValues = append(rowValues, boardValue)
		}
		board.Rows = append(board.Rows, rowValues)
	}
}

func NewBoards(input []string) []*Board {
	boards := make([]*Board, 0)
	for _, v := range input {
		boards = append(boards, NewBoard(v))
	}
	return boards
}

func (b *Board) isBingo() (board *Board) {
	bingo := true
	for _, row := range b.Rows {
		bingo = true
		for _, value := range row {
			if !value.Checked {
				bingo = false
				break
			}
		}
		if bingo {
			board = b
		}
	}

	for _, col := range b.Cols {
		bingo = true
		for _, value := range col {
			if !value.Checked {
				bingo = false
				break
			}
		}
		if bingo {
			board = b
		}
	}

	return
}
