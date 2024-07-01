package game

import (
	"main/network"
	"main/structures"
)

type Player struct {
	Id         int32
	MaxHealth  int32
	Health     int32
	Type       int32
	X          float32
	Y          float32
	LastX      float32
	LastY      float32
	flags      int32
	Dead       bool
	Connection *Connection
}

func NewPlayer(connection *Connection, id int32, x float32, y float32) *Player {
	return &Player{
		Id:         id,
		MaxHealth:  100,
		Health:     100,
		Type:       0x030e,
		X:          x,
		Y:          y,
		LastX:      x,
		LastY:      y,
		flags:      structures.NoFlags,
		Dead:       false,
		Connection: connection,
	}
}

func (player *Player) NewObjectData() network.NewObjectData {
	newObjectData := network.NewObjectData{
		ObjectType: player.Type,
		StatusData: network.StatusData{
			ObjectId: player.Id,
			X:        player.X,
			Y:        player.Y,
			Stats:    []network.StatData{},
		},
	}
	return newObjectData
}

func (player *Player) HasFlag(flag int32) bool {
	return (player.flags & flag) != 0
}

func (player *Player) SetPosition(x float32, y float32) {
	player.LastX = player.X
	player.LastY = player.Y
	player.X = x
	player.Y = y
	player.flags |= structures.MovedFlag
}

func (player *Player) Update(dt float64) bool {
	if player.Dead {
		return false
	}

	// logic

	return !player.Dead
}

func (player *Player) TakeDamage(damage int32) {
	player.Health -= damage
	if player.Health <= 0 {
		player.Kill()
	}
}

func (player *Player) Kill() {
	player.Dead = true
}

func (player *Player) ClearFlags() {
	player.flags = structures.NoFlags
}
