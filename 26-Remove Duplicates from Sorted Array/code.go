package main

func removeDuplicates(nums []int) int {
	var updatePtr int

	for i := 1; i < len(nums); i++ {
		if nums[updatePtr] != nums[i] {
			nums[updatePtr+1] = nums[i]
			updatePtr++
		}
	}

	return updatePtr + 1
}
