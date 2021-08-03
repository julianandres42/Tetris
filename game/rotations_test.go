package game

import "testing"

func TestRotateEle(t *testing.T) {
	positions := make([]*Position, 0)
	positions = append(positions,
		&Position{x: 5, y: 1},
		&Position{x: 5, y: 0},
		&Position{x: 6, y: 1})
	rotation := 2
	positions = rotateEle(positions, rotation)
	if positions[0].getX() != 5 && positions[0].getY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].getX()+1 != 5 && positions[1].getY()-1 != 0 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].getX()+1 != 6 && positions[2].getY()+1 != 1 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation += 1
	positions = rotateEle(positions, rotation)
	if positions[0].getX() != 5 && positions[0].getY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].getX()-1 != 4 && positions[1].getY()-1 != 1 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].getX()+1 != 5 && positions[2].getY()-1 != 0 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation += 1
	positions = rotateEle(positions, rotation)
	if positions[0].getX() != 5 && positions[0].getY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].getX()-1 != 4 && positions[1].getY()+1 != 2 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].getX()-1 != 4 && positions[2].getY()-1 != 1 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation = 1
	positions = rotateEle(positions, rotation)
	if positions[0].getX() != 5 && positions[0].getY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].getX()+1 != 6 && positions[1].getY()+1 != 1 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].getX()+1 != 7 && positions[2].getY()-1 != 0 {
		t.Errorf("Bad coordinates %d", 2)
	}

}

func TestRotateBar(t *testing.T) {

}
