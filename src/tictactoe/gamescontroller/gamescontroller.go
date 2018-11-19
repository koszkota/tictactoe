package gamescontroller

import (
	"os"
	"strings"
	. "../clui"
	. "../game"
	. "../gamescontroller/controllerstatus"
)

type GamesController struct {
	Clui      *Clui
	RunStatus ControllerStatus
	GameFactory *Factory
}

const(
	playGameOption = "yes"
	exitGameOption = "no"
)

func (gamesController *GamesController) Run() {
	for gamesController.RunStatus.GetRunStatus() {
		gamesController.runMainMenu()
	}
}

func (gamesController *GamesController) runMainMenu() {
	gamesController.Clui.ShowMainMenu()
	userInput := gamesController.Clui.ReadUserInput()
	if strings.EqualFold(playGameOption, userInput) {
		gamesController.playGame()
	} else if strings.EqualFold(exitGameOption, userInput)  {
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