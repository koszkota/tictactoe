package player

import (
	"strconv"
	. "tictactoe/board"
)

type Computer struct {
	Mark string
}

func (computer Computer) PickMove (board Board) string {
	maxPlayerMark := board.GetActivePlayerMark()
	minPlayerMark := board.GetPassivePlayerMark()
	boardClone := board.MakeACloneOfItself()
	return strconv.Itoa(computer.miniMaxAlgorithm(boardClone, 0, maxPlayerMark, minPlayerMark))
}

func (computer Computer) GetMark() string {
	return computer.Mark
}

func (computer Computer) GetType() int {
	return ComputerType
}

func (computer Computer) miniMaxAlgorithm(board Board, depth int, maxPlayerMark string, minPlayerMark string) int {

	const(
		maxValueOfPlaceOnBoard = 10
		tieValueOfPlaceOnBoard = 0
	)

	if board.IsWon(maxPlayerMark) {
		return maxValueOfPlaceOnBoard - depth
	} else if board.IsWon(minPlayerMark) {
		return - (maxValueOfPlaceOnBoard - depth)
	} else if board.IsTie() {
		return tieValueOfPlaceOnBoard
	}

	var freeCells = board.GetFreeCells()

	if computer.isMaxPlayerDepth(depth) {
		var bestScoreMaxPlayer = -1000
		bestPlace := "temporary"
		for _, cell := range freeCells {
			cellValueAsInteger, _ := strconv.Atoi(cell)
			clonedBoard := board.MakeACloneOfItself()
			clonedBoard.PutMarkOnBoard(maxPlayerMark, cellValueAsInteger -1)
			output := computer.miniMaxAlgorithm(clonedBoard, depth + 1, maxPlayerMark, minPlayerMark)
			if output > bestScoreMaxPlayer {
				bestPlace = cell
				bestScoreMaxPlayer = output
			}
		}
		if depth == 0 {
			 bestPlaceAsInteger, _ := strconv.Atoi(bestPlace)
			return bestPlaceAsInteger
		} else {
			return bestScoreMaxPlayer
		}
	} else {
		var bestScoreMinPlayer = 1000
		for _, cell := range freeCells {
			cellValueAsInteger, _ := strconv.Atoi(cell)
			clonedBoard := board.MakeACloneOfItself()
			clonedBoard.PutMarkOnBoard(minPlayerMark, cellValueAsInteger -1)
			output := computer.miniMaxAlgorithm(clonedBoard, depth +1, maxPlayerMark, minPlayerMark)
			if output < bestScoreMinPlayer {
				bestScoreMinPlayer = output
			}
		}
		return bestScoreMinPlayer
	}
}

func (computer Computer) isMaxPlayerDepth(depth int) bool {
	return depth % 2 == 0
}