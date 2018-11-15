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
		gameFactory.Clui.InformOfInvalidInput(pick)
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
		gameFactory.Clui.InformOfInvalidInput(pick)
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
	mark := gameFactory.Clui.ReadUserInput()
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if len(mark) == 1 && IsLetter(mark) && !(strings.EqualFold(mark, "o")) {
		return mark
	} else {
		gameFactory.Clui.InformOfNotAvailableMark(mark)
		return gameFactory.getPlayerOneMark()
	}
}

func (gameFactory *GameFactory) getPlayerTwoMark(playerOneMark string) string {
	gameFactory.Clui.AskPlayerTwoForMark()
	mark := gameFactory.Clui.ReadUserInput()
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if len(mark) == 1 && IsLetter(mark) && !(strings.EqualFold(mark, playerOneMark)) {
		return mark
	} else {
		gameFactory.Clui.InformOfNotAvailableMark(mark)
		return gameFactory.getPlayerTwoMark(playerOneMark)
	}
}