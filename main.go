package main

import (
	board2 "Tetris/board"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, 638, 894),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	board := &board2.Board{}
	board.Init(20, 10)

	imd := imdraw.New(nil)
	drawer := &board2.Drawer{ImdDrawer: imd}
	imd.Color = colornames.Grey
	for i, element := range board.Matrix {
		for j := range element {
			drawer.DrawSquare(float64(board.Matrix[i][j].Square.X), float64(board.Matrix[i][j].Square.Y))
		}
	}

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
