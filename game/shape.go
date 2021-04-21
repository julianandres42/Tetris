package game

type Position struct {
	x int
	y int
}

type Shape struct {
	positions       []*Position
	currentRotation int
	maxRotations    int
	rotate          func(positions []*Position) []*Position
}

func (shape *Shape) rotateShape(){
	shape.positions = shape.rotate(shape.positions)
}


