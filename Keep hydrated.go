package kata

import (
	"math"
)

func Litres(time float64) int {
	drankLiters := int(math.Floor(time * 0.5))
	return drankLiters
}
