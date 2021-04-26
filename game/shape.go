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

type ShapeType int

const (
	Cube ShapeType = iota
	Ele
	Te
	Bar
)

type Shape struct {
	positions       []*Position
	currentRotation int
	maxRotations    int
	shapeType       ShapeType
	rotate          func(positions []*Position) []*Position
}

func (shape *Shape) rotateShape() {
	shape.positions = shape.rotate(shape.positions)
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
