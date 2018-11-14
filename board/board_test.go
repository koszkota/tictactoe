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

func TestPutMarkOnEmptyCell(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 1)

	matchers.EqualLiterals(t, board.cells[1], "X")
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
	board.PutMarkOnBoard("X", 2)

	matchers.IsFalse(t, board.IsMoveValid(10))
}

func TestGetRowsReturnsArrayOfCellsSplitIntoRows(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedRows := [][]string{{"Y","2","X"}, {"4", "Y", "6"}, {"7", "8", "Y"}}

	matchers.DeepEqual(t, expectedRows, board.GetRows())
}

func TestGetColumnsReturnsArrayOfCellsSplitIntoColumns(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedColumns := [][]string{{"Y","4","7"}, {"2", "Y", "8"}, {"X", "6", "Y"}}

	matchers.DeepEqual(t, expectedColumns, board.getColumns())
}

func TestGetLeftDiagonal(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedLeftDiagonal := []string{"Y","Y","Y"}

	matchers.DeepEqual(t, expectedLeftDiagonal, board.getLeftDiagonal())
}

func TestGetRightDiagonal(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedRightDiagonal := []string{"X","Y","7"}

	matchers.DeepEqual(t, expectedRightDiagonal, board.getRightDiagonal())
}

func TestGetDiagonals(t *testing.T) {
	board := setupBoardWithPrePickedCells()
	expectedDiagonals := [][]string{{"Y","Y","Y"}, {"X", "Y", "7"}}

	matchers.DeepEqual(t, expectedDiagonals, board.getDiagonals())
}

func TestGetLinesReturnsAllLinesOnBoard(t *testing.T) {
	board := setupBoardWithPrePickedCells()
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

	matchers.DeepEqual(t, expectedAllLines, board.getLines())
}

func TestIsWinnerAcceptsMarkAndReturnsBool(t *testing.T) {
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

	matchers.EqualLiterals(t, board.GetActivePlayerMark(), "X")
}

func TestReturnsMarkOfThePassivePlayerAtTheBeginningOfTheGame(t *testing.T) {
	board := setupEmptyBoard()

	matchers.EqualLiterals(t, board.GetPassivePlayerMark(), "Y")
}

func TestReturnsMarkOfTheActivePlayerAfterFirstMove(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)

	matchers.EqualLiterals(t, board.GetActivePlayerMark(), "Y")
}

func TestReturnsMarkOfThePassivePlayerAfterFirstMove(t *testing.T) {
	board := setupEmptyBoard()
	board.PutMarkOnBoard("X", 2)

	matchers.EqualLiterals(t, board.GetPassivePlayerMark(), "X")
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

	matchers.DeepEqual(t, actual, expected)
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
	markOnOriginalBoard := board.cells[4]
	markOnClonedBoard := boardClone.cells[4]

	matchers.EqualLiterals(t, markOnOriginalBoard, markOnClonedBoard)
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