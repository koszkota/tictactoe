package game

import (
	"regexp"
	"strings"
	"tictactoe/src/board"
	"tictactoe/src/clui"
	"tictactoe/src/player"
)

type Factory struct {
	Clui *clui.Clui
	PlayerFactory *player.Factory
}

func (factory Factory) CreateGame() *Game {
	factory.Clui.AskForGameMode()
	userInput := factory.Clui.ReadUserInput()
	switch userInput {
	case humanHumanGame:
		return factory.getHumanVsHumanGame()
	case humanComputerGame:
		return factory.getMixedPlayersGame()
	case computerComputerGame:
		return factory.getComputerVsComputerGame()
	default:
		factory.Clui.InformOfInvalidInput(userInput)
		return factory.CreateGame()
	}
}

func (factory *Factory) getMixedPlayersGame() *Game {
	factory.Clui.AskWhoGoesFirst()
	userInput := factory.Clui.ReadUserInput()
	if strings.EqualFold(userInput, humanGoesFirst) {
		return factory.getHumanVsComputerGame()
	} else if strings.EqualFold(userInput, computerGoesFirst) {
		return factory.getComputerVsHumanGame()
	} else {
		factory.Clui.InformOfInvalidInput(userInput)
		return factory.getMixedPlayersGame()
	}
}

func (factory *Factory) getHumanVsHumanGame() *Game {
	playerOneMark := factory.getPlayerMark("one", "")
	playerTwoMark := factory.getPlayerMark("two", playerOneMark)
	return factory.getGame(player.HumanType, player.HumanType, playerOneMark, playerTwoMark)
}

func (factory *Factory) getComputerVsComputerGame() *Game {
	return factory.getGame(player.ComputerType, player.ComputerType, defaultPlayerOneMark, defaultPlayerTwoMark)
}

func (factory *Factory) getHumanVsComputerGame() *Game {
	playerOneMark := factory.getPlayerMark("one", defaultPlayerTwoMark)
	return factory.getGame(player.HumanType, player.ComputerType, playerOneMark, defaultPlayerTwoMark)
}

func (factory *Factory) getComputerVsHumanGame() *Game {
	playerTwoMark := factory.getPlayerMark("two", defaultPlayerOneMark)
	return factory.getGame(player.ComputerType, player.HumanType, defaultPlayerOneMark, playerTwoMark)
}

func (factory *Factory) getPlayerMark(playerOrder, forbiddenMark string) string {
	factory.Clui.AskPlayerForMark(playerOrder)
	userInput := factory.Clui.ReadUserInput()
	if markIsValid(userInput, forbiddenMark) {
		return userInput
	} else {
		factory.Clui.InformOfNotAvailableMark(userInput)
		return factory.getPlayerMark(playerOrder, forbiddenMark)
	}
}

func (factory *Factory) createPlayers(playerOneType, playerTwoType int, playerOneMark, playerTwoMark string) (player.Player, player.Player) {
	playerOne := factory.PlayerFactory.Create(playerOneType, playerOneMark)
	playerTwo := factory.PlayerFactory.Create(playerTwoType, playerTwoMark)
	return playerOne, playerTwo
}

func (factory *Factory) createBoard(playerOneMark, playerTwoMark string) board.Board {
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	aBoard := board.MakeBoard(3, &marksRepo)
	return aBoard
}

func markIsValid(userMark, forbiddenMark string) bool {
	return len(userMark) == 1 && stringIsLiteral(userMark) && !(strings.EqualFold(userMark, forbiddenMark))
}

func stringIsLiteral(input string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(input)
}

func (factory *Factory) getGame(playerOneType, playerTwoType int, playerOneMark, playerTwoMark string) *Game {
	playerOne, playerTwo := factory.createPlayers(playerOneType, playerTwoType, playerOneMark, playerTwoMark )
	aBoard := factory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{factory.Clui, &aBoard, playerOne, playerTwo}
}