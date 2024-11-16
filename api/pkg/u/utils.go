package u

import (
	"math"
)

const (
	SECONDS_PER_YEAR float64 = 31536000
)

// may be modified to take compounding frequency as an argument
func ToAPY(apr float64) float64 {
	return math.Pow(1+apr/SECONDS_PER_YEAR, SECONDS_PER_YEAR) - 1
}

func Map[T, R any] (arr []T, f func(T) R) []R {
	res := make([]R, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}
	return res
}

func All[T any] (arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}
	return true
}
