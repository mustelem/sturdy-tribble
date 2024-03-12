package main

import "fmt"

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}

	// var list1 *ListNode
	// var list2 *ListNode

	list3 := mergeTwoLists(list1, list2)
	if list3 != nil {
		list3.Print()
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (a *ListNode) Print() {
	for {
		fmt.Printf("%d,", a.Val)
		if a.Next == nil {
			break
		} else {
			a = a.Next
		}
	}
	fmt.Println()
}

func mergeTwoLists(p1 *ListNode, p2 *ListNode) *ListNode {
	if p1 == nil && p2 == nil {
		return nil
	}

	if p1 == nil {
		return p2
	}

	if p2 == nil {
		return p1
	}

	list3 := &ListNode{}
	p3 := list3

	for {

		if p1.Val <= p2.Val {
			p3.Next = p1
			if p1.Next == nil {
				p3.Next.Next = p2
				break
			} else {
				p1 = p1.Next
			}
			p3 = p3.Next
		} else {
			p3.Next = p2
			if p2.Next == nil {
				p3.Next.Next = p1
				break
			} else {
				p2 = p2.Next

			}
			p3 = p3.Next
		}
	}

	return list3.Next
}
