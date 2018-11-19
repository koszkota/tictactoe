package clui

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
	. "../clui/writer"
)

type Clui struct {
	reader *bufio.Reader
	writer Writer
}

func NewClui(input io.Reader, writer Writer) *Clui {
	reader := bufio.NewReader(input)
	return &Clui{reader: reader, writer: writer}
}

func (clui Clui) InformOfBeginningOfGame() {
	clui.writer.Write("Let's start the game!")
}

func (clui Clui) AskPlayerForMark(player string) {
	clui.writer.Write("Player " + player + ", please pick one-letter mark")
}

func (clui Clui) InformOfInvalidInput(input string) {
	clui.writer.WriteWarning("Option " + input + " is not allowed.")
}

func (clui Clui) ShowMainMenu() {
	clui.writer.Write("To play a game enter YES, to exit enter NO.")
}

func (clui Clui) AskForGameMode() {
	clui.writer.Write("To play Human vs Human enter 1. To play Human vs Computer enter 2. To see two computers playing enter 3.")
}

func (clui Clui) AskWhoGoesFirst() {
	clui.writer.Write("If Human player should go first, enter H; if computer, enter C.")
}

func (clui Clui) InformOfNotAvailableMark(mark string) {
	clui.writer.WriteWarning("Mark " + mark + " is not available, please pick another one.")
}

func (clui Clui) InformOfMove(pickedMove string, mark string)  {
	clui.writer.Write("Player " + mark + " picked position " + pickedMove)
}

func (clui Clui) InformOfWinner(winnerMark string) {
	clui.writer.Write("Player " + winnerMark + " won!")
}

func (clui Clui) InformOfTie()  {
	clui.writer.Write("It's a tie!")
}

func (clui Clui) AskForMove(mark string) {
	clui.writer.Write("Player " + mark + ", pick a position")
}

func (clui Clui) InformOfInvalidMove() {
	clui.writer.WriteWarning("This move is not available")
}

func (clui Clui) InformOfThinkingComputer(thinkingTime time.Duration) {
	clui.writer.Write("Computer is thinking...")
	time.Sleep(thinkingTime * time.Second)
}

func (clui Clui) ShowBoard(rows [][]string) {
	boardString := buildBoardString(rows)
	clui.writer.Write(boardString)
}

func (clui Clui) ReadUserInput() string {
	text, _ := clui.reader.ReadString('\n')
	return strings.TrimRight(text, "\n\r")
}

func buildBoardString(rows [][]string) string {
	var boardString string
	for _, row := range rows {
		for j, cell := range row {
			boardString += fmt.Sprintf("%3s", cell)
			if j < (len(row) - 1) {
				boardString += "  |"
			}
		}
		boardString += "\n"
	}
	return boardString
}
