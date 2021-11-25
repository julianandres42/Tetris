package board

type canFallFunctionMapper map[int]func([]*Position, [][]*Cell, int) bool

var mapperFalls = canFallFunctionMapper{
	int(Cube): func(positions []*Position, matrix [][]*Cell, currentRotation int) bool {
		return (positions[1].GetY()+1 <= len(matrix)-1) &&
			(!matrix[positions[1].GetY()+1][positions[1].GetX()].Active &&
				!matrix[positions[3].GetY()+1][positions[3].GetX()].Active)
	},
}
