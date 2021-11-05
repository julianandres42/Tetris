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
	game.Board.Init(20, 11)
}

func (game *Game) OpenWindow() {
	game.Drawer.OpenWindow()
}

func (game *Game) DrawBoard() {
	game.Drawer.DrawBoard(game.Board, colornames.Grey)
}

func (game *Game) Play() {
	game.CurrentShape = shape.ShapeFct[int(shape.Bar)]()
	game.CurrentShape.InitializeShape(5)
	game.Drawer.DrawShape(game.Board,game.CurrentShape,colornames.Yellow)
}

func (game *Game) UpdateScreen(){
	game.Drawer.UpdateScreen()
}

func (game *Game) UpdateScreenNoClose(){
	game.Drawer.UpdateScreenNoClose()
}
