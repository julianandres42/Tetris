package shape

import "testing"

func TestShapeFactory_getSquare(t *testing.T) {
	square := shapeFct[int(Cube)]()
	square.InitializeShape(1)
	if square.ShapeType != Cube {
		t.Errorf("Bad shape type, got %d, expected %d", square.ShapeType, Cube)
	}
}

func TestShapeFactory_getEle(t *testing.T) {
	square := shapeFct[int(Ele)]()
	square.InitializeShape(1)
	if square.ShapeType != Ele {
		t.Errorf("Bad shape type, got %d, expected %d", square.ShapeType, Ele)
	}
}

func TestShapeFactory_getCube(t *testing.T) {
	square := shapeFct[int(Bar)]()
	square.InitializeShape(1)
	if square.ShapeType != Bar {
		t.Errorf("Bad shape type, got %d, expected %d", square.ShapeType, Bar)
	}
}
