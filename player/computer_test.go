package player

import (
	"testing"
	"tictactoe/board"
	"tictactoe/testhelper"
)

var computerPlayer = Computer{Mark: "Y"}

func TestComputerPlayerHasMark(t *testing.T) {
	matchers.EqualLiterals(t, "Y", computerPlayer.Mark)
}

func TestGetMarkReturnsComputerMark(t *testing.T) {
	matchers.EqualLiterals(t, "Y", computerPlayer.GetMark())
}

func TestPickMethodReturnsWinningPosition(t *testing.T) {
	marksRepo := board.MarksRepo{"X", "Y"}
	aBoard := board.MakeBoard(3, &marksRepo)
	aBoard.PutMarkOnBoard("X", 0)
	aBoard.PutMarkOnBoard("X", 1)
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "3", pickedCell)
}

func TestPickMethodBlocksAWinningOpponent(t *testing.T) {
	marksRepo := board.MarksRepo{"X", "Y"}
	aBoard := board.MakeBoard(3, &marksRepo)
	aBoard.PutMarkOnBoard("X", 0)
	aBoard.PutMarkOnBoard("Y", 1)
	aBoard.PutMarkOnBoard("X", 3)
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "7", pickedCell)
}

func TestPicksTopLeftPositionOn2x2Board(t *testing.T) {
	marksRepo := board.MarksRepo{"X", "Y"}
	aBoard := board.MakeBoard(2, &marksRepo)
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "1", pickedCell)
}