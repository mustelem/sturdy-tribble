package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root TreeNode = TreeNode{}

	rootValIdx := len(nums) / 2
	root.Val = nums[rootValIdx]

	if len(nums) > 2 {
		root.Left = sortedArrayToBST(nums[:rootValIdx])
		root.Right = sortedArrayToBST(nums[rootValIdx+1:])
	}

	if len(nums) == 2 {
		root.Left = sortedArrayToBST(nums[:rootValIdx])
	}

	if len(nums) == 1 {

	}

	return &root
}

func main() {

}
