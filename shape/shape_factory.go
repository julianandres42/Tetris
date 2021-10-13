package shape

type shapeFactory map[int]func() *Shape

var shapeFct = shapeFactory{
	int(Cube): getSquare,
	int(Ele):  getEle,
	int(Bar):  getBar,
}

func getSquare() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Cube)],
		RotateFunction: rotateSquare,
	}
}

func getEle() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Ele)],
		RotateFunction: rotateEle,
	}
}

func getBar() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Bar)],
		RotateFunction: rotateBar,
	}
}
