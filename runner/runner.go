package runner

import (
	"os"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/gamescontroller"
	"tictactoe/player"
)

type Runner struct {}

func (runner *Runner) Start() {
	aClui := clui.NewClui(os.Stdin, clui.MakeConsoleWriter())
	runStatus := &gamescontroller.RunStatus{}
	playersFactory := &player.Factory{Clui: aClui}
	gameFactory := &game.Factory{Clui: aClui, PlayerFactory: playersFactory}
	gamesController := &gamescontroller.GamesController{Clui: aClui, RunStatus: runStatus, GameFactory: gameFactory}
	gamesController.Run()
}
