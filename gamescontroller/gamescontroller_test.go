package gamescontroller

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/testhelper"
)

func TestRunsAComputerVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("YES\n3\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.GameFactory{Clui:aClui}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}

func TestOnInvalidInputInMainMenuNeedsToSubmitAgain(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("invalid\nYES\n3\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.GameFactory{Clui:aClui}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}

func TestRunsAHumanVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("YES\n2\nH\nX\n1\n2\n6\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.GameFactory{Clui:aClui}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Player Y won!", stubWriter.GetLastMessage())
}

func TestRunsAHumanVsHumanGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("YES\n1\nX\nW\n1\n2\n4\n3\n7\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.GameFactory{Clui:aClui}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Player X won!", stubWriter.GetLastMessage())
}