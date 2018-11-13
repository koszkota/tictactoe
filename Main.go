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
	marksRepo := board.MarksRepo{"X", "Y"}
	board := board.MakeBoard(3, &marksRepo)
	playerX := player.Human{Mark: "X", Clui: clui}
	playerY := player.Computer{Mark: "Y"}
	aGame := game.MakeGame(clui, &board, playerX, playerY)

	aGame.Play()
}