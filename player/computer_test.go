package player

import (
	"testing"
	"tictactoe/board"
	"tictactoe/testhelper"
)

var computerPlayer = Computer{mark: "Y"}

func TestComputerPlayerHasMark(t *testing.T) {
	matchers.EqualLiterals(t, "Y", computerPlayer.mark)
}

func TestPickMethodReturns3(t *testing.T) {
	aBoard := board.MakeBoard(3)
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, 3, pickedCell)
}

func TestGetMarkReturnsComputerMark(t *testing.T) {
	matchers.EqualLiterals(t, "Y", computerPlayer.GetMark())
}