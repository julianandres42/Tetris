package shape

type initializationMapper map[int]func(int) []*Position

var mapperInitialization = initializationMapper{
	int(Cube): func(start int) []*Position {
		positions := initializePositions(4)
		positions[0].SetX(start)
		positions[0].SetY(0)
		positions[1].SetX(start)
		positions[1].SetY(1)
		positions[2].SetX(start + 1)
		positions[2].SetY(0)
		positions[3].SetX(start + 1)
		positions[3].SetY(1)
		return positions
	},
	int(Ele): func(start int) []*Position {
		positions := initializePositions(3)
		positions[0].SetX(start)
		positions[0].SetY(0)
		positions[1].SetX(start)
		positions[1].SetY(1)
		positions[2].SetX(start + 1)
		positions[2].SetY(1)
		return positions
	},
	int(Bar): func(start int) []*Position {
		positions := initializePositions(3)
		positions[0].SetX(start)
		positions[0].SetY(0)
		positions[1].SetX(start)
		positions[1].SetY(1)
		positions[2].SetX(start)
		positions[2].SetY(2)
		return positions
	},
}

func initializePositions(size int) []*Position {
	positions := make([]*Position, size)
	for h := range positions {
		positions[h] = &Position{}
	}
	return positions
}
