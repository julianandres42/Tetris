package shape

type Position struct {
	x int
	y int
}

func (position *Position) SetX(x int) {
	position.x = x
}

func (position *Position) SetY(y int) {
	position.y = y
}

func (position *Position) GetX() int {
	return position.x
}

func (position *Position) GetY() int {
	return position.y
}

func (position *Position) AddY() {
	position.y++
}

func (position *Position) AddX() {
	position.x++
}

func (position *Position) DescX() {
	position.x--
}

func (position *Position) DescY() {
	position.y--
}

type ShapeType int

const (
	Cube ShapeType = 1
	Ele            = 2
	Bar            = 3
)

type Shape struct {
	Positions       []*Position
	CurrentRotation int
	MaxRotations    int
	ShapeType       ShapeType
	RotateFunction  func([]*Position, int) []*Position
	initialize      func(int) []*Position
}

func (shape *Shape) rotate() {
	shape.Positions = shape.RotateFunction(shape.Positions, shape.CurrentRotation)
	if shape.CurrentRotation == shape.MaxRotations {
		shape.CurrentRotation = 1
	} else {
		shape.CurrentRotation++
	}
}

func (shape *Shape) initializeShape(start int) {
	shape.Positions = shape.initialize(start)

}

func (shape *Shape) fall() {
	for i := 0; i < len(shape.Positions); i++ {
		shape.Positions[i].AddY()
	}
}

func (shape *Shape) moveRight() {
	for i := 0; i < len(shape.Positions); i++ {
		shape.Positions[i].AddX()
	}
}

func (shape *Shape) moveLeft() {
	for i := 0; i < len(shape.Positions); i++ {
		shape.Positions[i].DescX()
	}
}
