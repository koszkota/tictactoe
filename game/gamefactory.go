package game

import (
	"regexp"
	"strings"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
)

type GameFactory struct {
	Clui clui.Clui
}

const(
	humanHumanGame       = "1"
	humanComputerGame    = "2"
	computerComputerGame = "3"
	humanGoesFirst       = "h"
	computeGoesFirst     = "c"
)

func (gameFactory *GameFactory) CreateGame() *Game {
	gameFactory.Clui.AskForGameMode()
	userInput := gameFactory.Clui.ReadUserInput()
	switch userInput {
	case humanHumanGame:
		return gameFactory.getHumanVsHumanGame()
	case humanComputerGame:
		return gameFactory.getMixedPlayersGame()
	case computerComputerGame:
		return gameFactory.getComputerVsComputerGame()
	default:
		gameFactory.Clui.InformOfInvalidInput(userInput)
		return gameFactory.CreateGame()
	}
}

func (gameFactory *GameFactory) getMixedPlayersGame() *Game {
	gameFactory.Clui.AskWhoGoesFirst()
	userInput := gameFactory.Clui.ReadUserInput()
	if strings.EqualFold(userInput, humanGoesFirst) {
		return gameFactory.getHumanVsComputerGame()
	} else if strings.EqualFold(userInput, computeGoesFirst) {
	return gameFactory.getComputerVsHumanGame()
	} else {
		gameFactory.Clui.InformOfInvalidInput(userInput)
		return gameFactory.getMixedPlayersGame()
	}
}

func (gameFactory *GameFactory) getHumanVsHumanGame() *Game {
	playerOneMark := gameFactory.getPlayerOneMark()
	playerTwoMark := gameFactory.getPlayerTwoMark(playerOneMark)
	playerOne := player.Human{Mark: playerOneMark, Clui:gameFactory.Clui}
	playerTwo := player.Human{Mark: playerTwoMark, Clui:gameFactory.Clui}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getComputerVsComputerGame() *Game {
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Computer{Mark: "O"}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getHumanVsComputerGame() *Game {
	playerOneMark := gameFactory.getPlayerOneMark()
	playerOne := player.Human{Mark: playerOneMark, Clui: gameFactory.Clui}
	playerTwo := player.Computer{Mark: "O"}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getComputerVsHumanGame() *Game {
	playerOne := player.Computer{Mark: "X"}
	playerTwoMark := gameFactory.getPlayerTwoMark(playerOne.GetMark())
	playerTwo := player.Human{Mark: playerTwoMark, Clui: gameFactory.Clui}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwoMark}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *GameFactory) getPlayerOneMark() string {
	gameFactory.Clui.AskPlayerOneForMark()
	userInput := gameFactory.Clui.ReadUserInput()
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if len(userInput) == 1 && IsLetter(userInput) && !(strings.EqualFold(userInput, "o")) {
		return userInput
	} else {
		gameFactory.Clui.InformOfNotAvailableMark(userInput)
		return gameFactory.getPlayerOneMark()
	}
}

func (gameFactory *GameFactory) getPlayerTwoMark(playerOneMark string) string {
	gameFactory.Clui.AskPlayerTwoForMark()
	userInput := gameFactory.Clui.ReadUserInput()
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if len(userInput) == 1 && IsLetter(userInput) && !(strings.EqualFold(userInput, playerOneMark)) {
		return userInput
	} else {
		gameFactory.Clui.InformOfNotAvailableMark(userInput)
		return gameFactory.getPlayerTwoMark(playerOneMark)
	}
}