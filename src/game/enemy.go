package game

import (
	"math"
)

type Enemy struct {
	Id        int32
	MaxHealth int32
	Health    int32
	Type      int32
	X         float32
	Y         float32
	LastX     float32
	LastY     float32
	Flags     int32
	dead      bool
}

func NewEnemy(id int32) *Enemy {
	return &Enemy{
		Id:        id,
		MaxHealth: 0,
		Health:    0,
		Type:      0,
		X:         0,
		Y:         0,
		LastX:     0,
		LastY:     0,
		Flags:     0,
		dead:      false,
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
}

func (enemy *Enemy) Update(dt float64) bool {
	return !enemy.dead
}

func (enemy *Enemy) TakeDamage(damage int32) bool {

	return true
}

func (enemy *Enemy) Kill() {
	enemy.dead = true
}
