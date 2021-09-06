package game

type initializationMapper map[int]func(int) []*Position

var mapperInitialization = initializationMapper{
	int(Cube): func(start int) []*Position {
		positions := initializePositions(4)
		positions[0].setX(start)
		positions[0].setY(0)
		positions[1].setX(start + 1)
		positions[1].setY(0)
		positions[2].setX(start)
		positions[2].setY(1)
		positions[3].setX(start + 1)
		positions[3].setY(1)
		return positions
	},
	int(Ele): func(start int) []*Position {
		positions := initializePositions(3)
		positions[0].setX(start)
		positions[0].setY(0)
		positions[1].setX(start)
		positions[1].setY(1)
		positions[2].setX(start + 1)
		positions[2].setY(2)
		return positions
	},
	int(Bar): func(start int) []*Position {
		positions := initializePositions(3)
		positions[0].setX(start)
		positions[0].setY(0)
		positions[1].setX(start)
		positions[1].setY(1)
		positions[2].setX(start)
		positions[2].setY(2)
		return positions
	},
}

func initializePositions(size int) []*Position {
	positions := make([]*Position, 4)
	for h := range positions {
		positions[h] = &Position{}
	}
	return positions
}
