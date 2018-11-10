package player

import "tictactoe/board"

type Computer struct {
	Mark string
}

func (computer Computer) PickMove (board board.Board) int {
	return 3
}