package main

func removeDuplicates(nums []int) int {
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[i-2] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
