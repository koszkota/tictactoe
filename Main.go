package main

import (
	"os"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/gamescontroller"
	"tictactoe/player"
)

func main() {
	aClui := clui.MakeClui(os.Stdin, clui.MakeConsoleWriter())
	runStatus := &gamescontroller.RunStatus{}
	playersFactory := &player.Factory{aClui}
	gameFactory := &game.Factory{Clui: aClui, PlayerFactory: playersFactory}
	gamesController := &gamescontroller.GamesController{Clui: aClui, RunStatus: runStatus, GameFactory: gameFactory}
	gamesController.Run()
}