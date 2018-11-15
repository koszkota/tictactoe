package player

import (
	"tictactoe/clui"
)

type Factory struct {
	Clui clui.Clui
}

func (factory *Factory) Create(player, mark string) Player {
	if player == "Human" {
		return Human{mark, factory.Clui}
	} else {
		return Computer{mark}
	}
}