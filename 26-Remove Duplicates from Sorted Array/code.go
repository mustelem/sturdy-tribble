package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// nums := []int{1, 1, 2}
	// nums := []int{1, 1}

	fmt.Println(removeDuplicates(nums))
}

func walkArray(nums []int, index int) []int {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}

	return nums
}

func removeDuplicates(nums []int) int {
	var retVal int = len(nums)

	last := nums[len(nums)-1]

	for i := 0; i < len(nums); i++ {
		if nums[i] == last {
			retVal = i + 1
			break
		}

		if i == 0 {
			continue
		}

		for nums[i] == nums[i-1] {
			nums = walkArray(nums, i)
		}

		if nums[i] == last {
			retVal = i + 1
			break
		}
	}

	return retVal
}
