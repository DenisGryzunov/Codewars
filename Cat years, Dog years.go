package kata

func CalculateYears(years int) (result [3]int) {
	if years == 1 {
		result[0] = years
		result[1] = 15
		result[2] = 15
	} else if years == 2 {
		result[0] = years
		result[1] = 24
		result[2] = 24
	} else if years > 2 {
		result[0] = years
		result[1] = 24 + ((years - 2) * 4)
		result[2] = 24 + ((years - 2) * 5)
	}
	return
}
