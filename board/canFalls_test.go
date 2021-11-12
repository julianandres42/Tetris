package board

import (
	"testing"
)

func TestShape_CanFallSquare(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	square := ShapeFct[int(Cube)]()
	square.InitializeShape(5)
	funcFall := mapperFalls[int(Cube)]
	if !funcFall(square.Positions, board.Matrix) {
		t.Errorf("Bad position")
	}
}

func TestShape_CanNotFallOneCellDown(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	square := ShapeFct[int(Cube)]()
	square.InitializeShape(5)
	board.Matrix[square.Positions[2].GetY()+1][square.Positions[2].GetX()].Active = true
	funcFall := mapperFalls[int(Cube)]
	if !funcFall(square.Positions, board.Matrix) {
		t.Errorf("Bad position")
	}
}
