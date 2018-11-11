package clui

import (
	"fmt"
	"strconv"
)

func HelloPlayers() {
	fmt.Println("Hello and welcome to tic tac toe")
}

func InformOfMove(positionOnBoard int, mark string)  {
	fmt.Println("Player " + mark + " picked position " + strconv.Itoa(positionOnBoard))
}

func InformOfWinner(winnerMark string) {
	fmt.Println("Player " + winnerMark + " won!")
}

func InformOfTie()  {
	fmt.Println("It's a tie!")
}

func AskForMove(mark string) {
	fmt.Println("Player " + mark + ", pick a position")
}

func InformOfInvalidMove() {
	fmt.Println("This move is not available")
}

func ShowBoard(rows [][]string) {
	boardString := buildBoardString(rows)
	fmt.Print(boardString)
}

func buildBoardString(rows [][]string) string {
	var boardString string
	for _, row := range rows {
		for j, cell := range row {
			boardString += fmt.Sprintf("%3s", cell)
			if j < (len(row) -1) {
				boardString += "  |"
			}
		}
		boardString += "\n"
	}
	return boardString
}