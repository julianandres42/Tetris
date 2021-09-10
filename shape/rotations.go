package shape

type rotationMapper map[int]func(positions []*Position) []*Position

var mapperEle = rotationMapper{
	2: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].addY()
		positions[2].descX()
		positions[2].descY()
		return positions
	},
	3: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].addY()
		positions[2].descX()
		positions[2].addY()
		return positions
	},
	4: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].descY()
		positions[2].addX()   
		positions[2].addY()
		return positions
	},
	1: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].descY()
		positions[2].addX()
		positions[2].descY()
		return positions
	},
}

var mapperBar = rotationMapper{
	2: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].addY()
		positions[2].addX()
		positions[2].descY()
		return positions
	},
	1: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].descY()
		positions[2].descX()
		positions[2].addY()
		return positions
	},
}

func rotateEle(positions []*Position, rotation int) []*Position {
	return mapperEle[rotation](positions)
}

func rotateBar(positions []*Position, rotation int) []*Position {
	return mapperBar[rotation](positions)
}
