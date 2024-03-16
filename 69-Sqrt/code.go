package main

import "math"

func mySqrt(x int) int {
	var result int
	i := 0

	for math.Pow(float64(i), 2) <= float64(x) {
		result = i
		i++
	}

	return result
}
