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
	cutoff := len(nums) - k

	copy(nums, append(nums[cutoff:], nums[:cutoff]...))

	fmt.Println(nums)
}
