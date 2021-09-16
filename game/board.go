package game

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
