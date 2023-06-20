package kata

func Grow(arr []int) int {
	product := arr[0]
	for i := 1; i < len(arr); i++ {
		product *= arr[i]
	}
	return product
}
