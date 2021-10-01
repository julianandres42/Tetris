package board

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Drawer struct {
	ImdDrawer *imdraw.IMDraw
}

func (drawer *Drawer) DrawSquare(x, y float64) {
	drawer.ImdDrawer.Push(pixel.V(x, y), pixel.V(x+40, y))
	drawer.ImdDrawer.Push(pixel.V(x+40, y), pixel.V(x+40, y-40))
	drawer.ImdDrawer.Push(pixel.V(x+40, y-40), pixel.V(x, y-40))
	drawer.ImdDrawer.Push(pixel.V(x, y-40), pixel.V(x, y))
	drawer.ImdDrawer.Line(5)
}
