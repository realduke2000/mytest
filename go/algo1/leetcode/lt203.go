package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	newHead := head
	head = head.Next
	newHead.Next = nil
	tail := newHead

	for head != nil {
		if head.Val == val {
			head = head.Next
			continue
		}
		tmp := head
		head = head.Next
		tmp.Next = nil
		tail.Next = tmp
		tail = tail.Next
	}
	return newHead
}
