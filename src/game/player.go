package game

type Player struct {
	Id         int32
	MaxHealth  int32
	Health     int32
	Type       int32
	X          float32
	Y          float32
	LastX      float32
	LastY      float32
	Flags      int32
	dead       bool
	Connection *Connection
}

func NewPlayer(connection *Connection, id int32, x float32, y float32) *Player {
	return &Player{
		Id:         id,
		MaxHealth:  0,
		Health:     0,
		Type:       0,
		X:          x,
		Y:          y,
		LastX:      x,
		LastY:      y,
		Flags:      0,
		dead:       false,
		Connection: connection,
	}
}

func (player *Player) Update(dt float64) bool {
	return !player.dead
}
