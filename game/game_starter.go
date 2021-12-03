package game

import (
	board2 "Tetris/board"
	"golang.org/x/image/colornames"
	"image/color"
)

type Game struct {
	CurrentShape *board2.Shape
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
	game.CurrentShape = board2.ShapeFct[int(board2.Cube)]()
	game.CurrentShape.InitializeShape(5)
	game.Drawer.DrawShape(game.Board, game.CurrentShape, colornames.Yellow)
	go playing(game)
}

func (game *Game) UpdateScreen() {
	game.Drawer.UpdateScreen()
}

func (game *Game) UpdateScreenNoClose() {
	game.Drawer.UpdateScreenNoClose()
}

func playing(game *Game) {
	for {
		if !game.CurrentShape.CanFall(game.Board.Matrix) {
			break
		} else {
			previousPositions := clonePositions(game.CurrentShape.Positions)
			game.CurrentShape.Fall()
			actualPositions := game.CurrentShape.Positions
			updateShape(game, previousPositions, false, colornames.Grey)
			updateShape(game, actualPositions, true, colornames.Yellow)
			game.UpdateScreenNoClose()
		}
	}
}

func updateShape(game *Game, positions []*board2.Position, truth bool, color color.Color) {
	game.CurrentShape.Positions = positions
	game.Board.UpdateShape(game.CurrentShape, truth)
	game.Drawer.DrawShape(game.Board, game.CurrentShape, color)
}

func clonePositions(positions []*board2.Position) []*board2.Position {
	newPositions := make([]*board2.Position, len(positions))
	for index, element := range positions {
		newPositions[index] = &board2.Position{
			X: element.X,
			Y: element.Y,
		}
	}
	return newPositions
}
