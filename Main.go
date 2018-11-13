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
	marksRepo := board.MarksRepo{PlayerOneMark: "X", PlayerTwoMark: "Y"}
	board := board.MakeBoard(3, &marksRepo)
	playerY := player.Human{Mark: "X", Clui:clui}
	playerX := player.Human{Mark: "Y", Clui:clui}
	aGame := game.MakeGame(clui, &board, playerX, playerY)

	aGame.Play()
}