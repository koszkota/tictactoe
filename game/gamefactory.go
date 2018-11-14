package game

import (
	"strings"
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
		return gameFactory.getHumanVsHumanGame()
	} else if pick == "2" {
		return gameFactory.getMixedPlayersGame()
	} else if pick == "3" {
		return gameFactory.getComputerVsComputerGame()
	} else {
		return gameFactory.CreateGame()
	}
}

func (gameFactory *GameFactory) getMixedPlayersGame() *Game {
	gameFactory.Clui.AskWhoGoesFirst()
	pick := gameFactory.Clui.ReadUserInput()
	if strings.EqualFold(pick, "h") {
		return gameFactory.getHumanVsComputerGame()
	} else if strings.EqualFold(pick, "c") {
	return gameFactory.getComputerVsHumanGame()
	} else {
		return gameFactory.getMixedPlayersGame()
	}
}

func (gameFactory *GameFactory) getHumanVsHumanGame() *Game {
	playerOne := player.Human{Mark: "X", Clui:gameFactory.Clui}
	playerTwo := player.Human{Mark: "Y", Clui:gameFactory.Clui}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getComputerVsComputerGame() *Game {
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Computer{Mark: "Y"}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getHumanVsComputerGame() *Game {
	playerOne := player.Human{Mark: "X", Clui: gameFactory.Clui}
	playerTwo := player.Computer{Mark: "Y"}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getComputerVsHumanGame() *Game {
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Human{Mark: "Y", Clui:gameFactory.Clui}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}