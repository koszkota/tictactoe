package game

import (
	"strings"
	"testing"
	"tictactoe/src/clui"
	"tictactoe/src/clui/writer"
	"tictactoe/src/player"
	"tictactoe/src/testhelper"
)

func TestReturnsHumanVsHumanGame(t *testing.T) {
	gameFactory := getGameFactory("1\nX\nO\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.HumanType, game.playerTwo.GetType())
	matchers.EqualLiterals(t, "O", game.playerTwo.GetMark())
}

func TestReturnsHumanVsHumanGameAfterInvalidModeInput(t *testing.T) {
	gameFactory := getGameFactory("invalid\n1\nX\nO\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.HumanType, game.playerTwo.GetType())
	matchers.EqualLiterals(t, "O", game.playerTwo.GetMark())
}

func TestReturnsHumanVsComputerGame(t *testing.T) {
	gameFactory := getGameFactory("2\nH\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.ComputerType, game.playerTwo.GetType())
}

func TestInHumanVsHumanGamePlayerOneCannotPickOAsMark(t *testing.T) {
	gameFactory := getGameFactory("1\n0\nH\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "H", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerTwo.GetMark())
}

func TestReturnsComputerVsHumanGame(t *testing.T) {
	gameFactory := getGameFactory("2\nC\nQ\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.ComputerType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.HumanType, game.playerTwo.GetType())
	matchers.EqualLiterals(t, "Q", game.playerTwo.GetMark())
}

func TestReturnsHumanVsComputerGameAfterInvalidInputInMode(t *testing.T) {
	gameFactory := getGameFactory("2\ninvalid\nH\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark(), )
	matchers.EqualLiterals(t, player.ComputerType, game.playerTwo.GetType())
	matchers.EqualLiterals(t, "O", game.playerTwo.GetMark())
}

func TestReturnsHumanVsComputerGameAfterInvalidInputInSign(t *testing.T) {
	gameFactory := getGameFactory("2\nH\ninvalid\nX\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t,  player.HumanType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.ComputerType, game.playerTwo.GetType())
	matchers.EqualLiterals(t, "O", game.playerTwo.GetMark())
}

func TestReturnsComputerVsComputerGame(t *testing.T) {
	gameFactory := getGameFactory("3\n")

	game := gameFactory.CreateGame()

	matchers.EqualLiterals(t, player.ComputerType, game.playerOne.GetType())
	matchers.EqualLiterals(t, "X", game.playerOne.GetMark())
	matchers.EqualLiterals(t, player.ComputerType, game.playerTwo.GetType())
	matchers.EqualLiterals(t, "O", game.playerTwo.GetMark())
}

func getGameFactory(input string) Factory {
	stubWriter := &writer.StubWriter{}
	aClui := clui.NewClui(strings.NewReader(input), stubWriter)
	thinkingTimer := &player.ThinkingTimer{ThinkingTime:0}
	playersFactory := &player.Factory{Clui:aClui, ThinkingTimer:thinkingTimer}
	return Factory{Clui: aClui, PlayerFactory:playersFactory}
}