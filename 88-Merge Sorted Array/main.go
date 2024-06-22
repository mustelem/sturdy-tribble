package main

import (
	"fmt"
)

func main() {
	foo := []int{1, 2, 3, 0, 0, 0}
	card1 := 3
	bar := []int{2, 5, 6}
	card2 := 3

	merge(foo, card1, bar, card2)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	retVal := make([]int, 0, m+n)
	p1, p2 := 0, 0

	for {
		if p2 >= n {
			retVal = append(retVal, nums1[p1:m]...)
			break
		}

		if p1 >= m {
			retVal = append(retVal, nums2[p2:n]...)
			break
		}

		if nums1[p1] > nums2[p2] {
			retVal = append(retVal, nums2[p2])
			p2++

		} else {
			retVal = append(retVal, nums1[p1])
			p1++
		}
	}

	fmt.Println(nums1)

	nums1 = retVal

	fmt.Println(nums1)
}
