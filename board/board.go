package board

type Board struct {
	cells []string
}

func MakeBoard(boardSize int) Board {
	aBoard := Board{}
	aBoard.initializeEmptyCells(boardSize)
	return aBoard
}

func (board *Board) initializeEmptyCells(boardSize int) {
	for i := 0; i < (boardSize * boardSize); i++ {
		board.cells = append(board.cells, " ")
	}
}

func (board *Board) putMarkOnBoard(mark string, cellIndex int) {
	board.cells[cellIndex] = mark
}

func (board *Board) cellIsAvailable(cellIndex int) bool {
	return board.cells[cellIndex] == " "
}
