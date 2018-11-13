package player

import (
	"strconv"
	"tictactoe/board"
)

type Computer struct {
	Mark string
}

func (computer Computer) PickMove (aboard board.Board) string {
	maxPlayerSign := aboard.GetActivePlayerSign()
	minPlayerSign := aboard.GetPassivePlayerSign()
	boardClone := computer.cloneBoard(aboard)
	return strconv.Itoa(computer.miniMaxAlgorithm(boardClone, 0, maxPlayerSign, minPlayerSign))
}

func (computer Computer) GetMark() string {
	return computer.Mark
}

func (computer Computer) miniMaxAlgorithm(aboard board.Board, depth int, maxPlayerSign string, minPlayerSign string) int {
	if aboard.IsWon(maxPlayerSign) {
		return 10 - depth
	} else if aboard.IsWon(minPlayerSign) {
		return - (10 - depth)
	} else if aboard.IsTie() {
		return 0
	}

	var freeCells = aboard.GetFreeCells()

	if computer.isMaxPlayerDepth(depth) {
		var bestScoreMaxPlayer = -1000
		bestPlace := "temporary"
		for _, cell := range freeCells {
			cellValueAsInteger, _ := strconv.Atoi(cell)
			clonedBoard := computer.cloneBoard(aboard)
			clonedBoard.PutMarkOnBoard(maxPlayerSign, cellValueAsInteger -1)
			output := computer.miniMaxAlgorithm(clonedBoard, depth + 1, maxPlayerSign, minPlayerSign)
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
			clonedBoard := computer.cloneBoard(aboard)
			clonedBoard.PutMarkOnBoard(minPlayerSign, cellValueAsInteger -1)
			output := computer.miniMaxAlgorithm(clonedBoard, depth +1, maxPlayerSign, minPlayerSign)
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

func (computer Computer) cloneBoard(oldBoard board.Board) board.Board {
	newBoard := oldBoard
	newBoard.Cells = make([]string, len(oldBoard.Cells))
	copy(newBoard.Cells, oldBoard.Cells)
	return newBoard
}
