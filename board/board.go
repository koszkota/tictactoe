package board

type Board struct {
	cells []string
	size int
}

func MakeBoard(boardSize int) Board {
	aBoard := Board{size: boardSize}

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

func (board *Board) isCellAvailable(cellIndex int) bool {
	return board.cells[cellIndex] == " "
}

func (board *Board) getRows() [][]string {
	var cellsSplitIntoRows [][]string
	for i := 0; i < len(board.cells); i += board.size {
		end := i + board.size
		cellsSplitIntoRows = append(cellsSplitIntoRows, board.cells[i:end])
	}
	return cellsSplitIntoRows
}

func (board Board) getColumns() [][]string {
	var cellsSplitIntoColumns [][]string
	rows := board.getRows()
	for i := 0; i < board.size; i++ {
		var singleColumn []string
		for j := 0; j < len(rows); j++ {
			singleColumn = append(singleColumn, rows[j][i])
		}
		cellsSplitIntoColumns = append(cellsSplitIntoColumns, singleColumn)
	}
	return cellsSplitIntoColumns
}