package board

type Position struct {
	X int
	Y int
}

func (position *Position) SetX(x int) {
	position.X = x
}

func (position *Position) SetY(y int) {
	position.Y = y
}

func (position *Position) GetX() int {
	return position.X
}

func (position *Position) GetY() int {
	return position.Y
}

func (position *Position) AddY() {
	position.Y++
}

func (position *Position) AddX() {
	position.X++
}

func (position *Position) DescX() {
	position.X--
}

func (position *Position) DescY() {
	position.Y--
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
	CanFallFunction func([]*Position, [][]*Cell) bool
	initialize      func(int) ([]*Position, int, ShapeType)
}

func (shape *Shape) InitializeShape(start int) {
	shape.Positions, shape.MaxRotations, shape.ShapeType = shape.initialize(start)
	shape.CurrentRotation = 1
}

func (shape *Shape) CanFall(matrix [][]*Cell) bool {
	return shape.CanFallFunction(shape.Positions, matrix)
}

func (shape *Shape) Rotate() {
	shape.Positions = shape.RotateFunction(shape.Positions, shape.CurrentRotation)
	if shape.CurrentRotation == shape.MaxRotations {
		shape.CurrentRotation = 1
	} else {
		shape.CurrentRotation++
	}
}

func (shape *Shape) Fall() {
	for i := 0; i < len(shape.Positions); i++ {
		shape.Positions[i].AddY()
	}
}

func (shape *Shape) MoveRight() {
	for i := 0; i < len(shape.Positions); i++ {
		shape.Positions[i].AddX()
	}
}

func (shape *Shape) MoveLeft() {
	for i := 0; i < len(shape.Positions); i++ {
		shape.Positions[i].DescX()
	}
}
