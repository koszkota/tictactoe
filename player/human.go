package player

import "tictactoe/board"

type Human struct {
	Mark string
}

func (human Human) PickMove(board board.Board) int {
	return 2
}