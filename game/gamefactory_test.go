package game

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

func TestReturnsHumanVsHumanGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("1\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Human")
}

func TestReturnsHumanVsHumanGameAfterInvalidInput(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("invalid\n1\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Human")
}

func TestReturnsHumanVsComputerGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\nH\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
}

func TestReturnsComputerVsHumanGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\nC\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Human")
}

func TestReturnsHumanVsComputerGameAfterInvalidInput(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\ninvalid\nH\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
}

func TestReturnsComputerVsComputerGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("3\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
}

