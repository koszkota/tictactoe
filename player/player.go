package player

import "tictactoe/board"

type Player interface {
	PickMove(board board.Board) string
	GetMark() string
	GetType() string
}