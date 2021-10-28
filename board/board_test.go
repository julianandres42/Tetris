package board

import (
	"fmt"
	"testing"
)

func TestBoard_Init(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	if len(board.Matrix) != 20 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", len(board.Matrix), 20))
	}
	if board.Matrix[0][0].Square.X != 119 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[0][0].Square.X, 119))
	}
	if board.Matrix[0][0].Square.Y != 487 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[0][0].Square.Y, 487))
	}
	if board.Matrix[19][9].Square.X != 479 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[19][9].Square.X, 479))
	}
	if board.Matrix[19][9].Square.Y != -273 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[19][9].Square.Y, -273))
	}
}

func TestBoard_EvaluateLines(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	setLine(board.Matrix[5])
	setLine(board.Matrix[10])
	setLine(board.Matrix[15])
	board.EvaluateLines()
	if isLine(board.Matrix[5]) {
		t.Errorf(fmt.Sprintf("line found in position %d", 5))
	}
	if isLine(board.Matrix[10]) {
		t.Errorf(fmt.Sprintf("line found in position %d", 10))
	}
	if isLine(board.Matrix[15]) {
		t.Errorf(fmt.Sprintf("line found in position %d", 15))
	}
}

func TestBoard_UpdateShape(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	//shape := shape2.Shape{}
}

func setLine(cells []*Cell) {
	for _, cell := range cells {
		cell.active = true
	}
}
