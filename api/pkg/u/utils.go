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
