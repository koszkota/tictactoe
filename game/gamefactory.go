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

const(
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
	playerOne, playerTwo := gameFactory.createPlayers("Human", "Human", playerOneMark, playerTwoMark )
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *Factory) getComputerVsComputerGame() *Game {
	playerOne, playerTwo := gameFactory.createPlayers("Computer", "Computer", defaultPlayerOneMark, defaultPlayerTwoMark )
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *Factory) getHumanVsComputerGame() *Game {
	playerOneMark := gameFactory.getPlayerMark("one", defaultPlayerTwoMark)
	playerOne, playerTwo := gameFactory.createPlayers("Human", "Computer", playerOneMark, defaultPlayerTwoMark )
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwo.GetMark()}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *Factory) getComputerVsHumanGame() *Game {
	playerTwoMark := gameFactory.getPlayerMark("two", defaultPlayerOneMark)
	playerOne, playerTwo := gameFactory.createPlayers("Computer", "Human", defaultPlayerOneMark, playerTwoMark )
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwoMark}
	aBoard := board.MakeBoard(3, &marksRepo)
	aGame := MakeGame(gameFactory.Clui, &aBoard, playerOne, playerTwo)
	return &aGame
}

func (gameFactory *Factory) getPlayerMark(playerOrder string, forbiddenMark string) string {
	gameFactory.Clui.AskPlayerForMark(playerOrder)
	userInput := gameFactory.Clui.ReadUserInput()
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if len(userInput) == 1 && IsLetter(userInput) && !(strings.EqualFold(userInput, forbiddenMark)) {
		return userInput
	} else {
		gameFactory.Clui.InformOfNotAvailableMark(userInput)
		return gameFactory.getPlayerMark(playerOrder, forbiddenMark)
	}
}

func (gameFactory *Factory) createPlayers(playerOneType , playerTwoType, playerOneMark, playerTwoMark string) (player.Player, player.Player) {
	playerOne := gameFactory.PlayerFactory.Create(playerOneType, playerOneMark)
	playerTwo := gameFactory.PlayerFactory.Create(playerTwoType, playerTwoMark)
	return playerOne, playerTwo
}