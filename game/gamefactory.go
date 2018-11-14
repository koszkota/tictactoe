package game

import (
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
)

type GameFactory struct {
	Clui clui.Clui
}

func (gameFactory *GameFactory) CreateGame() *Game {
	gameFactory.Clui.AskForGameMode()
	pick := gameFactory.Clui.ReadUserInput()
	if pick == "1" {
		playerOne := player.Human{Mark: "X", Clui:gameFactory.Clui}
		playerTwo := player.Human{Mark: "Y", Clui:gameFactory.Clui}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
		return &aGame
	} else if pick == "2" {
		return gameFactory.whoGoesFirst()
	} else if pick == "3" {
		playerOne := player.Computer{Mark: "X"}
		playerTwo := player.Computer{Mark: "Y"}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
		return &aGame
	} else {
		return gameFactory.CreateGame()
	}
}

func (gameFactory *GameFactory) whoGoesFirst() *Game {
	gameFactory.Clui.AskWhoGoesFirst()
	pick := gameFactory.Clui.ReadUserInput()
	if pick == "H" {
		playerOne := player.Human{Mark: "X", Clui:gameFactory.Clui}
		playerTwo := player.Computer{Mark: "Y"}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
		return &aGame
	} else if pick == "C" {
		playerOne := player.Computer{Mark: "X"}
		playerTwo := player.Human{Mark: "Y", Clui:gameFactory.Clui}
		marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
		aBoard := board.MakeBoard(3, &marksRepo)
		aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
		return &aGame
	} else {
		return gameFactory.whoGoesFirst()
	}
}