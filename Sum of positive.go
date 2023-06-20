package kata

func PositiveSum(numbers []int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			sum += numbers[i]
		} else {
			continue
		}
	}
	return sum
}
