package gamescontroller

import (
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
		playerOne := player.Computer{Mark: "X"}
		playerTwo := player.Computer{Mark: "Y"}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}

		aBoard := board.MakeBoard(3, &marksRepo)

		aGame := game.MakeGame(gamesController.Clui, &aBoard, playerTwo, playerOne)

		aGame.Play()
	}
}