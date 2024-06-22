package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	scores := map[int]int{}

	threshold := int(math.Floor(float64(len(nums)) / 2))

	for _, v := range nums {
		score, exists := scores[v]

		if exists {
			score++
		} else {
			score = 1
		}
		scores[v] = score

		if score > threshold {
			return v
		}
	}

	return 0
}
