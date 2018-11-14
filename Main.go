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
	playerOne := player.Computer{Mark: "X"}
	playerTwo := player.Computer{Mark: "Y"}
	marksRepo := board.MarksRepo{PlayerOneMark: playerOne.GetMark(), PlayerTwoMark: playerTwo.GetMark()}
	board := board.MakeBoard(3, &marksRepo)

	aGame := game.MakeGame(clui, &board, playerTwo, playerOne)

	aGame.Play()
}