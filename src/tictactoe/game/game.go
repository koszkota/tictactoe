package game

import (
	"strconv"
	. "tictactoe/src/tictactoe/board"
	. "tictactoe/src/tictactoe/clui"
	"tictactoe/src/tictactoe/player"
)

type Game struct {
	clui      *Clui
	board     *Board
	playerOne player.Player
	playerTwo player.Player
}

func (game *Game) Play() {
	game.clui.InformOfBeginningOfGame()
	game.clui.ShowBoard(game.board.GetRows())
	for !(game.board.IsGameOver(game.playerOne.GetMark(), game.playerTwo.GetMark())) {
		game.playTurn()
	}
	game.announceResult()
}

func (game Game) playTurn() {
	currentPlayer := game.getCurrentPlayer()
	cellValuePickedByPlayer := currentPlayer.PickMove(*game.board)
	moveReadyForBoard := game.transformMove(cellValuePickedByPlayer)
	game.board.PutMarkOnBoard(currentPlayer.GetMark(), moveReadyForBoard)
	game.clui.InformOfMove(cellValuePickedByPlayer, currentPlayer.GetMark())
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

func (game Game) transformMove(move string) int {
	moveAsInteger,_  := strconv.Atoi(move)
	return moveAsInteger - 1
}

func (game Game) getCurrentPlayer() player.Player {
	currentPlayerMark := game.board.GetActivePlayerMark()
	if currentPlayerMark == game.playerOne.GetMark() {
		return game.playerOne
	} else {
		return game.playerTwo
	}
}

