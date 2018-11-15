package player

import (
	"strings"
	"testing"
	"tictactoe/clui"
	"tictactoe/testhelper"
)

func TestFactoryCreatesHumanPlayer(t *testing.T) {
	factory := getFactory()

	player := factory.Create(HumanType, "X")

	matchers.EqualLiterals(t, player.GetType(), HumanType)
}

func TestFactoryCreatesComputerPlayer(t *testing.T) {
	factory := getFactory()

	player := factory.Create(ComputerType, "X")

	matchers.EqualLiterals(t, player.GetType(), ComputerType)
}

func getFactory() *Factory {
	stubWriter := &clui.StubWriter{}
	clui := clui.NewClui(strings.NewReader(""), stubWriter)
	return &Factory{Clui: *clui}
}