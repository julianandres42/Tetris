package game

import "testing"

func TestRotateEle(t *testing.T) {
	positions := make([]*Position, 0)
	positions = append(positions, &Position{x: 5, y: 0},
		&Position{x: 5, y: 1},
		&Position{x: 6, y: 1})
	rotation := 1
	positions , rotation = rotateEle(positions,rotation)

}
