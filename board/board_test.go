package board

import (
	"testing"
	"tictactoe/testhelper"
)

var aBoard = setupBoard()

func setupBoard() Board {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	aBoard.putMarkOnBoard("Y", 0)
	aBoard.putMarkOnBoard("X", 2)
	aBoard.putMarkOnBoard("Y", 4)
	aBoard.putMarkOnBoard("Y", 8)
	return aBoard
}

func TestMakeBoardReturnsBoardWithCells(t *testing.T) {
	expectedNumOfCellsInBoard := 9
	actualNumOfCells := len(aBoard.cells)

	matchers.EqualLiterals(t, expectedNumOfCellsInBoard, actualNumOfCells)
	matchers.EqualLiterals(t, 3, aBoard.size)
}

func TestPutMarkOnEmptyCell(t *testing.T) {
	aBoard := MakeBoard(3)
	aBoard.putMarkOnBoard("X", 1)

	matchers.EqualLiterals(t, aBoard.cells[1], "X")
}

func TestMoveIsValidForNonTakenField(t *testing.T) {
	matchers.IsTrue(t, aBoard.isMoveValid(1))
}

func TestMoveIsInValidForTakenField(t *testing.T) {
	matchers.IsFalse(t, aBoard.isMoveValid(8))
}

func TestMoveIsInValidForOutOfRangeIndex(t *testing.T) {
	matchers.IsFalse(t, aBoard.isMoveValid(10))
}

func TestGetRowsReturnsArrayOfCellsSplitIntoRows(t *testing.T) {
	expectedRows := [][]string{{"Y","2","X"}, {"4", "Y", "6"}, {"7", "8", "Y"}}

	matchers.DeepEqual(t, expectedRows, aBoard.getRows())
}

func TestGetColumnsReturnsArrayOfCellsSplitIntoColumns(t *testing.T) {
	expectedColumns := [][]string{{"Y","4","7"}, {"2", "Y", "8"}, {"X", "6", "Y"}}

	matchers.DeepEqual(t, expectedColumns, aBoard.getColumns())
}

func TestGetLeftDiagonal(t *testing.T) {
	expectedLeftDiagonal := []string{"Y","Y","Y"}

	matchers.DeepEqual(t, expectedLeftDiagonal, aBoard.getLeftDiagonal())
}

func TestGetRightDiagonal(t *testing.T) {
	expectedRightDiagonal := []string{"X","Y","7"}

	matchers.DeepEqual(t, expectedRightDiagonal, aBoard.getRightDiagonal())
}

func TestGetDiagonals(t *testing.T) {
	expectedDiagonals := [][]string{{"Y","Y","Y"}, {"X", "Y", "7"}}

	matchers.DeepEqual(t, expectedDiagonals, aBoard.getDiagonals())
}

func TestGetLinesReturnsAllLinesOnBoard(t *testing.T) {
	expectedAllLines := [][]string{
		{"Y","2","X"},
		{"4", "Y", "6"},
		{"7", "8", "Y"},
		{"Y", "4", "7"},
		{"2", "Y", "8"},
		{"X", "6", "Y"},
		{"Y", "Y", "Y"},
		{"X", "Y", "7"},
	}

	matchers.DeepEqual(t, expectedAllLines, aBoard.getLines())
}

func TestIsWinnerAcceptsMarkAndReturnsBool(t *testing.T) {
	expectedWinnerSign := "Y"
	expectedLoserSign := "X"

	matchers.IsTrue(t, aBoard.isWon(expectedWinnerSign))
	matchers.IsFalse(t, aBoard.isWon(expectedLoserSign))
}

func TestHasEmptyCellReturnsTrueWhenEmptyCellIsAvailable(t *testing.T) {
	matchers.IsTrue(t, aBoard.hasEmptyCell())
}