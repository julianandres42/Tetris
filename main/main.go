package main

import (
	board2 "Tetris/board"
	game2 "Tetris/game"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	game := game2.Game{Board: &board2.Board{},
		CurrentShape: &board2.Shape{},
		Drawer:       &board2.Drawer{},
	}
	game.Init()
	game.OpenWindow()
	game.DrawBoard()
	game.UpdateScreen()
	game.Play()
	game.UpdateScreenNoClose()
}

func main() {
	pixelgl.Run(run)
}
