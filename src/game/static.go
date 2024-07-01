package game

type Static struct {
	Id   int32
	Type int32
	X    float32
	Y    float32
	Dead bool
}

func (static *Static) SetPosition(x float32, y float32) {
	static.X = x
	static.Y = y
}
