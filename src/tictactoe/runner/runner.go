package runner

import (
	"os"
	. "tictactoe/src/tictactoe/clui"
	. "tictactoe/src/tictactoe/clui/writer"
	"tictactoe/src/tictactoe/game"
	. "tictactoe/src/tictactoe/gamescontroller"
	"tictactoe/src/tictactoe/gamescontroller/controllerstatus"
	"tictactoe/src/tictactoe/player"
)

func Start() {
	aClui := NewClui(os.Stdin, new(ConsoleWriter))
	runStatus := &controllerstatus.RunStatus{}
	thinkingTimer := &player.ThinkingTimer{ThinkingTime: 3}
	playersFactory := &player.Factory{Clui: aClui, ThinkingTimer: thinkingTimer}
	gameFactory := &game.Factory{Clui: aClui, PlayerFactory: playersFactory}
	gamesController := &GamesController{Clui: aClui, RunStatus: runStatus, GameFactory: gameFactory}
	gamesController.Run()
}
