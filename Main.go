package main

import (
	"os"
	"tictactoe/clui"
	"tictactoe/gamescontroller"
)

func main() {
	aClui := clui.MakeClui(os.Stdin, clui.MakeConsoleWriter())
	runStatus := &gamescontroller.RunStatus{}
	gamesController := &gamescontroller.GamesController{Clui: aClui, RunStatus: runStatus}
	gamesController.Run()
}