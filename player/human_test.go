package player

import (
	"strings"
	"testing"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

func TestHumanPlayerHasMark(t *testing.T) {
	var stubWriter = new(clui.StubWriter)
	clui := clui.MakeClui(strings.NewReader("1"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: clui}

	matchers.EqualLiterals(t, "X", humanPlayer.Mark)
}

func TestPickMoveMethodReturnsPlayersPick(t *testing.T) {
	var stubWriter = new(clui.StubWriter)
	clui := clui.MakeClui(strings.NewReader("1"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: clui}
	aBoard := board.MakeBoard(3)

	pickedCell := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "1", pickedCell)
}

func TestPickMoveMethodPromptsPlayerFormMove(t *testing.T) {
	var stubWriter = new(clui.StubWriter)
	clui := clui.MakeClui(strings.NewReader("1"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: clui}
	aBoard := board.MakeBoard(3)

	humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "Player X, pick a position", stubWriter.GetLastMessage())
}

func TestGetMarkReturnsHumanMark(t *testing.T) {
	var stubWriter = new(clui.StubWriter)
	clui := clui.MakeClui(strings.NewReader("1"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: clui}

	matchers.EqualLiterals(t, "X", humanPlayer.GetMark())
}

func TestHumanIsPrompterForValidMoveAndSubmitsIt(t *testing.T) {
	var stubWriter = new(clui.StubWriter)
	clui := clui.MakeClui(strings.NewReader("6\n2\n"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: clui}
	aBoard := board.MakeBoard(3)
	aBoard.PutMarkOnBoard("X", 5)

	move := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "2", move)
	matchers.EqualLiterals(t, "This move is not available", stubWriter.GetOutputs()[1])
}