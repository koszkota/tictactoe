package player

import (
	"strconv"
	"tictactoe/board"
)

type Computer struct {
	Mark string
}

func (computer Computer) PickMove (aBoard board.Board) string {
	maxPlayerMark := aBoard.GetActivePlayerMark()
	minPlayerMark := aBoard.GetPassivePlayerMark()
	boardClone := aBoard.MakeACloneOfItself()
	return strconv.Itoa(computer.miniMaxAlgorithm(boardClone, 0, maxPlayerMark, minPlayerMark))
}

func (computer Computer) GetMark() string {
	return computer.Mark
}

func (computer Computer) miniMaxAlgorithm(aBoard board.Board, depth int, maxPlayerMark string, minPlayerMark string) int {
	var maxValueOfPlace = 10
	var tieValueOfPlace = 0

	if aBoard.IsWon(maxPlayerMark) {
		return maxValueOfPlace - depth
	} else if aBoard.IsWon(minPlayerMark) {
		return - (maxValueOfPlace - depth)
	} else if aBoard.IsTie() {
		return tieValueOfPlace
	}

	var freeCells = aBoard.GetFreeCells()

	if computer.isMaxPlayerDepth(depth) {
		var bestScoreMaxPlayer = -1000
		bestPlace := "temporary"
		for _, cell := range freeCells {
			cellValueAsInteger, _ := strconv.Atoi(cell)
			clonedBoard := aBoard.MakeACloneOfItself()
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
			clonedBoard := aBoard.MakeACloneOfItself()
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
