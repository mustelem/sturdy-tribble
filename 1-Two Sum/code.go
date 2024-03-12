package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)

	for i, v := range nums {
		complIndex, isMapContainsComplement := numbers[target-v]

		if !isMapContainsComplement {
			numbers[v] = i
		} else {
			return []int{complIndex, i}
		}

	}

	return []int{}
}

func main() {
	nums := []int{3, 9, 3, 8}
	target := 6

	matchingIndexes := twoSum(nums, target)

	fmt.Println(matchingIndexes)
}
