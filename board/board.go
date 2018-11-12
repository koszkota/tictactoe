package board

import (
	"strconv"
)

type Board struct {
	cells []string
	size int
}

func MakeBoard(boardSize int) Board {
	aBoard := Board{size: boardSize}

	aBoard.initializeEmptyCells(boardSize)
	return aBoard
}

func (board *Board) initializeEmptyCells(boardSize int) {
	for i := 0; i < (boardSize * boardSize); i++ {
		board.cells = append(board.cells, strconv.Itoa(i + 1))
	}
}

func (board *Board) PutMarkOnBoard (mark string, cellIndex int) {
	board.cells[cellIndex] = mark
}

func (board Board) GetRows() [][]string {
	var cellsSplitIntoRows [][]string
	for i := 0; i < len(board.cells); i += board.size {
		end := i + board.size
		cellsSplitIntoRows = append(cellsSplitIntoRows, board.cells[i:end])
	}
	return cellsSplitIntoRows
}

func (board Board) isMoveValid(cellIndex int) bool {
	return board.isCellIndexWithinBoardSize(cellIndex) && board.isCellNumeric(cellIndex)
}

func (board Board) isCellIndexWithinBoardSize(cellIndex int) bool {
	return cellIndex < len(board.cells) && cellIndex > 0
}

func (board Board) isCellNumeric(cellIndex int) bool {
	if _, err := strconv.Atoi(board.cells[cellIndex]); err == nil {
		return true
	}
	return false
}

func (board Board) getColumns() [][]string {
	var cellsSplitIntoColumns [][]string
	rows := board.GetRows()
	for i := 0; i < board.size; i++ {
		var singleColumn []string
		for j := 0; j < len(rows); j++ {
			singleColumn = append(singleColumn, rows[j][i])
		}
		cellsSplitIntoColumns = append(cellsSplitIntoColumns, singleColumn)
	}
	return cellsSplitIntoColumns
}

func (board Board) getDiagonals() [][]string {
	var diagonals [][]string
	diagonals = append(diagonals, board.getLeftDiagonal())
	diagonals = append(diagonals, board.getRightDiagonal())
	return diagonals
}

func (board Board) getLeftDiagonal() []string {
	var leftDiagonal []string
	rows := board.GetRows()
	for i := 0; i < board.size; i++ {
		leftDiagonal = append(leftDiagonal, rows[i][i])
	}
	return leftDiagonal
}
func (board Board) getRightDiagonal() []string {
	var rightDiagonal []string
	rows := board.GetRows()
	endIndex := board.size - 1
	for i := 0; i < board.size; i++ {
		rightDiagonal = append(rightDiagonal, rows[i][endIndex])
		endIndex -= 1
	}
	return rightDiagonal
}

func (board Board) getLines() [][]string {
	return append(append(board.GetRows(), board.getColumns()...), board.getDiagonals()...)
}

func (board Board) isWon(mark string) bool {
	var allLines = board.getLines()
	for _, line := range allLines {
		if lineIsWon(line, mark) {
			return true
		}
	}
	return false
}

func lineIsWon(line []string, mark string) bool{
	for i := 0; i < len(line); i++ {
		if line[i] != mark {
			return false
		}
	}
	return true
}

func (board Board) hasEmptyCell() bool {
	for i := range board.cells {
		if board.isCellNumeric(i) {
			return true
		}
	}
	return false
}