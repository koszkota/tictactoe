package board

import (
	"testing"
	"tictactoe/testHelper"
)

func TestMakeBoardReturnsBoardWithCells(t *testing.T) {
	aBoard := MakeBoard(3)
	expectedNumOfCells := 9
	actualNumOfCells := len(aBoard.cells)

	matchers.EqualLiterals(t, expectedNumOfCells, actualNumOfCells)
}

func TestPutMarkOnEmptyCell(t *testing.T) {
	aBoard := MakeBoard(3)
	cellIndex := 0
	aBoard.putMarkOnBoard("X", cellIndex)

	matchers.EqualLiterals(t, aBoard.cells[cellIndex], "X")
}

func TestCheckIfCellIsTakenOrEmpty(t *testing.T) {
	aBoard := MakeBoard(3)
	cellIndex := 0
	aBoard.putMarkOnBoard("X", cellIndex)

	matchers.IsTrue(t, aBoard.cellIsAvailable(1))
	matchers.IsFalse(t, aBoard.cellIsAvailable(cellIndex))
}