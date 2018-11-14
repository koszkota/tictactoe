package gamescontroller

import (
	"os"
	"strings"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/player"
)

type GamesController struct {
	Clui      clui.Clui
	RunStatus ControllerStatus
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
		gamesController.pickGameMode()
	} else if strings.EqualFold("no", mainMenuAnswer)  {
		os.Exit(1)
	} else {
		gamesController.runMainMenu()
	}
}


func (gamesController *GamesController) pickGameMode() {
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Computer{Mark: "Y"}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}

	aBoard := board.MakeBoard(3, &marksRepo)

	aGame := game.MakeGame(gamesController.Clui, &aBoard, playerTwo, playerOne)

	aGame.Play()
}