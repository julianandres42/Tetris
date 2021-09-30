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
