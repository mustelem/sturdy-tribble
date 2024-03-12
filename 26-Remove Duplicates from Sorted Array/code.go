package main

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	cache := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == cache {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			cache = nums[i]
		}
	}

	return len(nums)
}
