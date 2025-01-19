package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack := make([]ListNode, 0)

	var atHand int
	for {

		if l1 == nil && l2 == nil {
			if atHand == 1 {
				stack = append(stack, ListNode{1, nil})
				break
			}

			break
		}

		var res int
		var left int
		var right int

		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}

		res = left + right + atHand
		if res > 9 {
			atHand = 1
			res = res - atHand*10
		} else {
			atHand = 0
		}

		stack = append(stack, ListNode{res, nil})
	}

	root := stack[len(stack)-1]
	for i := len(stack) - 2; i >= 0; i-- {
		child := root
		root = stack[i]
		root.Next = &child
	}

	return &root
}
