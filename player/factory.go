package player

import (
	"tictactoe/clui"
)

type Factory struct {
	Clui clui.Clui
}

func (factory *Factory) Create(player int, mark string) Player {
	if player == HumanType {
		return Human{mark, factory.Clui}
	} else {
		return Computer{mark}
	}
}