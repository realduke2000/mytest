package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cntA := 0
	cntB := 0
	for n := headA; n.Next != nil; n = n.Next {
		cntA++
	}
	for n := headB; n.Next != nil; n = n.Next {
		cntB++
	}

	var diff int
	na := headA
	nb := headB
	if cntA > cntB {
		diff = cntA - cntB
		for i := 0; i < diff; i++ {
			na = na.Next
		}
	} else {
		diff = cntB - cntA
		for i := 0; i < diff; i++ {
			nb = nb.Next
		}
	}
	for na != nil && nb != nil && na != nb {
		na = na.Next
		nb = nb.Next
	}
	return na
}
