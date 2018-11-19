package board

import (
	"testing"
	"tictactoe/testhelper"
)

func TestMakeBoardReturnsBoardWithCells(t *testing.T) {
	board := setupEmptyBoard()
	expectedNumOfCells := 9
	actualNumOfCells := len(board.cells)

	matchers.EqualLiterals(t, expectedNumOfCells, actualNumOfCells)
	matchers.EqualLiterals(t, 3, board.size)
}

func TestPutMarkOnEmptyCellMarksCell(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 1)

	matchers.EqualLiterals(t, "X", board.cells[1])
}

func TestMoveIsValidForNonTakenField(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)

	matchers.IsTrue(t, board.IsMoveValid(8))
}

func TestMoveIsInvalidForTakenField(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)

	matchers.IsFalse(t, board.IsMoveValid(2))
}

func TestMoveIsInValidForOutOfRangeIndex(t *testing.T) {
	board := setupEmptyBoard()

	matchers.IsFalse(t, board.IsMoveValid(10))
}

func TestGetRowsReturnsArrayOfCellsSplitIntoRows(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedRows := [][]string{{"Y", "2", "X"}, {"4", "Y", "6"}, {"7", "8", "Y"}}

	matchers.DeepEqual(t, expectedRows, board.GetRows())
}

func TestGetColumnsReturnsArrayOfCellsSplitIntoColumns(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedColumns := [][]string{{"Y", "4", "7"}, {"2", "Y", "8"}, {"X", "6", "Y"}}

	matchers.DeepEqual(t, expectedColumns, board.getColumns())
}

func TestGetLeftDiagonalReturnsLeftDiagonal(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedLeftDiagonal := []string{"Y", "Y", "Y"}

	matchers.DeepEqual(t, expectedLeftDiagonal, board.getLeftDiagonal())
}

func TestGetRightDiagonalReturnsRightDiagonal(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedRightDiagonal := []string{"X", "Y", "7"}

	matchers.DeepEqual(t, expectedRightDiagonal, board.getRightDiagonal())
}

func TestGetDiagonalsReturnsBothDiagonals(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedDiagonals := [][]string{{"Y", "Y", "Y"}, {"X", "Y", "7"}}

	matchers.DeepEqual(t, expectedDiagonals, board.getDiagonals())
}

func TestGetLinesReturnsAllLinesOnBoard(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedAllLines := [][]string{
		{"Y", "2", "X"},
		{"4", "Y", "6"},
		{"7", "8", "Y"},
		{"Y", "4", "7"},
		{"2", "Y", "8"},
		{"X", "6", "Y"},
		{"Y", "Y", "Y"},
		{"X", "Y", "7"},
	}

	matchers.DeepEqual(t, expectedAllLines, board.getLines())
}

func TestIsWinnerAcceptsMarkAndReturnsTrueIfGameWon(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedWinnerMark := "Y"
	expectedLoserMark := "X"

	matchers.IsTrue(t, board.IsWon(expectedWinnerMark))
	matchers.IsFalse(t, board.IsWon(expectedLoserMark))
}

func TestHasEmptyCellReturnsTrueWhenEmptyCellIsAvailable(t *testing.T) {
	board := setupBoardWithPrePickedCells()

	matchers.IsTrue(t, board.HasEmptyCell())
}

func TestReturnsMarkOfTheActivePlayerAtTheBeginningOfTheGame(t *testing.T) {
	board := setupEmptyBoard()

	matchers.EqualLiterals(t, "X", board.GetActivePlayerMark())
}

func TestReturnsMarkOfThePassivePlayerAtTheBeginningOfTheGame(t *testing.T) {
	board := setupEmptyBoard()

	matchers.EqualLiterals(t, "Y", board.GetPassivePlayerMark())
}

func TestReturnsMarkOfTheActivePlayerAfterFirstMove(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)

	matchers.EqualLiterals(t, "Y", board.GetActivePlayerMark())
}

func TestReturnsMarkOfThePassivePlayerAfterFirstMove(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)

	matchers.EqualLiterals(t, "X", board.GetPassivePlayerMark())
}

func TestGameIsFinishedWhenBoardIsWon(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 0)
	board.PutMarkOnBoard("X", 2)
	board.PutMarkOnBoard("X", 1)

	matchers.IsTrue(t, board.IsGameOver("X", "Y"))
}

func TestGameIsNotFinishedWhenEmptyPlacesLeftAndNoWinner(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)
	board.PutMarkOnBoard("Y", 1)

	matchers.IsFalse(t, board.IsGameOver("X", "Y"))
}

func TestGameIsFinishedWhenNoMoreFreePlaces(t *testing.T) {
	board := setupEmptyBoard()
	for i := range board.cells {
		board.PutMarkOnBoard("X", i)
	}

	matchers.IsTrue(t, board.IsGameOver("X", "Y"))
}

func TestBoardReturnsFreePlaces(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)
	board.PutMarkOnBoard("Y", 1)
	actual := board.GetFreeCells()
	expected := []string{"1", "4", "5", "6", "7", "8", "9"}

	matchers.DeepEqual(t, expected, actual)
}

func TestBoardHasTie(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 4)
	board.PutMarkOnBoard("Y", 0)
	board.PutMarkOnBoard("X", 2)
	board.PutMarkOnBoard("Y", 6)
	board.PutMarkOnBoard("X", 3)
	board.PutMarkOnBoard("Y", 5)
	board.PutMarkOnBoard("X", 1)
	board.PutMarkOnBoard("Y", 7)
	board.PutMarkOnBoard("X", 8)

	matchers.IsTrue(t, board.IsTie())
}

func TestBoardClonesItself(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 4)
	boardClone := board.MakeACloneOfItself()

	matchers.DeepEqual(t, board.cells, boardClone.cells)
}

func setupEmptyBoard() Board {
	boardSize := 3
	marksRepo := MarksRepo{"X", "Y"}
	board := MakeBoard(boardSize, &marksRepo)
	return board
}

func setupBoardWithPrePickedCells() Board {
	boardSize := 3
	marksRepo := MarksRepo{"X", "Y"}
	board := MakeBoard(boardSize, &marksRepo)
	board.PutMarkOnBoard("Y", 0)
	board.PutMarkOnBoard("X", 2)
	board.PutMarkOnBoard("Y", 4)
	board.PutMarkOnBoard("Y", 8)
	return board
}
