package player

import "tictactoe/board"

type Computer struct {
	Mark string
}

func (computer Computer) PickMove (board board.Board) string {
	return "3"
}

func (computer Computer) GetMark() string {
	return computer.Mark
}