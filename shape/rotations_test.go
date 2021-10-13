package shape

import "testing"

func TestRotateEle(t *testing.T) {
	positions := make([]*Position, 0)
	positions = append(positions,
		&Position{x: 5, y: 1},
		&Position{x: 5, y: 0},
		&Position{x: 6, y: 1})
	rotation := 2
	positions = rotateEle(positions, rotation)
	if positions[0].GetX() != 5 && positions[0].GetY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].GetX()+1 != 5 && positions[1].GetY()-1 != 0 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].GetX()+1 != 6 && positions[2].GetY()+1 != 1 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation += 1
	positions = rotateEle(positions, rotation)
	if positions[0].GetX() != 5 && positions[0].GetY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].GetX()-1 != 4 && positions[1].GetY()-1 != 1 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].GetX()+1 != 5 && positions[2].GetY()-1 != 0 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation += 1
	positions = rotateEle(positions, rotation)
	if positions[0].GetX() != 5 && positions[0].GetY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].GetX()-1 != 4 && positions[1].GetY()+1 != 2 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].GetX()-1 != 4 && positions[2].GetY()-1 != 1 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation = 1
	positions = rotateEle(positions, rotation)
	if positions[0].GetX() != 5 && positions[0].GetY() != 1 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].GetX()+1 != 6 && positions[1].GetY()+1 != 1 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].GetX()+1 != 7 && positions[2].GetY()-1 != 0 {
		t.Errorf("Bad coordinates %d", 2)
	}

}

func TestRotateBar(t *testing.T) {
	positions := make([]*Position, 0)
	positions = append(positions,
		&Position{x: 5, y: 2},
		&Position{x: 5, y: 1},
		&Position{x: 5, y: 3})
	rotation := 2
	positions = rotateBar(positions, rotation)
	if positions[0].GetX() != 5 && positions[0].GetY() != 2 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].GetX()+1 != 5 && positions[1].GetY()-1 != 1 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].GetX()-1 != 7 && positions[2].GetY()+1 != 3 {
		t.Errorf("Bad coordinates %d", 2)
	}
	rotation = 1
	positions = rotateBar(positions, rotation)
	if positions[0].GetX() != 5 && positions[0].GetY() != 2 {
		t.Errorf("Bad coordinates %d", 0)
	}
	if positions[1].GetX()-1 != 4 && positions[1].GetY()+1 != 2 {
		t.Errorf("Bad coordinates %d", 1)
	}
	if positions[2].GetX()+1 != 6 && positions[2].GetY()-1 != 2 {
		t.Errorf("Bad coordinates %d", 1)
	}

}
