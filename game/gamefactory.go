package game

import (
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
)

type GameFactory struct {}

func (gameFactory *GameFactory) CreateGame(clui clui.Clui) *Game {
	clui.AskForGameMode()
	pick := clui.ReadUserInput()
	if pick == "1" {
		playerOne := player.Human{Mark: "X", Clui:clui}
		playerTwo := player.Human{Mark: "Y", Clui: clui}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(clui, &aBoard, playerTwo, playerOne)
		return &aGame
	} else if pick == "2" {
		playerOne := player.Human{Mark: "X", Clui:clui}
		playerTwo := player.Computer{Mark: "Y"}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(clui, &aBoard, playerTwo, playerOne)
		return &aGame
	} else if pick == "3" {
		playerOne := player.Computer{Mark: "X"}
		playerTwo := player.Computer{Mark: "Y"}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(clui, &aBoard, playerTwo, playerOne)
		return &aGame
	} else {
		return gameFactory.CreateGame(clui)
	}
}