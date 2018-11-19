package player

import (
	. "../clui"
)

type Factory struct {
	Clui *Clui
	ThinkingTimer *ThinkingTimer
}

func (factory *Factory) Create(player int, mark string) Player {
	if player == HumanType {
		return Human{mark, factory.Clui}
	} else {
		return Computer{mark, factory.Clui, factory.ThinkingTimer}
	}
}