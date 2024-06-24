package main

import "fmt"

func main() {
	rotate([]int{8, 9, 1, 10, 5, 2}, 3)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
}

func rotate(nums []int, k int) {
	fmt.Println(nums)
	cutoff := len(nums) - k
	nums = append(nums[cutoff:], nums[:cutoff]...)
	fmt.Println(nums)
}
