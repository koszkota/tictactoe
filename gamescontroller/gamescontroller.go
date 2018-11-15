package gamescontroller

import (
	"os"
	"strings"
	. "tictactoe/clui"
	. "tictactoe/game"
)

type GamesController struct {
	Clui      *Clui
	RunStatus ControllerStatus
	GameFactory *Factory
}

func (gamesController *GamesController) Run() {
	for gamesController.RunStatus.GetRunStatus() {
		gamesController.runMainMenu()
	}
}

func (gamesController *GamesController) runMainMenu() {
	gamesController.Clui.ShowMainMenu()
	userInput := gamesController.Clui.ReadUserInput()
	if strings.EqualFold("yes", userInput) {
		gamesController.playGame()
	} else if strings.EqualFold("no", userInput)  {
		os.Exit(0)
	} else {
		gamesController.Clui.InformOfInvalidInput(userInput)
		gamesController.runMainMenu()
	}
}

func (gamesController *GamesController) playGame() {
	game := gamesController.GameFactory.CreateGame()
	game.Play()
}