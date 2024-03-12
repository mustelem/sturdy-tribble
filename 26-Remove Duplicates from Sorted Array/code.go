package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// nums := []int{1, 1, 2}
	// nums := []int{1, 1}

	newNums := removeDuplicates(nums)
	fmt.Println(newNums)
}

func walkArray(nums []int, index int) []int {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}

	nums[len(nums)-1] = -999

	return nums
}

func removeDuplicates(nums []int) int {
	var count int = 0

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	for i := 1; i < len(nums); i++ {
		for nums[i] == nums[i-1] {
			nums = walkArray(nums, i)
		}

		if nums[i] < nums[i-1] {
			return i
		}

		count = i + 1
	}

	return count
}
