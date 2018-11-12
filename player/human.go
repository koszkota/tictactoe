package player

import (
	"strings"
	"tictactoe/board"
	"tictactoe/clui"
)

type Human struct {
	Mark string
	Clui clui.Clui
}

func (human Human) PickMove(board board.Board) string {
	human.Clui.AskForMove(human.Mark)
	return strings.TrimRight(human.Clui.ReadUserInput(), "\n\r")
}

func (human Human) GetMark() string {
	return human.Mark
}