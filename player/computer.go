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
	boardCopy := aboard
	boardCopy.Cells = make([]string, len(aboard.Cells))
	copy(boardCopy.Cells, aboard.Cells)

	return strconv.Itoa(computer.miniMaxAlgorithm(boardCopy, 0, maxPlayerSign, minPlayerSign))
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
			newBoard := computer.boardClone(aboard, maxPlayerSign, cellValueAsInteger)
			output := computer.miniMaxAlgorithm(newBoard, depth + 1, maxPlayerSign, minPlayerSign)
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
			newBoard := computer.boardClone(aboard, minPlayerSign, cellValueAsInteger)
			output := computer.miniMaxAlgorithm(newBoard, depth +1, maxPlayerSign, minPlayerSign)
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

func (computer Computer) boardClone(oldBoard board.Board, mark string, place int) board.Board {
	newBoard := oldBoard
	newBoard.Cells = make([]string, len(oldBoard.Cells))
	copy(newBoard.Cells, oldBoard.Cells)
	newBoard.PutMarkOnBoard(mark, place -1)
	return newBoard
}
