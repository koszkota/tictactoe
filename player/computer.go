package player

import "tictactoe/board"

type Computer struct {
	mark string
}

func (computer Computer) PickMove (board board.Board) int {
	return 3
}

func (computer Computer) GetMark() string {
	return computer.mark
}