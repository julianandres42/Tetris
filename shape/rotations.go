package shape

type rotationMapper map[int]func(positions []*Position) []*Position

var mapperEle = rotationMapper{
	2: func(positions []*Position) []*Position {
		positions[1].DescX()
		positions[1].AddY()
		positions[2].DescX()
		positions[2].DescY()
		return positions
	},
	3: func(positions []*Position) []*Position {
		positions[1].AddX()
		positions[1].AddY()
		positions[2].DescX()
		positions[2].AddY()
		return positions
	},
	4: func(positions []*Position) []*Position {
		positions[1].AddX()
		positions[1].DescY()
		positions[2].AddX()
		positions[2].AddY()
		return positions
	},
	1: func(positions []*Position) []*Position {
		positions[1].DescX()
		positions[1].DescY()
		positions[2].AddX()
		positions[2].DescY()
		return positions
	},
}

var mapperBar = rotationMapper{
	2: func(positions []*Position) []*Position {
		positions[1].DescX()
		positions[1].AddY()
		positions[2].AddX()
		positions[2].DescY()
		return positions
	},
	1: func(positions []*Position) []*Position {
		positions[1].AddX()
		positions[1].DescY()
		positions[2].DescX()
		positions[2].AddY()
		return positions
	},
}

func rotateEle(positions []*Position, rotation int) []*Position {
	return mapperEle[rotation](positions)
}

func rotateBar(positions []*Position, rotation int) []*Position {
	return mapperBar[rotation](positions)
}

func rotateSquare(positions []*Position, rotation int) []*Position {
	return positions
}
