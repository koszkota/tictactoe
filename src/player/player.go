package player

import . "tictactoe/src/board"

type Player interface {
	PickMove(board Board) string
	GetMark() string
	GetType() int
}