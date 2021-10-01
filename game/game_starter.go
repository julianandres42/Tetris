package game

import (
	board2 "Tetris/board"
	"Tetris/shape"
)

type Game struct {
	CurrentShape *shape.Shape
	Board        board2.Board
	Drawer       board2.Drawer
}
