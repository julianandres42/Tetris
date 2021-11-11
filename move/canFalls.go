package move

import (
	"Tetris/board"
)

type canFallFunctionMapper map[int]func([]*board.Position, [][]*board.Cell) bool

var mapperFalls = canFallFunctionMapper{
	int(board.Cube): func(positions []*board.Position, matrix [][]*board.Cell) bool {
		return (!matrix[positions[2].GetY()+1][positions[2].GetY()].Active && !matrix[positions[3].GetY()+1][positions[3].GetX()].Active) ||
			(positions[2].GetY()+1 <= len(matrix)-1 && positions[3].GetY()+1 <= len(matrix)-1)
	},
}
