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
}
