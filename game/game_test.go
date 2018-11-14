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
	stubWriter := &clui.StubWriter{}
	aClui := makeCluiWithInput("5\n1\n3\n7\n4\n6\n2\n8\n9\n", stubWriter)
	playerOne := player.Human{Mark: "X", Clui: aClui}
	playerTwo := player.Human{Mark: "Y", Clui: aClui}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := MakeGame(aClui, &aBoard, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "It's a tie!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func TestGamePlaysWholeHumanVsHumanWonGameForPlayerOne(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := makeCluiWithInput("1\n2\n4\n5\n7\n", stubWriter)
	playerOne := player.Human{Mark: "X", Clui: aClui}
	playerTwo := player.Human{Mark: "Y", Clui: aClui}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := MakeGame(aClui, &aBoard, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "Player X won!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func TestGamePlaysWholeHumanVsHumanWonGameForPlayerTwo(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := makeCluiWithInput("1\n2\n3\n5\n6\n8\n", stubWriter)
	playerOne := player.Human{Mark: "X", Clui: aClui}
	playerTwo := player.Human{Mark: "Y", Clui: aClui}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := MakeGame(aClui, &aBoard, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "Player Y won!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func TestGamePlaysWholeComputerVsComputerGameWithTie(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader(""), stubWriter)
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Computer{Mark: "Y"}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := MakeGame(aClui, &aBoard, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "It's a tie!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}

func makeCluiWithInput(hardcodedMovesOfPlayers string, stubWriter *clui.StubWriter) clui.Clui {
	return clui.MakeClui(strings.NewReader(hardcodedMovesOfPlayers), stubWriter)
}

func makeBoard(size int, playerOneMark string, playerTwoMark string) board.Board {
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	return board.MakeBoard(size, &marksRepo)
}