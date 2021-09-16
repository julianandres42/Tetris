package game

import "testing"

func TestInitCube(t *testing.T) {
	shape := mapperInitialization[int(Cube)](5)
	if shape[0].getX() != 5 || shape[0].getY() != 0 ||
		shape[1].getX() != 5 || shape[1].getY() != 1 ||
		shape[2].getX() != 6 || shape[2].getY() != 0 ||
		shape[3].getX() != 6 || shape[3].getY() != 1 {
		t.Errorf("Bad position for cube")
	}
}

func TestInitBar(t *testing.T) {
	shape := mapperInitialization[int(Bar)](5)
	if shape[0].getX() != 5 || shape[0].getY() != 0 ||
		shape[1].getX() != 5 || shape[1].getY() != 1 ||
		shape[2].getX() != 5 || shape[2].getY() != 2 {
		t.Errorf("Bad position")
	}
}

func TestInitEle(t *testing.T) {
	shape := mapperInitialization[int(Ele)](5)
	if shape[0].getX() != 5 || shape[0].getY() != 0 ||
		shape[1].getX() != 5 || shape[1].getY() != 1 ||
		shape[2].getX() != 6 || shape[2].getY() != 1 {
		t.Errorf("Bad position")
	}
}


