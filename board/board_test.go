package board

import (
	"testing"
	"tictactoe/testHelper"
)

func createBoard(cells [9]string) Board{
	return Board{cells: cells}
}

func TestBoardHasCells(t *testing.T) {
	cells := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	aBoard := createBoard(cells)
	expectedNumOfCells := 9
	actualNumOfCells := len(aBoard.cells)

	matchers.EqualLiterals(t, expectedNumOfCells, actualNumOfCells)
}

func TestPutMarkOnEmptyCell(t *testing.T) {
	cells := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	aBoard := createBoard(cells)
	cellIndex := 0
	aBoard.putMarkOnBoard("X", cellIndex)

	matchers.EqualLiterals(t, aBoard.cells[cellIndex], "X")
}  