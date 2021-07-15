package game

type rotationMapper map[int]func(positions []*Position) []*Position

var mapperEle = rotationMapper{
	1: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].addY()
		positions[2].descX()
		positions[2].descY()
		return positions
	},
	2: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].addY()
		positions[2].descX()
		positions[2].addY()
		return positions
	},
	3: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].descY()
		positions[2].addX()   
		positions[2].addY()
		return positions
	},
	4: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].descY()
		positions[2].addX()
		positions[2].descY()
		return positions
	},
}

var mapperBar = rotationMapper{
	1: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].addY()
		positions[2].descX()
		positions[2].descY()
		return positions
	},
	2: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].addY()
		positions[2].descX()
		positions[2].addY()
		return positions
	},
	3: func(positions []*Position) []*Position {
		positions[1].addX()
		positions[1].descY()
		positions[2].addX()   
		positions[2].addY()
		return positions
	},
	4: func(positions []*Position) []*Position {
		positions[1].descX()
		positions[1].descY()
		positions[2].addX()
		positions[2].descY()
		return positions
	},
}

func rotateEle(positions []*Position, rotation int) ([]*Position, int) {
	return mapperEle[rotation](positions), rotation
}
