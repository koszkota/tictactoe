package main

import (
	"os"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/gamescontroller"
)

func main() {
	aClui := clui.MakeClui(os.Stdin, clui.MakeConsoleWriter())
	runStatus := &gamescontroller.RunStatus{}
	gameFactory := &game.GameFactory{}
	gamesController := &gamescontroller.GamesController{Clui: aClui, RunStatus: runStatus, GameFactory: gameFactory}
	gamesController.Run()
}