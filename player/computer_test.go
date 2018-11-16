package player

import (
	"strings"
	"testing"
	"tictactoe/board"
	"tictactoe/clui"
	"tictactoe/clui/writer"
	"tictactoe/testhelper"
)

var computerPlayer = getComputerPlayer()

func TestComputerPlayerHasMark(t *testing.T) {
	matchers.EqualLiterals(t, "Y", computerPlayer.Mark)
}

func TestComputerPlayerHasType(t *testing.T) {
	matchers.EqualLiterals(t, ComputerType, computerPlayer.GetType())
}

func TestGetMarkReturnsComputerMark(t *testing.T) {
	matchers.EqualLiterals(t, "Y", computerPlayer.GetMark())
}

func TestPickMethodReturnsWinningPosition(t *testing.T) {
	aBoard := createBoard(3, "X", "Y")
	aBoard.PutMarkOnBoard("X", 0)
	aBoard.PutMarkOnBoard("X", 1)
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "3", pickedCell)
}

func TestPickMethodBlocksAWinningOpponent(t *testing.T) {
	aBoard := createBoard(3, "X", "Y")
	aBoard.PutMarkOnBoard("X", 0)
	aBoard.PutMarkOnBoard("Y", 1)
	aBoard.PutMarkOnBoard("X", 3)
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "7", pickedCell)
}

func TestPicksTopLeftPositionOn2x2Board(t *testing.T) {
	aBoard := createBoard(2, "X", "Y")
	pickedCell := computerPlayer.PickMove(aBoard)

	matchers.EqualLiterals(t, "1", pickedCell)
}

func createBoard(size int, playerOneMark string, playerTwoMark string) board.Board {
	marksRepo := board.MarksRepo{PlayerOneMark: playerOneMark, PlayerTwoMark: playerTwoMark}
	return board.MakeBoard(size, &marksRepo)
}

func getComputerPlayer() Computer {
	var stubWriter = new(writer.StubWriter)
	clui := clui.NewClui(strings.NewReader(""), stubWriter)
	thinkingTimer := &ThinkingTimer{ThinkingTime: 0}
	return Computer{Mark: "Y", Clui: clui, ThinkingTimer: thinkingTimer}
}