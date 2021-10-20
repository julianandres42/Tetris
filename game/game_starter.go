package game

import (
	board2 "Tetris/board"
	"Tetris/shape"
	"golang.org/x/image/colornames"
)

type Game struct {
	CurrentShape *shape.Shape
	Board        *board2.Board
	Drawer       *board2.Drawer
}

func (game *Game) Init() {
	game.Drawer.Init()
	game.Board.Init(10, 20)
	game.Drawer.DrawBoard(game.Board, colornames.Grey)
}
