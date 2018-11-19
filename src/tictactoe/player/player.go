package player

import . "tictactoe/src/tictactoe/board"

type Player interface {
	PickMove(board Board) string
	GetMark() string
	GetType() int
}