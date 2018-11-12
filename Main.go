package main

import (
	"os"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/game"
	"tictactoe/player"
)

func main() {
	cluiReader := clui.MakeClui(os.Stdin, clui.MakeConsoleWriter())
	board := board.MakeBoard(3)
	playerX := player.Human{"X"}
	playerY := player.Human{"Y"}
	aGame := game.MakeGame(cluiReader, &board, playerX, playerY)

	aGame.Play()
}