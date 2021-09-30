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
	imd.Color = colornames.Grey
	imd.EndShape = imdraw.RoundEndShape
	for i, element := range board.Matrix {
		for j := range element {
			imd.Push(pixel.V(float64(board.Matrix[i][j].Square.X), float64(board.Matrix[i][j].Square.Y)),
				pixel.V(float64(board.Matrix[i][j].Square.X)+40, float64(board.Matrix[i][j].Square.Y)))
			imd.Push(pixel.V(float64(board.Matrix[i][j].Square.X)+40, float64(board.Matrix[i][j].Square.Y)),
				pixel.V(float64(board.Matrix[i][j].Square.X)+40, float64(board.Matrix[i][j].Square.Y)-40))
			imd.Push(pixel.V(float64(board.Matrix[i][j].Square.X)+40, float64(board.Matrix[i][j].Square.Y)-40),
				pixel.V(float64(board.Matrix[i][j].Square.X), float64(board.Matrix[i][j].Square.Y)-40))
			imd.Push(pixel.V(float64(board.Matrix[i][j].Square.X), float64(board.Matrix[i][j].Square.Y)-40),
				pixel.V(float64(board.Matrix[i][j].Square.X), float64(board.Matrix[i][j].Square.Y)))
			imd.Line(5)
		}
	}

	//imd.Push(pixel.V(119, 47), pixel.V(519, 47))
	//imd.Push(pixel.V(519, 47), pixel.V(519, 847))
	//imd.Push(pixel.V(519, 847), pixel.V(119, 847))
	//imd.Push(pixel.V(119, 847), pixel.V(119, 47))
	//imd.Line(5)
	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
