package game

import (
	"regexp"
	"strings"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
)

type Factory struct {
	Clui clui.Clui
	PlayerFactory *player.Factory
}

const (
	humanHumanGame       = "1"
	humanComputerGame    = "2"
	computerComputerGame = "3"
	humanGoesFirst       = "h"
	computeGoesFirst     = "c"
	defaultPlayerOneMark = "X"
	defaultPlayerTwoMark = "O"
)

func (gameFactory *Factory) CreateGame() *Game {
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

func (gameFactory *Factory) getMixedPlayersGame() *Game {
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

func (gameFactory *Factory) getHumanVsHumanGame() *Game {
	playerOneMark := gameFactory.getPlayerMark("one", "")
	playerTwoMark := gameFactory.getPlayerMark("two", playerOneMark)
	playerOne, playerTwo := gameFactory.createPlayers(player.HumanType, player.HumanType, playerOneMark, playerTwoMark )
	aBoard := gameFactory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{gameFactory.Clui, &aBoard, playerOne, playerTwo}
}

func (gameFactory *Factory) getComputerVsComputerGame() *Game {
	playerOne, playerTwo := gameFactory.createPlayers(player.ComputerType, player.ComputerType, defaultPlayerOneMark, defaultPlayerTwoMark )
	aBoard := gameFactory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{gameFactory.Clui, &aBoard, playerOne, playerTwo}
}

func (gameFactory *Factory) getHumanVsComputerGame() *Game {
	playerOneMark := gameFactory.getPlayerMark("one", defaultPlayerTwoMark)
	playerOne, playerTwo := gameFactory.createPlayers(player.HumanType, player.ComputerType, playerOneMark, defaultPlayerTwoMark )
	aBoard := gameFactory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{gameFactory.Clui, &aBoard, playerOne, playerTwo}
}

func (gameFactory *Factory) getComputerVsHumanGame() *Game {
	playerTwoMark := gameFactory.getPlayerMark("two", defaultPlayerOneMark)
	playerOne, playerTwo := gameFactory.createPlayers(player.ComputerType, player.HumanType, defaultPlayerOneMark, playerTwoMark )
	aBoard := gameFactory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{gameFactory.Clui, &aBoard, playerOne, playerTwo}
}

func (gameFactory *Factory) getPlayerMark(playerOrder string, forbiddenMark string) string {
	gameFactory.Clui.AskPlayerForMark(playerOrder)
	userInput := gameFactory.Clui.ReadUserInput()
	if len(userInput) == 1 && stringIsLiteral(userInput) && !(strings.EqualFold(userInput, forbiddenMark)) {
		return userInput
	} else {
		gameFactory.Clui.InformOfNotAvailableMark(userInput)
		return gameFactory.getPlayerMark(playerOrder, forbiddenMark)
	}
}

func (gameFactory *Factory) createPlayers(playerOneType , playerTwoType int, playerOneMark, playerTwoMark string) (player.Player, player.Player) {
	playerOne := gameFactory.PlayerFactory.Create(playerOneType, playerOneMark)
	playerTwo := gameFactory.PlayerFactory.Create(playerTwoType, playerTwoMark)
	return playerOne, playerTwo
}

func (gameFactory *Factory) createBoard(playerOneMark, playerTwoMark string) board.Board {
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	aBoard := board.MakeBoard(3, &marksRepo)
	return aBoard
}

func stringIsLiteral(input string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(input)
}