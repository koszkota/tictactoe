package clui

import (
	"os"
	"strings"
	"testing"
	. "tictactoe/clui/writer"
	"tictactoe/testhelper"
)

func TestInformOfBeginningOfGame(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfBeginningOfGame()

	matchers.EqualLiterals(t, "Let's start the game!", stubWriter.GetLastMessage())
}

func TestInformOfMove(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfMove("1", "X")

	matchers.EqualLiterals(t, "Player X picked position 1", stubWriter.GetLastMessage())
}

func TestInformOfWinner(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfWinner("X")

	matchers.EqualLiterals(t, "Player X won!", stubWriter.GetLastMessage())
}

func TestShowMainMenu(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.ShowMainMenu()

	matchers.EqualLiterals(t, "To play a game enter YES, to exit enter NO.", stubWriter.GetLastMessage())
}

func TestShowGameMode(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.AskForGameMode()

	matchers.EqualLiterals(t, "To play Human vs Human enter 1. To play Human vs Computer enter 2. To see two computers playing enter 3.", stubWriter.GetLastMessage())
}

func TestShowWhoGoesFirstMenu(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.AskWhoGoesFirst()

	matchers.EqualLiterals(t, "If Human player should go first, enter H; if computer, enter C.", stubWriter.GetLastMessage())
}

func TestAskPlayerOneForMark(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.AskPlayerForMark("one")

	matchers.EqualLiterals(t, "Player one, please pick one-letter mark", stubWriter.GetLastMessage())
}

func TestInformOfNotAvailableMark(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfNotAvailableMark("X")

	matchers.EqualLiterals(t, "Mark X is not available, please pick another one.", stubWriter.GetLastMessage())
}

func TestInformOfTie(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfTie()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}

func TestAskForMove(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.AskForMove("X")

	matchers.EqualLiterals(t, "Player X, pick a position", stubWriter.GetLastMessage())
}

func TestInformOfInvalidMove(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfInvalidMove()

	matchers.EqualLiterals(t, "This move is not available", stubWriter.GetLastMessage())
}

func TestInformOfInvalidInput(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)

	clui.InformOfInvalidInput("test")

	matchers.EqualLiterals(t, "Option test is not allowed.", stubWriter.GetLastMessage())
}

func TestInformOfThinkingComputer(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)
	expected := "\r\x1b[36mComputer is thinking...\x1b[m %s "

	clui.InformOfThinkingComputer(0)

	matchers.EqualLiterals(t, expected, stubWriter.GetOutputs()[len(stubWriter.GetOutputs())-2])
}

func TestShowBoard(t *testing.T) {
	stubWriter := &StubWriter{}
	clui := NewClui(os.Stdin, stubWriter)
	rows := [][]string{{"1","2","3"}, {"4", "X", "6"}, {"7", "8", "9"}}

	clui.ShowBoard(rows)

	matchers.EqualLiterals(t, "  1  |  2  |  3\n  4  |  X  |  6\n  7  |  8  |  9\n", stubWriter.GetOutputs()[0])
}

func TestBuildBoardString(t *testing.T) {
	expectedString := "  1  |  2  |  3\n  4  |  X  |  6\n  7  |  8  |  9\n"
	rows := [][]string{{"1","2","3"}, {"4", "X", "6"}, {"7", "8", "9"}}
	actualString := buildBoardString(rows)
	matchers.EqualLiterals(t, expectedString, actualString)
}

func TestReadUserInputAndReturnIt(t *testing.T) {
	writer := ConsoleWriter{}
	hardcodedInput := "testWord"
	clui := NewClui(strings.NewReader(hardcodedInput), writer)

	actualString := clui.ReadUserInput()

	matchers.EqualLiterals(t, hardcodedInput, actualString)
}