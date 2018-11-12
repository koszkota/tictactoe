package game

import (
	"strconv"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/player"
)

type Game struct {
	clui      clui.Clui
	board     *board.Board
	playerOne player.Player
	playerTwo player.Player
}

func MakeGame(clui clui.Clui, board *board.Board, playerOne player.Player, playerTwo player.Player) Game {
	return Game{clui: clui, board: board, playerOne: playerOne, playerTwo: playerTwo}
}

func (game *Game) Play() {
	game.clui.HelloPlayers()
	game.clui.ShowBoard(game.board.GetRows())
	for !(game.board.IsGameOver(game.playerOne.GetMark(), game.playerTwo.GetMark())) {
		game.playTurn()
	}
	game.announceResult()
}

func (game Game) playTurn() {
	currentPlayer := game.getCurrentPlayer()
	move := currentPlayer.PickMove(*game.board)
	moveReadyForBoard := game.transformMoveForTheBoard(move)
	game.board.PutMarkOnBoard(currentPlayer.GetMark(), moveReadyForBoard)
	game.clui.InformOfMove(move, currentPlayer.GetMark())
	game.clui.ShowBoard(game.board.GetRows())
}

func (game Game) announceResult() {
	if game.board.IsWon(game.playerOne.GetMark()) {
		game.clui.InformOfWinner(game.playerOne.GetMark())
	} else if  game.board.IsWon(game.playerTwo.GetMark()) {
		game.clui.InformOfWinner(game.playerTwo.GetMark())
	} else {
		game.clui.InformOfTie()
	}
}

func (game Game) transformMoveForTheBoard(move string) int {
	moveAsInteger,_  := strconv.Atoi(move)
	return moveAsInteger - 1
}

func (game Game) getCurrentPlayer() player.Player {
	currentPlayerMark := game.board.GetCurrentPlayerMark(game.playerOne.GetMark(), game.playerTwo.GetMark())
	if currentPlayerMark == game.playerOne.GetMark() {
		return game.playerOne
	} else {
		return game.playerTwo
	}
}