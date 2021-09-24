package board

type Square struct {
	x int
	y int
}

type Cell struct {
	square *Square
	active bool
}

type Board struct {
	matrix [][]*Cell
}

func (board *Board) init(lenght, height int) {
	board.matrix = make([][]*Cell, lenght, height)
}
