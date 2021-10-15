package shape

import "testing"

func TestShape_InitializeShape(t *testing.T) {
	shape := shapeFct[int(Ele)]()
	shape.InitializeShape(5)
	if len(shape.Positions) != 3 {
		t.Errorf("Bad size, got %d, expected %d", len(shape.Positions), 3)
	}
}

func TestShape_Rotate(t *testing.T) {
	shape := shapeFct[int(Ele)]()
	shape.InitializeShape(5)
	if shape.CurrentRotation != 1 {
		t.Errorf("Bad rotation, got %d, expected %d", shape.CurrentRotation, 1)
	}
	shape.Rotate()
	if shape.CurrentRotation != 2 {
		t.Errorf("Bad rotation, got %d, expected %d", shape.CurrentRotation, 2)
	}
	shape.Rotate()
	if shape.CurrentRotation != 3 {
		t.Errorf("Bad rotation, got %d, expected %d", shape.CurrentRotation, 3)
	}
	shape.Rotate()
	if shape.CurrentRotation != shape.MaxRotations {
		t.Errorf("Bad rotation, got %d, expected %d", shape.CurrentRotation, shape.CurrentRotation)
	}

}

func TestShape_Fall(t *testing.T) {
	shape := shapeFct[int(Cube)]()
	shape.InitializeShape(5)
	yInit := shape.Positions[0].GetY()
	shape.Fall()
	if shape.Positions[0].GetY() != yInit+1 {
		t.Errorf("Bad position, got %d, expected %d", shape.Positions[0].GetY(), yInit+1)
	}
}

func TestShape_MoveLeft(t *testing.T) {
	shape := shapeFct[int(Cube)]()
	shape.InitializeShape(5)
	xInit := shape.Positions[0].GetX()
	shape.MoveLeft()
	if shape.Positions[0].GetX() != xInit-1 {
		t.Errorf("Bad position, got %d, expected %d", shape.Positions[0].GetY(), xInit-1)
	}
}

func TestShape_MoveRight(t *testing.T) {
	shape := shapeFct[int(Cube)]()
	shape.InitializeShape(5)
	xInit := shape.Positions[0].GetX()
	shape.MoveRight()
	if shape.Positions[0].GetX() != xInit+1 {
		t.Errorf("Bad position, got %d, expected %d", shape.Positions[0].GetY(), xInit-1)
	}
}
