package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	right := head
	// shift right one more step to find left previous node
	for i := 0; i < n+1; i++ {
		right = right.Next
		if right == nil {
			if i < n-1 {
				return head
			} else {
				if i == n-1 {
					// remove head
					ret := head.Next
					head.Next = nil
					return ret
				} else if i == n {
					// remove second
					removed := head.Next
					head.Next = removed.Next
					removed.Next = nil
					return head
				}
			}

		}
	}

	prev_left := head
	for right != nil {
		right = right.Next
		prev_left = prev_left.Next
	}
	toRemoved := prev_left.Next
	prev_left.Next = toRemoved.Next
	toRemoved.Next = nil
	return head
}
