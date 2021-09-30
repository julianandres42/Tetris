package board

import (
	"fmt"
	"testing"
)

func TestInitBoard(t *testing.T) {
	board := &Board{}
	board.Init(20, 10)
	if len(board.Matrix) != 20 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", len(board.Matrix), 20))
	}
	if board.Matrix[0][0].Square.X != 119 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[0][0].Square.X, 119))
	}
	if board.Matrix[0][0].Square.Y != 847 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[0][0].Square.Y, 847))
	}
	if board.Matrix[19][9].Square.X != 479 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[19][9].Square.X, 479))
	}
	if board.Matrix[19][9].Square.Y != 87 {
		t.Errorf(fmt.Sprintf("Got %d, expected %d", board.Matrix[19][9].Square.Y, 87))
	}
}
