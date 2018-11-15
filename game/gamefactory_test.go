package game

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/player"
	"tictactoe/testhelper"
)

func TestReturnsHumanVsHumanGame(t *testing.T) {
	gameFactory := getGameFactory("1\nX\nO\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func TestReturnsHumanVsHumanGameAfterInvalidModeInput(t *testing.T) {
	gameFactory := getGameFactory("invalid\n1\nX\nO\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func TestReturnsHumanVsComputerGame(t *testing.T) {
	gameFactory := getGameFactory("2\nH\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.ComputerType)
}

func TestInHumanVsHumanGamePlayerOneCannotPickOAsMark(t *testing.T) {
	gameFactory := getGameFactory("1\n0\nH\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "H")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "X")
}

func TestReturnsComputerVsHumanGame(t *testing.T) {
	gameFactory := getGameFactory("2\nC\nQ\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.ComputerType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "Q")
}

func TestReturnsHumanVsComputerGameAfterInvalidInputInMode(t *testing.T) {
	gameFactory := getGameFactory("2\ninvalid\nH\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.ComputerType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func TestReturnsHumanVsComputerGameAfterInvalidInputInSign(t *testing.T) {
	gameFactory := getGameFactory("2\nH\ninvalid\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.HumanType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.ComputerType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func TestReturnsComputerVsComputerGame(t *testing.T) {
	gameFactory := getGameFactory("3\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, game.playerOne.GetType(), player.ComputerType)
	matchers.EqualLiterals(t, game.playerOne.GetMark(), "X")
	matchers.EqualLiterals(t, game.playerTwo.GetType(), player.ComputerType)
	matchers.EqualLiterals(t, game.playerTwo.GetMark(), "O")
}

func getGameFactory(input string) Factory {
	stubWriter := &clui.StubWriter{}
	clui := clui.NewClui(strings.NewReader(input), stubWriter)
	playersFactory := &player.Factory{clui}
	return Factory{Clui: *clui, PlayerFactory:playersFactory}
}