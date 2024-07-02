package utils

import "math"

// these are causing weird floating point inaccuracies

func Distance(x1 float32, y1 float32, x2 float32, y2 float32) float32 {
	dx := x1 - x2
	dy := y1 - y2
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func DistanceSqr(x1 float32, y1 float32, x2 float32, y2 float32) float32 {
	dx := x1 - x2
	dy := y1 - y2
	return dx*dx + dy*dy
}
