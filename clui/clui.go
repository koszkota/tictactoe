package clui

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type CluiReader struct {
	reader *bufio.Reader
}

func makeCluiReader(input io.Reader) CluiReader {
	reader := bufio.NewReader(input)
	return CluiReader{reader: reader}
}

func helloPlayers() {
	fmt.Println("Hello and welcome to tic tac toe")
}

func informOfMove(positionOnBoard int, mark string)  {
	fmt.Println("Player " + mark + " picked position " + strconv.Itoa(positionOnBoard))
}

func informOfWinner(winnerMark string) {
	fmt.Println("Player " + winnerMark + " won!")
}

func informOfTie()  {
	fmt.Println("It's a tie!")
}

func askForMove(mark string) {
	fmt.Println("Player " + mark + ", pick a position")
}

func informOfInvalidMove() {
	fmt.Println("This move is not available")
}

func ShowBoard(rows [][]string) {
	boardString := buildBoardString(rows)
	fmt.Print(boardString)
}

func (clui CluiReader) ReadUserInput() string {
	text, _ := clui.reader.ReadString('\n')
	return text
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