package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nextVal := 0
	preHead := &ListNode{}
	node := preHead
	for {
		if l1 != nil || l2 != nil {
			sum := 0
			if l1 != nil && l2 != nil {
				sum = l1.Val + l2.Val + nextVal
				l1, l2 = l1.Next, l2.Next
			} else if l1 == nil && l2 != nil{
				sum = l2.Val + nextVal
				l2 = l2.Next
			} else if l1 != nil && l2 == nil {
				sum = l1.Val + nextVal
				l1 = l1.Next
			}
			nextVal = sum / 10
			node.Next = &ListNode{sum % 10,nil}
			node = node.Next
		} else {
			if nextVal != 0 {
				node.Next = &ListNode{nextVal,nil}
			}
			return preHead.Next
		}
	}
}


