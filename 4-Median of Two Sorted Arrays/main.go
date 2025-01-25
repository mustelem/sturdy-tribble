package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	var medianIdxUp, medianIdxLow int

	total := n1 + n2

	// var median float64

	if total%2 == 0 {
		medianIdxUp = total / 2
		medianIdxLow = medianIdxUp - 1
		fmt.Println("Merged median Idx Low", medianIdxLow)
		fmt.Println("Merged median Idx Up", medianIdxUp)
	} else {
		medianIdxUp = (total/2 + 1) - 1
		fmt.Println("Median Idx", medianIdxUp)
	}

	var medianLongArr float64

	if n1 >= n2 {
		medianLongArr = findMedian(nums1)
		fmt.Println("Shorter array is: ", nums2)
	} else {
		medianLongArr = findMedian(nums2)
		fmt.Println("Shorter array is: ", nums1)
	}

	fmt.Println("Median of the longer array is: ", medianLongArr)

	return float64(medianIdxUp)
}

func findMedian(nums []int) float64 {
	var length int = len(nums)

	if length%2 == 0 {
		lowIdx := (length / 2) - 1
		result := (float64(nums[lowIdx]) + float64(nums[lowIdx+1])) / 2
		return result
	}

	return float64(nums[length/2])
}

func countBumps(nums []int, median float64) {

}

func main() {
	findMedianSortedArrays([]int{1, 6, 13}, []int{4, 5, 7, 34})
}
