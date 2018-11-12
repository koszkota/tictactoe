package game

import (
	"strings"
	"testing"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
	"tictactoe/testhelper"
)

func TestGamePlaysTheWholeTieGame(t *testing.T) {
	movesOfPlayers := "5\n1\n3\n7\n4\n6\n2\n8\n9\n"
	stubWriter := new(clui.StubWriter)
	clui := clui.MakeClui(strings.NewReader(movesOfPlayers), stubWriter)
	board := board.MakeBoard(3)
	playerOne := player.Human{"X"}
	playerTwo := player.Human{"Y"}
	aGame := MakeGame(clui, &board, playerOne, playerTwo)

	aGame.Play()

	expectedFinalMessage :=  "It's a tie!"
	matchers.EqualLiterals(t, expectedFinalMessage, stubWriter.GetLastMessage())
	stubWriter.CleanOutputs()
}