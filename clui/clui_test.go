package clui

import (
	"strings"
	"testing"
	"tictactoe/testhelper"
)

func ExampleHelloPlayers() {
	HelloPlayers()
	// Output: Hello and welcome to tic tac toe
}

func ExampleInformOfMove() {
	informOfMove(1, "X")
	// Output: Player X picked position 1
}

func ExampleInformOfWinner() {
	informOfWinner("X")
	// Output: Player X won!
}

func ExampleInformOfTie() {
	informOfTie()
	// Output: It's a tie!
}

func ExampleAskForMove() {
	askForMove("X")
	// Output: Player X, pick a position
}

func ExampleInformOfInvalidMove() {
	informOfInvalidMove()
	// Output: This move is not available
}

func TestBuildBoardString(t *testing.T) {
	expectedString := "  1  |  2  |  3\n  4  |  X  |  6\n  7  |  8  |  9\n"
	rows := [][]string{{"1","2","3"}, {"4", "X", "6"}, {"7", "8", "9"}}
	actualString := buildBoardString(rows)
	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestReadUserInputAndReturnIt(t *testing.T) {
	hardcodedInput := "testWord"
	// in real MakeCluiReader pass os.Stdin
	aClui := MakeCluiReader(strings.NewReader(hardcodedInput))
	actualString := aClui.ReadUserInput()

	matchers.EqualLiterals(t, hardcodedInput, actualString)
}
