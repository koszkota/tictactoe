package player

import (
	"testing"
	"tictactoe/board"
	"tictactoe/testhelper"
)

var humanPlayer = Human{Mark: "X"}

func TestHumanPlayerHasMark(t *testing.T) {
	matchers.EqualLiterals(t, "X", humanPlayer.Mark)
}

func TestPickMethodReturns2(t *testing.T) {
	aBoard := board.MakeBoard(3)
	pickedCell := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, 2, pickedCell)
}

func TestGetMarkReturnsHumanMark(t *testing.T) {
	matchers.EqualLiterals(t, "X", humanPlayer.GetMark())
}