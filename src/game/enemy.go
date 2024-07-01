package game

import (
	assets "main/assets/xml"
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
	Flags     int32
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
		Flags:     structures.NoFlags,
		Dead:      false,
	}
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
	enemy.Flags |= structures.MovedFlag
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
	enemy.Flags &= structures.NoFlags
}
