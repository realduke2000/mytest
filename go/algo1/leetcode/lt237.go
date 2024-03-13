package leetcode

import "fmt"

func deleteNode(node *ListNode) {
	fmt.Printf("deleting %d\n", node.Val)
	if node == nil || node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
