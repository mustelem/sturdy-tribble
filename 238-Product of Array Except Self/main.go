package main

func productExceptSelf(nums []int) []int {
	var retSlice = make([]int, len(nums))
	var preMult = make([]int, len(nums))
	var postMult = make([]int, len(nums))

	preMult[0] = 1
	postMult[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		preMult[i] = preMult[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		postMult[i] = postMult[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		retSlice[i] = preMult[i] * postMult[i]
	}

	return retSlice
}
