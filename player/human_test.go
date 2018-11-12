package player

import (
	"strings"
	"testing"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

var stubWriter = new(clui.StubWriter)
var humanPlayer = setupPlayer(stubWriter)


func setupPlayer(writer *clui.StubWriter) Human {
	var hardcodedInput = "1"
	clui := clui.MakeClui(strings.NewReader(hardcodedInput), writer)
	return Human{Mark: "X", Clui: clui}
}


func TestHumanPlayerHasMark(t *testing.T) {
	matchers.EqualLiterals(t, "X", humanPlayer.Mark)
}

func TestPickMoveMethodReturnsPlayersPick(t *testing.T) {
	aBoard := board.MakeBoard(3)
	pickedCell := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "1", pickedCell)
}

func TestPickMoveMethodPromptsPlayerFormMove(t *testing.T) {
	aBoard := board.MakeBoard(3)
	humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "Player X, pick a position", stubWriter.GetLastMessage())

	stubWriter.CleanOutputs()
}

func TestGetMarkReturnsHumanMark(t *testing.T) {
	matchers.EqualLiterals(t, "X", humanPlayer.GetMark())
}