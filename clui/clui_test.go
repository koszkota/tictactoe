package clui

import (
	"strings"
	"testing"
	"tictactoe/testhelper"
)

var clui = setupClui()
var hardcodedInput = "testWord"

func setupClui() Clui {
	writer := MakeConsoleWriter()
	clui := MakeClui(strings.NewReader(hardcodedInput), writer)
	return clui
}

func ExampleHelloPlayers() {
	clui.HelloPlayers()
	// Output: Hello and welcome to tic tac toe
}

func ExampleInformOfMove() {
	clui.InformOfMove("1", "X")
	// Output: Player X picked position 1
}

func ExampleInformOfWinner() {
	clui.InformOfWinner("X")
	// Output: Player X won!
}

func ExampleShowMainMenu() {
	clui.ShowMainMenu()
	// Output: To play a game enter YES, to exit enter NO.
}

func ExampleInformOfTie() {
	clui.InformOfTie()
	// Output: It's a tie!
}

func ExampleAskForMove() {
	clui.AskForMove("X")
	// Output: Player X, pick a position
}

func ExampleInformOfInvalidMove() {
	clui.InformOfInvalidMove()
	// Output: This move is not available
}

func TestBuildBoardString(t *testing.T) {
	expectedString := "  1  |  2  |  3\n  4  |  X  |  6\n  7  |  8  |  9\n"
	rows := [][]string{{"1","2","3"}, {"4", "X", "6"}, {"7", "8", "9"}}
	actualString := buildBoardString(rows)
	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestReadUserInputAndReturnIt(t *testing.T) {
	actualString := clui.ReadUserInput()

	matchers.EqualLiterals(t, hardcodedInput, actualString)
}
