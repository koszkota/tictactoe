package game

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

func TestReturnsHumanVsHumanGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("1\nX\nO\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func TestReturnsHumanVsHumanGameAfterInvalidModeInput(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("invalid\n1\nX\nO\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func TestReturnsHumanVsComputerGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\nH\nX\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
}

func TestReturnsComputerVsHumanGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\nC\nQ\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "Q")
}

func TestReturnsHumanVsComputerGameAfterInvalidInputInMode(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\ninvalid\nH\nX\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "Y")
}

func TestReturnsHumanVsComputerGameAfterInvalidInputInSign(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("2\nH\ninvalid\nX\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Human")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "Y")
}

func TestReturnsComputerVsComputerGame(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	clui := clui.MakeClui(strings.NewReader("3\n"), stubWriter)
	gameFactory := GameFactory{Clui:clui}

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), "Computer")
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "Y")
}

