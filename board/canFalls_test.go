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
	if !funcFall(square.Positions, board.Matrix, 1) {
		t.Errorf("Bad position")
	}
}

func TestShape_CanNotFallOneCellDown(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	square := ShapeFct[int(Cube)]()
	square.InitializeShape(5)
	board.Matrix[square.Positions[1].GetY()+1][square.Positions[1].GetX()].Active = true
	funcFall := mapperFalls[int(Cube)]
	if funcFall(square.Positions, board.Matrix, 1) {
		t.Errorf("Bad position")
	}
	board.Matrix[square.Positions[1].GetY()+1][square.Positions[1].GetX()].Active = false
	board.Matrix[square.Positions[3].GetY()+1][square.Positions[3].GetX()].Active = true
	if funcFall(square.Positions, board.Matrix, 1) {
		t.Errorf("Bad position")
	}
}

func TestCanNotFallBoardBorder(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	square := ShapeFct[int(Cube)]()
	square.InitializeShape(5)
	for _, pos := range square.Positions {
		pos.Y += 18
	}
	funcFall := mapperFalls[int(Cube)]
	board.UpdateShape(square, true)
	if funcFall(square.Positions, board.Matrix, 1) {
		t.Errorf("Bad position")
	}
}
