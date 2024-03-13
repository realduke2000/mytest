package leetcode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	size := 0
	p := head
	for p != nil {
		p = p.Next
		size++
	}
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	p = head
	for i := 0; i < size/2; i++ {
		tmpPNext := p.Next
		tmpDummyNext := dummy.Next
		dummy.Next = p
		p.Next = tmpDummyNext
		p = tmpPNext
	}
	if size%2 != 0 {
		p = p.Next
	}

	dp := dummy.Next
	for p != nil && dp != nil {
		if p.Val != dp.Val {
			return false
		}
		p = p.Next
		dp = dp.Next
	}
	if p != nil || dp != nil {
		return false
	}
	return true
}
