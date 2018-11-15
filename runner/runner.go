package runner

import (
	"os"
	. "tictactoe/clui"
	"tictactoe/game"
	. "tictactoe/gamescontroller"
	"tictactoe/player"
)

//type Runner struct {}

func Start() {
	aClui := NewClui(os.Stdin, MakeConsoleWriter())
	runStatus := &RunStatus{}
	playersFactory := &player.Factory{Clui: aClui}
	gameFactory := &game.Factory{Clui: aClui, PlayerFactory: playersFactory}
	gamesController := &GamesController{Clui: aClui, RunStatus: runStatus, GameFactory: gameFactory}
	gamesController.Run()
}
