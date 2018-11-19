package runner

import (
	"os"
	. "tictactoe/src/clui"
	. "tictactoe/src/clui/writer"
	"tictactoe/src/game"
	. "tictactoe/src/gamescontroller"
	"tictactoe/src/gamescontroller/controllerstatus"
	"tictactoe/src/player"
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
