package kata

func EvenOrOdd(number int) string {
	switch {
	case number%2 == 0:
		return "Even"
	default:
		return "Odd"
	}
}
