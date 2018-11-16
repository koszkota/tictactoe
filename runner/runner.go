package runner

import (
	"os"
	. "tictactoe/clui"
	"tictactoe/game"
	. "tictactoe/gamescontroller"
	"tictactoe/player"
)

func Start() {
	aClui := NewClui(os.Stdin, MakeConsoleWriter())
	runStatus := &RunStatus{}
	thinkingTimer := &player.ThinkingTimer{3}
	playersFactory := &player.Factory{Clui: aClui, ThinkingTimer: thinkingTimer}
	gameFactory := &game.Factory{Clui: aClui, PlayerFactory: playersFactory}
	gamesController := &GamesController{Clui: aClui, RunStatus: runStatus, GameFactory: gameFactory}
	gamesController.Run()
}
