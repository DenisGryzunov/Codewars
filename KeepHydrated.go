package main

import (
	"fmt"
	"math"
)

func Litres(time float64) int {
	drankLiters := int(math.Floor(time * 0.5))
	return drankLiters
}

func main() {
	fmt.Println(Litres(6.7))
}
