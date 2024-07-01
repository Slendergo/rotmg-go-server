package game

import (
	assets "main/assets/xml"
	"main/network"
	"main/structures"
	"math"
)

type Enemy struct {
	Id        int32
	MaxHealth int32
	Health    int32
	Size      float32
	Type      int32
	X         float32
	Y         float32
	LastX     float32
	LastY     float32
	flags     int32
	Dead      bool
}

func NewEnemy(xmlEnemy *assets.XMLEnemy, id int32, x float32, y float32) *Enemy {
	return &Enemy{
		Id:        id,
		MaxHealth: xmlEnemy.MaxHitPoints,
		Health:    xmlEnemy.MaxHitPoints,
		Size:      float32(xmlEnemy.Size),
		Type:      xmlEnemy.Type,
		X:         x,
		Y:         y,
		LastX:     x,
		LastY:     y,
		flags:     structures.NoFlags,
		Dead:      false,
	}
}

func (enemy *Enemy) NewObjectData() network.NewObjectData {
	newObjectData := network.NewObjectData{
		ObjectType: enemy.Type,
		StatusData: network.StatusData{
			ObjectId: enemy.Id,
			X:        enemy.X,
			Y:        enemy.Y,
			Stats:    []network.StatData{},
		},
	}
	return newObjectData
}

func (enemy *Enemy) HasFlag(flag int32) bool {
	return (enemy.flags & flag) != 0
}

func (enemy *Enemy) Facing() float32 {
	dx := enemy.X - enemy.LastX
	dy := enemy.Y - enemy.LastY
	return float32(math.Atan2(float64(dy), float64(dx)))
}

func (enemy *Enemy) SetPosition(x float32, y float32) {
	enemy.LastX = enemy.X
	enemy.LastY = enemy.Y
	enemy.X = x
	enemy.Y = y
	enemy.flags |= structures.MovedFlag
}

func (enemy *Enemy) Update(dt float64) bool {
	if enemy.Dead {
		return false
	}

	// logic

	return !enemy.Dead
}

func (enemy *Enemy) TakeDamage(damage int32) {
}

func (enemy *Enemy) Kill() {
	enemy.Dead = true
}

func (enemy *Enemy) ClearFlags() {
	enemy.flags &= structures.NoFlags
}
