package game

import (
	"strings"
	"testing"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/clui/writer"
	"tictactoe/player"
	"tictactoe/testhelper"
)

func TestGamePlaysWholeHumanVsHumanTieGame(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	aClui := makeCluiWithInput("5\n1\n3\n7\n4\n6\n2\n8\n9\n", stubWriter)
	playerOne := player.Human{Mark: "X", Clui: aClui}
	playerTwo := player.Human{Mark: "O", Clui: aClui}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := Game{aClui, &aBoard, playerOne, playerTwo}
	expectedFinalMessage :=  "It's a tie!"

	aGame.Play()

	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
}

func TestGamePlaysWholeHumanVsHumanWonGameForPlayerOne(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	aClui := makeCluiWithInput("1\n2\n4\n5\n7\n", stubWriter)
	playerOne := player.Human{Mark: "X", Clui: aClui}
	playerTwo := player.Human{Mark: "O", Clui: aClui}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := Game{aClui, &aBoard, playerOne, playerTwo}
	expectedFinalMessage :=  "Player X won!"

	aGame.Play()

	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
}

func TestGamePlaysWholeHumanVsHumanWonGameForPlayerTwo(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	aClui := makeCluiWithInput("1\n2\n3\n5\n6\n8\n", stubWriter)
	playerOne := player.Human{Mark: "X", Clui: aClui}
	playerTwo := player.Human{Mark: "O", Clui: aClui}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := Game{aClui, &aBoard, playerOne, playerTwo}
	expectedFinalMessage :=  "Player O won!"

	aGame.Play()

	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
}

func TestGamePlaysWholeComputerVsComputerGameWithTie(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	aClui := clui.NewClui(strings.NewReader(""), stubWriter)
	thinkingTimer := &player.ThinkingTimer{0}
	playerOne := player.Computer{Mark: "X", Clui: aClui, ThinkingTimer: thinkingTimer}
	playerTwo := player.Computer{Mark: "O", Clui: aClui, ThinkingTimer: thinkingTimer}
	aBoard := makeBoard(3, playerOne.GetMark(), playerTwo.GetMark())
	aGame := Game{aClui, &aBoard, playerOne, playerTwo}
	expectedFinalMessage :=  "It's a tie!"

	aGame.Play()

	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
}

func makeCluiWithInput(hardcodedMovesOfPlayers string, stubWriter *writer.StubWriter) *clui.Clui {
	return clui.NewClui(strings.NewReader(hardcodedMovesOfPlayers), stubWriter)
}

func makeBoard(size int, playerOneMark string, playerTwoMark string) board.Board {
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	return board.MakeBoard(size, &marksRepo)
}