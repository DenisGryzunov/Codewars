package kata

func Divisors(n int) int {
	var counter int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			counter++
		} else {
			continue
		}
	}
	return counter
}
