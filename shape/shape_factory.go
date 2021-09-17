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
		rotateFunction: rotateSquare,
	}
}

func getEle() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Ele)],
		rotateFunction: rotateEle,
	}
}

func getBar() *Shape {
	return &Shape{
		initialize:     mapperInitialization[int(Bar)],
		rotateFunction: rotateBar,
	}
}
