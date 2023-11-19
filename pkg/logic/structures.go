package logic

type Coordinates struct {
	X, Y, Z float32
}

type Data struct {
	BedTemp          float32
	ExtruderTemp     float32
	PrintSpeed       int32
	FanSpeed         int32
	VolumetricFlow   float32
	HeadCoordinates  Coordinates
	ZOffset          float32
	PrintingFileName string
}

func NewData() Data {
	return Data{HeadCoordinates: Coordinates{}}
}
