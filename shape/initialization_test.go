package shape

import "testing"

func TestInitCube(t *testing.T) {
	shape := mapperInitialization[int(Cube)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 6 || shape[2].GetY() != 0 ||
		shape[3].GetX() != 6 || shape[3].GetY() != 1 {
		t.Errorf("Bad position for cube")
	}
}

func TestInitBar(t *testing.T) {
	shape := mapperInitialization[int(Bar)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 5 || shape[2].GetY() != 2 {
		t.Errorf("Bad position")
	}
}

func TestInitEle(t *testing.T) {
	shape := mapperInitialization[int(Ele)](5)
	if shape[0].GetX() != 5 || shape[0].GetY() != 0 ||
		shape[1].GetX() != 5 || shape[1].GetY() != 1 ||
		shape[2].GetX() != 6 || shape[2].GetY() != 1 {
		t.Errorf("Bad position")
	}
}
