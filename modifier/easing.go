package modifier

import "math"

// EaseInOut returns an eased value for animations.
func EaseInOut(t float32) float32 {
	return float32(-(math.Cos(math.Pi*float64(t)) - 1) / 2)
}

// EaseOut returns an ease-out value.
func EaseOut(t float32) float32 {
	return float32(1 - math.Pow(1-float64(t), 3))
}
