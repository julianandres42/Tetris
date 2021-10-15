package shape

import "testing"

func TestShapeFactory_getSquare(t *testing.T) {
	square := shapeFct[int(Cube)]()
	if square.ShapeType != Cube {
		t.Errorf("Bad shape type, got %d, expected %d", square.ShapeType, Cube)
	}
}
