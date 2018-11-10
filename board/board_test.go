package board

import (
	"testing"
	"tictactoe/testHelper"
)

func TestMakeBoardReturnsBoardWithCells(t *testing.T) {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	expectedNumOfCells := 9
	actualNumOfCells := len(aBoard.cells)

	matchers.EqualLiterals(t, expectedNumOfCells, actualNumOfCells)
	matchers.EqualLiterals(t, boardSize, aBoard.size)
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

	matchers.IsTrue(t, aBoard.isCellAvailable(1))
	matchers.IsFalse(t, aBoard.isCellAvailable(cellIndex))
}

func TestGetRowsReturnsArrayOfCellsSplitIntoRows(t *testing.T) {
	aBoard := MakeBoard(3)
	aBoard.putMarkOnBoard("X", 1)
	aBoard.putMarkOnBoard("Y", 4)
	expected := [][]string{{" ","X"," "}, {" ", "Y", " "}, {" ", " ", " "}}

	matchers.DeepEqual(t, expected, aBoard.getRows())
}

func TestGetColumnsReturnsArrayOfCellsSplitIntoColumns(t *testing.T) {
	aBoard := MakeBoard(3)
	aBoard.putMarkOnBoard("X", 1)
	aBoard.putMarkOnBoard("Y", 4)
	expected := [][]string{{" "," "," "}, {"X", "Y", " "}, {" ", " ", " "}}

	matchers.DeepEqual(t, expected, aBoard.getColumns())
}

func TestGetLeftDiagonal(t *testing.T) {
	aBoard := MakeBoard(3)
	aBoard.putMarkOnBoard("X", 0)
	expected := []string{"X"," "," "}

	matchers.DeepEqual(t, expected, aBoard.getLeftDiagonal())
}

func TestGetRightDiagonal(t *testing.T) {
	aBoard := MakeBoard(3)
	aBoard.putMarkOnBoard("X", 2)
	aBoard.putMarkOnBoard("Y", 4)
	expected := []string{"X","Y"," "}

	matchers.DeepEqual(t, expected, aBoard.getRightDiagonal())
}

func TestGetDiagonals(t *testing.T) {
	aBoard := MakeBoard(3)
	aBoard.putMarkOnBoard("Y", 0)
	aBoard.putMarkOnBoard("X", 2)
	aBoard.putMarkOnBoard("Y", 4)
	expected := [][]string{{"Y","Y"," "}, {"X", "Y", " "}}

	matchers.DeepEqual(t, expected, aBoard.getDiagonals())
}
