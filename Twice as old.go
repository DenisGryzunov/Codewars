package kata

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	diff := dadYearsOld - 2*sonYearsOld
	if diff < 0 {
		diff = int(math.Abs(float64(diff)))
	}
	return diff
}
