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

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 1)
	ret[0] = root.Val
	ret = visit(root.Left, ret)
	ret = visit(root.Right, ret)
	return ret
}

func visit(node *TreeNode, ret []int) []int {
	if node == nil {
		return ret
	}
	ret = append(ret, node.Val)
	ret = visit(node.Left, ret)
	ret = visit(node.Right, ret)
	return ret
}
