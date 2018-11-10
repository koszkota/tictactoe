package player

import "tictactoe/board"

type Player interface {
	PickMove(board board.Board) int
}