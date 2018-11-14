package gamescontroller

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

func TestRusAComputerVsComputerGameAndEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("YES\nNO"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus}

	gamesController.Run()

	matchers.EqualLiterals(t, "It's a tie!", stubWriter.GetLastMessage())
}