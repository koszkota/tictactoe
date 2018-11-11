package player_test

import (
	"testing"
	"tictactoe/board"
	"tictactoe/player"
	"tictactoe/testhelper"
)

var humanPlayer = player.Human{Mark: "X"}

func TestHumanPlayerHasMark(t *testing.T) {
	matchers.EqualLiterals(t, humanPlayer.Mark, "X")
}

func TestPickMethodReturns2(t *testing.T) {
	aBoard := board.MakeBoard(3)
	pickedCell := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, pickedCell, 2)
}