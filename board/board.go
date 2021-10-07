package board

type Square struct {
	X int
	Y int
}

type Cell struct {
	Square *Square
	active bool
}

type Board struct {
	Matrix [][]*Cell
}

func (board *Board) Init(lenght, height int) {
	board.Matrix = make([][]*Cell, lenght)
	initSquare := &Square{119, 847}
	for i := range board.Matrix {
		board.Matrix[i] = make([]*Cell, height)
		for h := range board.Matrix[i] {
			board.Matrix[i][h] = &Cell{&Square{initSquare.X, initSquare.Y}, false}
			initSquare.X += 40
		}
		initSquare.X = 119
		initSquare.Y -= 40
	}
}

func (board *Board) EvaluateLines() {
	for i, row := range board.Matrix {
		if isLine(row) {
			board.downRows(i)
		}
	}
}

func isLine(cellRow []*Cell) bool {
	for _, cell := range cellRow {
		if !cell.active {
			return false
		}
	}
	return true
}

func (board *Board) downRows(rowIndex int) {
	for i := rowIndex; i >= 1; i-- {
		switchCellValues(board.Matrix[i-1], board.Matrix[i])
	}
}

func switchCellValues(source []*Cell, destiny []*Cell) {
	for i := range source {
		destiny[i].active = source[i].active
		source[i].active = false
	}
}
