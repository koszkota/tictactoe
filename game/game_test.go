package game

import (
	"strings"
	"testing"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
	"tictactoe/testhelper"
)

func TestGamePlaysWholeHumanVsHumanTieGame(t *testing.T) {
	movesOfPlayers := "5\n1\n3\n7\n4\n6\n2\n8\n9\n"
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader(movesOfPlayers), stubWriter)
	marksRepo := board.MarksRepo{"X", "Y"}
	board := board.MakeBoard(3, &marksRepo)
	playerOne := player.Human{Mark: "X", Clui: clui}
	playerTwo := player.Human{Mark: "Y", Clui: clui}
	aGame := MakeGame(clui, &board, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "It's a tie!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func TestGamePlaysWholeHumanVsHumanWonGameForPlayerOne(t *testing.T) {
	movesOfPlayers := "1\n2\n4\n5\n7\n"
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader(movesOfPlayers), stubWriter)
	marksRepo := board.MarksRepo{"X", "Y"}
	board := board.MakeBoard(3, &marksRepo)
	playerOne := player.Human{Mark: "X", Clui: clui}
	playerTwo := player.Human{Mark: "Y", Clui: clui}
	aGame := MakeGame(clui, &board, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "Player X won!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func TestGamePlaysWholeHumanVsHumanWonGameForPlayerTwo(t *testing.T) {
	movesOfPlayers := "1\n2\n3\n5\n6\n8\n"
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader(movesOfPlayers), stubWriter)
	marksRepo := board.MarksRepo{"X", "Y"}
	board := board.MakeBoard(3, &marksRepo)
	playerOne := player.Human{Mark: "X", Clui: clui}
	playerTwo := player.Human{Mark: "Y", Clui: clui}
	aGame := MakeGame(clui, &board, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "Player Y won!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func TestGamePlaysWholeComputerVsVomputerGameWithTie(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader(""), stubWriter)
	marksRepo := board.MarksRepo{"X", "Y"}
	board := board.MakeBoard(3, &marksRepo)
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Computer{Mark: "Y"}
	aGame := MakeGame(clui, &board, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "It's a tie!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}