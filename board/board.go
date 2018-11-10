package board

type Board struct {
	cells [9]string
}

func (board *Board) putMarkOnBoard(mark string, cellIndex int) {
	board.cells[cellIndex] = mark
}
