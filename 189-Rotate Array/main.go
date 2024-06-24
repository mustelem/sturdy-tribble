package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 10)
}

func rotate(nums []int, k int) {
	fmt.Println(nums)

	if k > len(nums) {
		k = k % len(nums)
	}

	cutoff := len(nums) - k // 4 = 7 - 3

	temp := make([]int, k)
	copy(temp, nums[cutoff:]) // temp = {5,6,7}

	for i := cutoff - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	} // nums = {1,2,3,1,2,3,4}

	for i, _ := range temp {
		nums[i] = temp[i]
	}

	fmt.Println(nums)
}
