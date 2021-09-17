package game

import (
	"Tetris/shape"
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	currentShape *shape.Shape
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
