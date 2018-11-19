package player

import (
	"strings"
	"testing"
	"../board"
	"../clui"
	"../clui/writer"
	"../testhelper"
)

func TestHumanPlayerHasMark(t *testing.T) {
	aClui := getCluiWithHardcodedInput("1")

	humanPlayer := Human{Mark: "X", Clui: aClui}

	matchers.EqualLiterals(t, "X", humanPlayer.Mark)
}

func TestHumanPlayerHasType(t *testing.T) {
	aClui := getCluiWithHardcodedInput("1")

	humanPlayer := Human{Mark: "X", Clui: aClui}

	matchers.EqualLiterals(t, HumanType, humanPlayer.GetType())
}

func TestPickMoveMethodReturnsPlayersPick(t *testing.T) {
	aClui := getCluiWithHardcodedInput("1")
	humanPlayer := Human{Mark: "X", Clui: aClui}
	marksRepo := board.MarksRepo{"X", "Y"}
	aBoard := board.MakeBoard(3, &marksRepo)

	pickedCell := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "1", pickedCell)
}

func TestPickMoveMethodPromptsPlayerFormMove(t *testing.T) {
	var stubWriter = new(writer.StubWriter)
	aClui := clui.NewClui(strings.NewReader("1"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: aClui}
	marksRepo := board.MarksRepo{"X", "Y"}
	aBoard := board.MakeBoard(3, &marksRepo)

	humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "Player X, pick a position", stubWriter.GetLastMessage())
}

func TestGetMarkReturnsHumanMark(t *testing.T) {
	var stubWriter = new(writer.StubWriter)
	aClui := clui.NewClui(strings.NewReader("1"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: aClui}

	matchers.EqualLiterals(t, "X", humanPlayer.GetMark())
}

func TestHumanIsPrompterForValidMoveAndSubmitsIt(t *testing.T) {
	var stubWriter = new(writer.StubWriter)
	aClui := clui.NewClui(strings.NewReader("6\n2\n"), stubWriter)
	humanPlayer := Human{Mark: "X", Clui: aClui}
	marksRepo := board.MarksRepo{"X", "Y"}
	aBoard := board.MakeBoard(3, &marksRepo)
	aBoard.PutMarkOnBoard("X", 5)

	move := humanPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "This move is not available", stubWriter.GetOutputs()[1])
	matchers.EqualLiterals(t, "2", move)
}

func getCluiWithHardcodedInput(input string) *clui.Clui {
	var stubWriter = new(writer.StubWriter)
	return clui.NewClui(strings.NewReader(input), stubWriter)
}