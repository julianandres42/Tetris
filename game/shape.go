package game

type Position struct {
	x int
	y int
}

func (position *Position) setX(x int) {
	position.x = x
}

func (position *Position) setY(y int) {
	position.y = y
}

func (position *Position) getX() int {
	return position.x
}

func (position *Position) getY() int {
	return position.y
}

func (position *Position) addY() {
	position.y++
}

func (position *Position) addX() {
	position.x++
}

func (position *Position) descX() {
	position.x--
}

func (position *Position) descY() {
	position.y--
}

type ShapeType int

const (
	Cube ShapeType = 1
	Ele            = 2
	Bar             = 3
)

type Shape struct {
	positions       []*Position
	currentRotation int
	maxRotations    int
	shapeType       ShapeType
	rotateFunction  func([]*Position,  int) ([]*Position)
	initialize      func(int) []*Position
}

func (shape *Shape) rotate() {
	shape.positions = shape.rotateFunction(shape.positions, shape.currentRotation)
	if shape.currentRotation == shape.maxRotations {
		shape.currentRotation = 1
	} else {
		shape.currentRotation++
	}
}

func (shape *Shape) initializeShape(start int) {
	shape.positions = shape.initialize(start)

}

func (shape *Shape) fall() {
	for i := 0; i < len(shape.positions); i++ {
		shape.positions[i].addY()
	}
}

func (shape *Shape) moveRight() {
	for i := 0; i < len(shape.positions); i++ {
		shape.positions[i].addX()
	}
}

func (shape *Shape) moveLeft() {
	for i := 0; i < len(shape.positions); i++ {
		shape.positions[i].descX()
	}
}
