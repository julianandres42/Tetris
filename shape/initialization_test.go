package shape

import "testing"

func TestInitCube(t *testing.T) {
	shape, maxRotations := mapperInitialization[int(Cube)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 6 || shape[2].GetY() != 0 ||
		shape[3].GetX() != 6 || shape[3].GetY() != 1 {
		t.Errorf("Bad position for cube")
	}
	if maxRotations != 0 {
		t.Errorf("Bad max rotations, got %d, expected %d", maxRotations, 0)
	}
}

func TestInitBar(t *testing.T) {
	shape, maxRotations := mapperInitialization[int(Bar)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 5 || shape[2].GetY() != 2 {
		t.Errorf("Bad position")
	}
	if maxRotations != 1 {
		t.Errorf("Bad max rotations, got %d, expected %d", maxRotations, 1)
	}
}

func TestInitEle(t *testing.T) {
	shape, maxRotations := mapperInitialization[int(Ele)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 6 || shape[2].GetY() != 1 {
		t.Errorf("Bad position")
	}
	if maxRotations != 2 {
		t.Errorf("Bad max rotations, got %d, expected %d", maxRotations, 1)
	}
}
