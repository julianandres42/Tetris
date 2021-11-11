package board

import "testing"

func TestInitCube(t *testing.T) {
	shape, maxRotations, shapeType := mapperInitialization[int(Cube)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 6 || shape[2].GetY() != 0 ||
		shape[3].GetX() != 6 || shape[3].GetY() != 1 {
		t.Errorf("Bad position for cube")
	}
	if maxRotations != 0 {
		t.Errorf("Bad max rotations, got %d, expected %d", maxRotations, 0)
	}
	if shapeType != Cube {
		t.Errorf("Bad shape type, got %d, expected %d", shapeType, Cube)
	}
}

func TestInitBar(t *testing.T) {
	shape, maxRotations, shapeType := mapperInitialization[int(Bar)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 5 || shape[2].GetY() != 2 {
		t.Errorf("Bad position")
	}
	if maxRotations != 2 {
		t.Errorf("Bad max rotations, got %d, expected %d", maxRotations, 1)
	}
	if shapeType != Bar {
		t.Errorf("Bad shape type, got %d, expected %d", shapeType, Bar)
	}
}

func TestInitEle(t *testing.T) {
	shape, maxRotations, shapeType := mapperInitialization[int(Ele)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 6 || shape[2].GetY() != 1 {
		t.Errorf("Bad position")
	}
	if maxRotations != 4 {
		t.Errorf("Bad max rotations, got %d, expected %d", maxRotations, 4)
	}
	if shapeType != Ele {
		t.Errorf("Bad shape type, got %d, expected %d", shapeType, Ele)
	}
}
