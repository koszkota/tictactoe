package game

import (
	"regexp"
	"strings"
	. "tictactoe/board"
	. "tictactoe/clui"
	"tictactoe/player"
)

type Factory struct {
	Clui *Clui
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
	} else if strings.EqualFold(userInput, computeGoesFirst) {
	return factory.getComputerVsHumanGame()
	} else {
		factory.Clui.InformOfInvalidInput(userInput)
		return factory.getMixedPlayersGame()
	}
}

func (factory *Factory) getHumanVsHumanGame() *Game {
	playerOneMark := factory.getPlayerMark("one", "")
	playerTwoMark := factory.getPlayerMark("two", playerOneMark)
	playerOne, playerTwo := factory.createPlayers(player.HumanType, player.HumanType, playerOneMark, playerTwoMark )
	aBoard := factory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{factory.Clui, &aBoard, playerOne, playerTwo}
}

func (factory *Factory) getComputerVsComputerGame() *Game {
	playerOne, playerTwo := factory.createPlayers(player.ComputerType, player.ComputerType, defaultPlayerOneMark, defaultPlayerTwoMark)
	aBoard := factory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{factory.Clui, &aBoard, playerOne, playerTwo}
}

func (factory *Factory) getHumanVsComputerGame() *Game {
	playerOneMark := factory.getPlayerMark("one", defaultPlayerTwoMark)
	playerOne, playerTwo := factory.createPlayers(player.HumanType, player.ComputerType, playerOneMark, defaultPlayerTwoMark)
	aBoard := factory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{factory.Clui, &aBoard, playerOne, playerTwo}
}

func (factory *Factory) getComputerVsHumanGame() *Game {
	playerTwoMark := factory.getPlayerMark("two", defaultPlayerOneMark)
	playerOne, playerTwo := factory.createPlayers(player.ComputerType, player.HumanType, defaultPlayerOneMark, playerTwoMark)
	aBoard := factory.createBoard(playerOne.GetMark(), playerTwo.GetMark())
	return &Game{factory.Clui, &aBoard, playerOne, playerTwo}
}

func (factory *Factory) getPlayerMark(playerOrder, forbiddenMark string) string {
	factory.Clui.AskPlayerForMark(playerOrder)
	userInput := factory.Clui.ReadUserInput()
	if len(userInput) == 1 && stringIsLiteral(userInput) && !(strings.EqualFold(userInput, forbiddenMark)) {
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

func (factory *Factory) createBoard(playerOneMark, playerTwoMark string) Board {
	marksRepo := MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	aBoard := MakeBoard(3, &marksRepo)
	return aBoard
}

func stringIsLiteral(input string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(input)
}