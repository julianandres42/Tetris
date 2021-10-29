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
	if len(game.Board.Matrix) == 0 {
		t.Errorf("Bad matrix type, got %d", 0)
	}
	if game.Drawer.ImdDrawer == nil {
		t.Errorf("Bad value, got nill")
	}
}
