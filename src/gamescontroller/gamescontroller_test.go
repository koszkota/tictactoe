package gamescontroller

import (
	"strings"
	"testing"
	"tictactoe/src/clui"
	"tictactoe/src/clui/writer"
	"tictactoe/src/game"
	"tictactoe/src/gamescontroller/controllerstatus"
	"tictactoe/src/player"
	"tictactoe/src/testhelper"
)

func TestRunsAComputerVsComputerGameThatEndsWithTie(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	gamesController := getGamesController(stubWriter, "YES\n3\nNO")

	gamesController.Run()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}

func TestRunsAHumanVsComputerGameWonByComputer(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	gamesController := getGamesController(stubWriter, "YES\n2\nH\nX\n1\n2\n6\nNO")

	gamesController.Run()

	matchers.EqualLiterals(t, "Player O won!", stubWriter.GetLastMessage())
}

func TestRunsAHumanVsHumanGameWonByPlayerX(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	gamesController := getGamesController(stubWriter, "YES\n1\nX\nW\n1\n2\n4\n3\n7\nNO")

	gamesController.Run()

	matchers.EqualLiterals(t, "Player X won!", stubWriter.GetLastMessage())
}

func TestOnInvalidInputInMainMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	gamesController := getGamesController(stubWriter, "invalidGameOption\nYES\n2\nh\nQ\n1\n2\n6\nNO")

	gamesController.Run()

	matchers.EqualLiterals(t, "Option invalidGameOption is not allowed.", stubWriter.GetOutputs()[1])
}

func TestOnInvalidInputInGameTypesMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	gamesController := getGamesController(stubWriter, "YES\nInvalidGameType\n2\nh\nQ\n1\n2\n6\nNO")

	gamesController.Run()

	matchers.EqualLiterals(t, "Option InvalidGameType is not allowed.", stubWriter.GetOutputs()[2])
}

func TestOnInvalidInputInWhoGoesFirstMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &writer.StubWriter{}
	gamesController := getGamesController(stubWriter, "YES\n2\nInvalidOrder\nh\nQ\n1\n2\n6\nNO")

	gamesController.Run()

	matchers.EqualLiterals(t, "Option InvalidOrder is not allowed.", stubWriter.GetOutputs()[3])
}

func getGamesController(stubWriter writer.Writer, input string) GamesController {
	aClui := clui.NewClui(strings.NewReader(input), stubWriter)
	thinkingTimer := &player.ThinkingTimer{ThinkingTime:0}
	playersFactory := &player.Factory{Clui: aClui, ThinkingTimer:thinkingTimer}
	stubRunStatus := &controllerstatus.StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: aClui, PlayerFactory:playersFactory}
	return GamesController{Clui: aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}
}