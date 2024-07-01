package game

import (
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
	Flags      int32
	Dead       bool
	Connection *Connection
}

func NewPlayer(connection *Connection, id int32, x float32, y float32) *Player {
	return &Player{
		Id:         id,
		MaxHealth:  0,
		Health:     0,
		Type:       0x030e,
		X:          x,
		Y:          y,
		LastX:      x,
		LastY:      y,
		Flags:      structures.NoFlags,
		Dead:       false,
		Connection: connection,
	}
}

func (enemy *Player) SetPosition(x float32, y float32) {
	enemy.LastX = enemy.X
	enemy.LastY = enemy.Y
	enemy.X = x
	enemy.Y = y
	enemy.Flags |= structures.MovedFlag
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
	player.Flags = structures.NoFlags
}
