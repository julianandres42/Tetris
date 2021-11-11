package board

type initializationMapper map[int]func(int) ([]*Position, int, ShapeType)

var mapperInitialization = initializationMapper{
	int(Cube): func(start int) ([]*Position, int, ShapeType) {
		positions := initializePositions(4)
		positions[0].SetX(start)
		positions[0].SetY(0)
		positions[1].SetX(start)
		positions[1].SetY(1)
		positions[2].SetX(start + 1)
		positions[2].SetY(0)
		positions[3].SetX(start + 1)
		positions[3].SetY(1)
		return positions, 0, Cube
	},
	int(Ele): func(start int) ([]*Position, int, ShapeType) {
		positions := initializePositions(3)
		positions[0].SetX(start)
		positions[0].SetY(0)
		positions[1].SetX(start)
		positions[1].SetY(1)
		positions[2].SetX(start + 1)
		positions[2].SetY(1)
		return positions, 4, Ele
	},
	int(Bar): func(start int) ([]*Position, int, ShapeType) {
		positions := initializePositions(3)
		positions[0].SetX(start)
		positions[0].SetY(0)
		positions[1].SetX(start)
		positions[1].SetY(1)
		positions[2].SetX(start)
		positions[2].SetY(2)
		return positions, 2, Bar
	},
}

func initializePositions(size int) []*Position {
	positions := make([]*Position, size)
	for h := range positions {
		positions[h] = &Position{}
	}
	return positions
}
