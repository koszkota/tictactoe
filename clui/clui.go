package clui

import "strconv"

func HelloPlayers() string {
	return "Hello and welcome to tic tac toe"
}

func InformOfMove(positionOnBoard int, mark string) string {
	return "Player " + mark + " picked position " + strconv.Itoa(positionOnBoard)
}

func InformOfWinner(winnerMark string) string {
	return "Player " + winnerMark + "won!"
}

func InformOfTie() string {
	return "It's a tie!"
}

func AskForMove(mark string) string {
	return "Player " + mark + ", pick a position"
}

func InformOfInvalidMove() string {
	return "This move is not available"
}