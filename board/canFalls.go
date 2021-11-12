package board

type canFallFunctionMapper map[int]func([]*Position, [][]*Cell) bool

var mapperFalls = canFallFunctionMapper{
	int(Cube): func(positions []*Position, matrix [][]*Cell) bool {
		return (!matrix[positions[2].GetY()+1][positions[2].GetX()].Active && !matrix[positions[3].GetY()+1][positions[3].GetX()].Active) ||
			(positions[2].GetY()+1 <= len(matrix)-1 && positions[3].GetY()+1 <= len(matrix)-1)
	},
}
