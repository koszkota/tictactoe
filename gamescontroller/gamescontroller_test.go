package gamescontroller

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/player"
	"tictactoe/testhelper"
)

func TestRunsAComputerVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.NewClui(strings.NewReader("YES\n3\nNO"), stubWriter)

	playersFactory := &player.Factory{aClui}
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: *aClui, PlayerFactory:playersFactory}
	gamesController := GamesController{Clui: *aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}

func TestRunsAHumanVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.NewClui(strings.NewReader("YES\n2\nH\nX\n1\n2\n6\nNO"), stubWriter)
	playersFactory := &player.Factory{aClui}
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: *aClui, PlayerFactory:playersFactory}
	gamesController := GamesController{Clui: *aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Player O won!", stubWriter.GetLastMessage())
}

func TestRunsAHumanVsHumanGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.NewClui(strings.NewReader("YES\n1\nX\nW\n1\n2\n4\n3\n7\nNO"), stubWriter)
	playersFactory := &player.Factory{aClui}
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: *aClui, PlayerFactory:playersFactory}
	gamesController := GamesController{Clui: *aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Player X won!", stubWriter.GetLastMessage())
}

func TestOnInvalidInputInMainMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.NewClui(strings.NewReader("invalidGameOption\nYES\n2\nh\nQ\n1\n2\n6\nNO"), stubWriter)
	playersFactory := &player.Factory{aClui}
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: *aClui, PlayerFactory:playersFactory}
	gamesController := GamesController{Clui: *aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Option invalidGameOption is not allowed.", stubWriter.GetOutputs()[1])
}

func TestOnInvalidInputInGameTypesMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.NewClui(strings.NewReader("YES\nInvalidGameType\n2\nh\nQ\n1\n2\n6\nNO"), stubWriter)
	playersFactory := &player.Factory{aClui}
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: *aClui, PlayerFactory:playersFactory}
	gamesController := GamesController{Clui: *aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Option InvalidGameType is not allowed.", stubWriter.GetOutputs()[2])
}

func TestOnInvalidInputInWhoGoesFirstMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.NewClui(strings.NewReader("YES\n2\nInvalidOrder\nh\nQ\n1\n2\n6\nNO"), stubWriter)
	playersFactory := &player.Factory{aClui}
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.Factory{Clui: *aClui, PlayerFactory:playersFactory}
	gamesController := GamesController{Clui: *aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Option InvalidOrder is not allowed.", stubWriter.GetOutputs()[3])
}