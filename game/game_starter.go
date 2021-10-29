package game

import (
	board2 "Tetris/board"
	"Tetris/shape"
	"golang.org/x/image/colornames"
	"math/rand"
)

type Game struct {
	CurrentShape *shape.Shape
	Board        *board2.Board
	Drawer       *board2.Drawer
}

func (game *Game) Init() {
	game.Drawer.Init()
	game.Board.Init(10, 20)
}

func (game *Game) OpenWindow() {
	game.Drawer.OpenWindow()
}

func (game *Game) DrawBoard() {
	game.Drawer.DrawBoard(game.Board, colornames.Grey)
}

func (game *Game) Play() {
	game.CurrentShape = shape.ShapeFct[rand.Intn(4)]()
	game.Board.UpdateShape(game.CurrentShape, true)
}
