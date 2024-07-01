package structures

import (
	"math"
)

type Position struct {
	X float32
	Y float32
}

func NewPosition(x float32, y float32) Position {
	return Position{
		X: x,
		Y: y,
	}
}

func (position Position) AngleTo(other Position) float32 {
	dx := other.X - position.X
	dy := other.Y - position.Y
	return float32(math.Atan2(float64(dy), float64(dx)))
}

func (position Position) PointAt(distance float32, angle float32) Position {
	newX := position.X + distance*float32(math.Cos(float64(angle)))
	newY := position.Y + distance*float32(math.Sin(float64(angle)))
	return Position{X: newX, Y: newY}
}

func (position Position) Distance(other Position) float32 {
	dx := other.X - position.X
	dy := other.Y - position.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (p Position) DistanceSqr(other Position) float32 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	return dx*dx + dy*dy
}

func (position Position) Dot(other Position) float32 {
	return position.X*other.X + position.Y*other.Y
}
