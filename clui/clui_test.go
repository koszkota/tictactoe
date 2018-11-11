package clui_test

import (
	"testing"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

func TestHelloPlayers(t *testing.T) {
	expectedString := "Hello and welcome to tic tac toe"
	actualString := clui.HelloPlayers()

	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestInformOfMove(t *testing.T) {
	expectedString := "Player X picked position 1"
	actualString := clui.InformOfMove(1, "X")

	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestInformOfWinner(t *testing.T) {
	expectedString := "Player X won!"
	actualString := clui.InformOfWinner("X")

	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestInformOfTie(t *testing.T) {
	expectedString := "It's a tie!"
	actualString := clui.InformOfTie()

	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestAskForMove(t *testing.T) {
	expectedString := "Player X, pick a position"
	actualString := clui.AskForMove("X")

	matchers.EqualLiterals(t, actualString, expectedString)
}

func TestInformOfInvalidMove(t *testing.T) {
	expectedString := "This move is not available"
	actualString := clui.InformOfInvalidMove()

	matchers.EqualLiterals(t, actualString, expectedString)
}

