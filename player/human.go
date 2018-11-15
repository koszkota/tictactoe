package player

import (
	"strconv"
	"tictactoe/board"
	"tictactoe/clui"
)

type Human struct {
	Mark string
	Clui *clui.Clui
}

func (human Human) PickMove(board board.Board) string {
	human.Clui.AskForMove(human.Mark)
	moveAsString :=human.Clui.ReadUserInput()
	moveAsInt, _ := strconv.Atoi(moveAsString)
	if board.IsMoveValid(moveAsInt - 1) {
		return moveAsString
	} else {
		human.Clui.InformOfInvalidMove()
		return human.PickMove(board)
	}
}

func (human Human) GetMark() string {
	return human.Mark
}

func (human Human) GetType() int {
	return HumanType
}