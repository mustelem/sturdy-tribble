package main

import (
	"math"
)

func plusOne(digits []int) []int {
	temp := toNum(digits)
	temp++
	array := toArray((temp))

	return array
}

func toNum(digits []int) float64 {
	var result float64

	var pow float64 = 0
	for i := len(digits) - 1; i >= 0; i-- {
		result += float64(digits[i]) * math.Pow(10, pow)
		pow++
	}

	return result
}

func toArray(number float64) []int {
	var revArr = make([]int, 0, 100)

	for number > 0 {
		digit := int(math.Mod(number, 10))
		revArr = append(revArr, digit)
		number = (number - float64(digit)) / 10
	}

	var arr = make([]int, 0, 100)
	for i := len(revArr) - 1; i >= 0; i-- {
		arr = append(arr, revArr[i])
	}

	return arr
}
