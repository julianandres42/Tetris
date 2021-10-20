package game

import (
	board2 "Tetris/board"
	"testing"
)

func TestGame_Init(t *testing.T) {
	game := &Game{
		Board:  &board2.Board{},
		Drawer: &board2.Drawer{},
	}
 	game.Init()
	if len(game.Board.Matrix) < 0{
		t.Errorf("Matrix is empty")
	}
}
