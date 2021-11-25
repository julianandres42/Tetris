package board

type shapeFactory map[int]func() *Shape

var ShapeFct = shapeFactory{
	int(Cube): GetSquare,
	int(Ele):  GetEle,
	int(Bar):  GetBar,
}

func GetSquare() *Shape {
	return &Shape{
		initialize:      mapperInitialization[int(Cube)],
		RotateFunction:  rotateSquare,
		CanFallFunction: mapperFalls[int(Cube)],
	}
}

func GetEle() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Ele)],
		RotateFunction: rotateEle,
	}
}

func GetBar() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Bar)],
		RotateFunction: rotateBar,
	}
}
