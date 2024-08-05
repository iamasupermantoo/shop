package utils

import "math"

// FloatAccuracy 返回指定精度float
func FloatAccuracy(f float64, n int) float64 {
	shift := math.Pow(10, float64(n))
	return math.Round(f*shift) / shift
}
