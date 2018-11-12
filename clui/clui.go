package clui

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Clui struct {
	reader *bufio.Reader
	writer Writer
}

func MakeClui(input io.Reader, writer Writer) Clui {
	reader := bufio.NewReader(input)
	return Clui{reader: reader, writer: writer}
}

func (clui Clui) HelloPlayers() {
	clui.writer.Write("Hello and welcome to tic tac toe")
}

func (clui Clui) informOfMove(positionOnBoard int, mark string)  {
	clui.writer.Write("Player " + mark + " picked position " + strconv.Itoa(positionOnBoard))
}

func (clui Clui) informOfWinner(winnerMark string) {
	clui.writer.Write("Player " + winnerMark + " won!")
}

func (clui Clui) InformOfTie()  {
	clui.writer.Write("It's a tie!")
}

func (clui Clui) askForMove(mark string) {
	clui.writer.Write("Player " + mark + ", pick a position")
}

func (clui Clui) informOfInvalidMove() {
	clui.writer.Write("This move is not available")
}

func (clui Clui) ShowBoard(rows [][]string) {
	boardString := buildBoardString(rows)
	clui.writer.Write(boardString)
}

func (clui Clui) ReadUserInput() string {
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