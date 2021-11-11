package board

import (
	"errors"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type Drawer struct {
	ImdDrawer *imdraw.IMDraw
	Cfg       pixelgl.WindowConfig
	Window    *pixelgl.Window
}

func (drawer *Drawer) Init() {
	drawer.ImdDrawer = imdraw.New(nil)
	drawer.Cfg = pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, 638, 894),
		VSync:  true,
	}
}

func (drawer *Drawer) OpenWindow() {
	err := errors.New("")
	drawer.Window, err = pixelgl.NewWindow(drawer.Cfg)
	if err != nil {
		panic(err)
	}
}

func (drawer *Drawer) DrawBoard(board *Board, color color.Color) {
	drawer.ImdDrawer.Color = color
	for i, element := range board.Matrix {
		for j := range element {
			drawer.DrawSquare(float64(board.Matrix[i][j].Square.X), float64(board.Matrix[i][j].Square.Y))
		}
	}
}

func (drawer *Drawer) DrawShape(board *Board, shape *Shape, color color.Color) {
	drawer.ImdDrawer.Color = color
	board.UpdateShape(shape, true)
	for _, element := range shape.Positions {
		square := board.Matrix[element.GetY()][element.GetX()].Square
		drawer.DrawSquare(float64(square.X), float64(square.Y))
	}
}

func (drawer *Drawer) DrawSquare(x, y float64) {
	drawer.ImdDrawer.Push(pixel.V(x, y), pixel.V(x+40, y))
	drawer.ImdDrawer.Push(pixel.V(x+40, y), pixel.V(x+40, y-40))
	drawer.ImdDrawer.Push(pixel.V(x+40, y-40), pixel.V(x, y-40))
	drawer.ImdDrawer.Push(pixel.V(x, y-40), pixel.V(x, y))
	drawer.ImdDrawer.Line(5)
}

func (drawer *Drawer) UpdateScreen() {
	drawer.Window.Clear(colornames.Aliceblue)
	drawer.ImdDrawer.Draw(drawer.Window)
	drawer.Window.Update()
}

func (drawer *Drawer) UpdateScreenNoClose() {
	for !drawer.Window.Closed() {
		drawer.ImdDrawer.Draw(drawer.Window)
		drawer.Window.Update()
	}
}
