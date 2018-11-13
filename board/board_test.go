package board

import (
	"testing"
	"tictactoe/testhelper"
)

var aBoard = setupBoard()

func setupBoard() Board {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	aBoard.PutMarkOnBoard("Y", 0)
	aBoard.PutMarkOnBoard("X", 2)
	aBoard.PutMarkOnBoard("Y", 4)
	aBoard.PutMarkOnBoard("Y", 8)
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
	aBoard.PutMarkOnBoard("X", 1)

	matchers.EqualLiterals(t, aBoard.cells[1], "X")
}

func TestMoveIsValidForNonTakenField(t *testing.T) {
	matchers.IsTrue(t, aBoard.IsMoveValid(1))
}

func TestMoveIsInValidForTakenField(t *testing.T) {
	matchers.IsFalse(t, aBoard.IsMoveValid(8))
}

func TestMoveIsInValidForOutOfRangeIndex(t *testing.T) {
	matchers.IsFalse(t, aBoard.IsMoveValid(10))
}

func TestGetRowsReturnsArrayOfCellsSplitIntoRows(t *testing.T) {
	expectedRows := [][]string{{"Y","2","X"}, {"4", "Y", "6"}, {"7", "8", "Y"}}

	matchers.DeepEqual(t, expectedRows, aBoard.GetRows())
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

	matchers.IsTrue(t, aBoard.IsWon(expectedWinnerSign))
	matchers.IsFalse(t, aBoard.IsWon(expectedLoserSign))
}

func TestHasEmptyCellReturnsTrueWhenEmptyCellIsAvailable(t *testing.T) {
	matchers.IsTrue(t, aBoard.HasEmptyCell())
}

func TestReturnsMarkOfTheCurrentPlayerAtTheBeginningOfTheGame(t *testing.T) {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	matchers.EqualLiterals(t, aBoard.GetMarkWithLessEntries("X", "Y"), "X")
}

func TestReturnsMarkOfTheCurrentPlayerAfterFirstMove(t *testing.T) {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	aBoard.PutMarkOnBoard("X", 2)
	matchers.EqualLiterals(t, aBoard.GetMarkWithLessEntries("X", "Y"), "Y")
}

func TestGameIsFinishedWhenBoardIsWon(t *testing.T) {
	matchers.IsTrue(t, aBoard.IsGameOver("X", "Y"))
}

func TestGameIsNotFinishedWHenEmptyPlacesAndNoWinner(t *testing.T) {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	aBoard.PutMarkOnBoard("Y", 0)
	matchers.IsFalse(t, aBoard.IsGameOver("X", "Y"))
}

func TestGameIsFinishedWhenNoMoreFreePlaces(t *testing.T) {
	boardSize := 3
	aBoard := MakeBoard(boardSize)
	for i := range aBoard.cells {
		aBoard.PutMarkOnBoard("X", i)
	}
	matchers.IsTrue(t, aBoard.IsGameOver("X", "Y"))
}

func TestBoardReturnsFreePlaces(t *testing.T) {
	actual := aBoard.GetFreeCells()
	expected := []string{"2", "4", "6", "7", "8"}

	matchers.DeepEqual(t, actual, expected)
}