package main

import (
	 game "Tetris/game"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	game1 := &game.Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tetris")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game1); err != nil {
		log.Fatal(err)
	}
}