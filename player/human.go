package player

import "tictactoe/board"

type Human struct {
	mark string
}

func (human Human) PickMove(board board.Board) int {
	return 2
}

func (human Human) GetMark() string {
	return human.mark
}