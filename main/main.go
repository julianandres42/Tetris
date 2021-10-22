package main

import (
	board2 "Tetris/board"
	game2 "Tetris/game"
	"Tetris/shape"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	game := game2.Game{Board: &board2.Board{},
		CurrentShape: &shape.Shape{},
		Drawer:       &board2.Drawer{},
	}
	game.Init()
}

func main() {
	pixelgl.Run(run)
}
