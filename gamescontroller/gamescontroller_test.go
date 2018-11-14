package gamescontroller

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/testhelper"
)

func TestRusAComputerVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("YES\n3\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.GameFactory{}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}

func TestRusAHumanVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("YES\n2\n1\n2\n6\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gameFactory := &game.GameFactory{}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus, GameFactory:gameFactory}

	gamesController.Run()

	matchers.EqualLiterals(t, "Player Y won!", stubWriter.GetLastMessage())
}