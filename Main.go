package main

import (
	"os"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/player"
)

func main() {
	clui := clui.MakeClui(os.Stdin, clui.MakeConsoleWriter())
	board := board.MakeBoard(3)
	playerX := player.Human{Mark: "X", Clui: clui}
	playerY := player.Human{Mark: "Y", Clui: clui}
	aGame := game.MakeGame(clui, &board, playerX, playerY)

	aGame.Play()
}