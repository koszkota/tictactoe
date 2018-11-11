package clui

import (
	"testing"
	"tictactoe/testhelper"
)

func ExampleHelloPlayers() {
	HelloPlayers()
	// Output: Hello and welcome to tic tac toe
}

func ExampleInformOfMove() {
	InformOfMove(1, "X")
	// Output: Player X picked position 1
}

func ExampleInformOfWinner() {
	InformOfWinner("X")
	// Output: Player X won!
}

func ExampleInformOfTie() {
	InformOfTie()
	// Output: It's a tie!
}

func ExampleAskForMove() {
	AskForMove("X")
	// Output: Player X, pick a position
}

func ExampleInformOfInvalidMove() {
	InformOfInvalidMove()
	// Output: This move is not available
}

func TestBuildBoardString(t *testing.T) {
	expectedString := "  1  |  2  |  3\n  4  |  X  |  6\n  7  |  8  |  9\n"
	rows := [][]string{{"1","2","3"}, {"4", "X", "6"}, {"7", "8", "9"}}
	actualString := buildBoardString(rows)
	matchers.EqualLiterals(t, actualString, expectedString)
}
