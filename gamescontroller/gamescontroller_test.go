package gamescontroller

import (
	"strings"
	"testing"
	"tictactoe/clui"
)

func TestRunningEndsWhenRunStatusIsFalse(t *testing.T) {
	stubWriter := &clui.StubWriter{}
	aClui := clui.MakeClui(strings.NewReader("stubinput"), stubWriter)
	stubRunStatus := &StubRunStatus{Counter: 0}
	gamesController := GamesController{Clui:aClui, RunStatus: stubRunStatus}
	gamesController.Run()
}