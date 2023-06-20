package kata

func MakeNegative(x int) int {
	if x > 0 {
		x = -x
	} else {
		x = x
	}
	return x
}
