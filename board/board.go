package board

type Square struct {
	X int
	Y int
}

type Cell struct {
	Square *Square
	Active bool
}

type Board struct {
	Matrix [][]*Cell
}

func (board *Board) Init(lenght, height int) {
	board.Matrix = make([][]*Cell, lenght)
	initSquare := &Square{120, 850}
	for i := range board.Matrix {
		board.Matrix[i] = make([]*Cell, height)
		for h := range board.Matrix[i] {
			board.Matrix[i][h] = &Cell{&Square{initSquare.X, initSquare.Y}, false}
			initSquare.X += 40
		}
		initSquare.X = 120
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

func (board *Board) UpdateShape(shape *Shape, value bool) {
	for _, element := range shape.Positions {
		board.Matrix[element.GetY()][element.GetX()].Active = value
	}
}

func isLine(cellRow []*Cell) bool {
	for _, cell := range cellRow {
		if !cell.Active {
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
		destiny[i].Active = source[i].Active
		source[i].Active = false
	}
}
