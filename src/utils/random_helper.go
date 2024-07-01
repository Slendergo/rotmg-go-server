package utils

import (
	"math"
	"math/rand"
)

func NextInt32(min int32, max int32) int32 {
	return rand.Int31n(max-min) + min
}

func Normal32(mean, stdDev float32) float32 {
	var u1, u2, w, normal float32
	for {
		u1 = 2.0*rand.Float32() - 1.0
		u2 = 2.0*rand.Float32() - 1.0
		w = u1*u1 + u2*u2
		if w < 1 && w > 1e-30 {
			break
		}
	}
	w = float32(math.Sqrt(-2.0 * math.Log(float64(w)) / float64(w)))
	normal = u1 * w
	return normal*stdDev + mean
}
