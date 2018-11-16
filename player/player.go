package player

import . "tictactoe/board"

type Player interface {
	PickMove(board Board) string
	GetMark() string
	GetType() int
}