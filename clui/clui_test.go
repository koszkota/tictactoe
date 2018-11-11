package clui_test

import (
	"tictactoe/clui"
)

func ExampleHelloPlayers() {
	clui.HelloPlayers()
	// Output: Hello and welcome to tic tac toe
}

func ExampleInformOfMove() {
	clui.InformOfMove(1, "X")
	// Output: "Player X picked position 1"
}

func ExampleInformOfWinner() {
	clui.InformOfWinner("X")
	// Output: "Player X won!"
}

func ExampleInformOfTie() {
	clui.InformOfTie()
	// Output: "It's a tie!"
}

func ExampleAskForMove() {
	clui.AskForMove("X")
	// Output: "Player X, pick a position"
}

func ExampleInformOfInvalidMove() {
	clui.InformOfInvalidMove()
	// Output: "This move is not available"
}

