package gamescontroller

import (
	"os"
	"strings"
	"tictactoe/clui"
	"tictactoe/game"
)

type GamesController struct {
	Clui      clui.Clui
	RunStatus ControllerStatus
	GameFactory *game.GameFactory
}

func (gamesController *GamesController) Run() {
	for gamesController.RunStatus.GetRunStatus() {
		gamesController.runMainMenu()
	}
}

func (gamesController *GamesController) runMainMenu() {
	gamesController.Clui.ShowMainMenu()
	mainMenuAnswer := gamesController.Clui.ReadUserInput()
	if strings.EqualFold("yes", mainMenuAnswer) {
		gamesController.playGame()
	} else if strings.EqualFold("no", mainMenuAnswer)  {
		os.Exit(1)
	} else {
		gamesController.runMainMenu()
	}
}

func (gamesController *GamesController) playGame() {
	game := gamesController.GameFactory.CreateGame(gamesController.Clui)
	game.Play()
}